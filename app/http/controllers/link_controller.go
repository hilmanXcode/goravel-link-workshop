package controllers

import (
	"crypto/rand"
	"math/big"
	nethttp "net/http"

	"github.com/goravel/framework/contracts/http"

	"goravel/app/facades"
	"goravel/app/http/requests"
	"goravel/app/models"
)

const (
	codeAlphabet = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	codeLength   = 6
)

type LinkController struct{}

func NewLinkController() *LinkController {
	return &LinkController{}
}

// Index menampilkan halaman utama berisi form pembuat link.
func (r *LinkController) Index(ctx http.Context) http.Response {
	return r.render(ctx, map[string]any{
		"created": ctx.Request().Query("created"),
	})
}

// Store memvalidasi URL, menyimpannya dengan kode acak, lalu kembali ke halaman utama.
func (r *LinkController) Store(ctx http.Context) http.Response {
	var req requests.StoreLinkRequest
	errs, err := ctx.Request().ValidateRequest(&req)
	if err != nil {
		return ctx.Response().String(nethttp.StatusInternalServerError, err.Error())
	}
	if errs != nil {
		return r.render(ctx, map[string]any{
			"error": errs.One(),
			"url":   ctx.Request().Input("url"),
		})
	}

	link := models.Link{Code: newCode(), URL: req.URL}
	// shortcut: tabrakan kode (peluang sangat kecil) jadi error 500, tambah retry kalau trafik besar.
	if err := facades.Orm().Query().Create(&link); err != nil {
		return ctx.Response().String(nethttp.StatusInternalServerError, err.Error())
	}

	return ctx.Response().Redirect(nethttp.StatusSeeOther, "/?created="+link.Code)
}

// Show mencari link berdasarkan kode, menambah hitungan klik, lalu redirect ke URL asli.
func (r *LinkController) Show(ctx http.Context) http.Response {
	var link models.Link
	if err := facades.Orm().Query().Where("code", ctx.Request().Route("code")).First(&link); err != nil {
		return ctx.Response().String(nethttp.StatusInternalServerError, err.Error())
	}
	if link.ID == 0 {
		return ctx.Response().String(nethttp.StatusNotFound, "Link tidak ditemukan")
	}

	// UPDATE atomik di database, aman kalau banyak klik bersamaan
	if _, err := facades.Orm().Query().Exec("UPDATE links SET clicks = clicks + 1 WHERE id = ?", link.ID); err != nil {
		return ctx.Response().String(nethttp.StatusInternalServerError, err.Error())
	}

	return ctx.Response().Redirect(nethttp.StatusFound, link.URL)
}

func (r *LinkController) render(ctx http.Context, data map[string]any) http.Response {
	var links []models.Link
	if err := facades.Orm().Query().Order("id desc").Limit(20).Get(&links); err != nil {
		return ctx.Response().String(nethttp.StatusInternalServerError, err.Error())
	}

	data["title"] = "Goravel Link"
	data["host"] = ctx.Request().Host()
	data["links"] = links

	return ctx.Response().View().Make("home.tmpl", data)
}

// newCode membuat kode acak yang aman (crypto/rand) dari alfabet base62.
func newCode() string {
	code := make([]byte, codeLength)
	for i := range code {
		n, err := rand.Int(rand.Reader, big.NewInt(int64(len(codeAlphabet))))
		if err != nil {
			panic(err) // crypto/rand gagal berarti OS bermasalah, tidak ada yang bisa dipulihkan
		}
		code[i] = codeAlphabet[n.Int64()]
	}

	return string(code)
}
