package controllers

import (
	"crypto/rand"
	"math/big"
	nethttp "net/http"

	"github.com/goravel/framework/contracts/http"
	"github.com/goravel/framework/validation"

	"goravel/app/facades"
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
	validator, err := ctx.Request().Validate(map[string]any{
		// regex: hanya izinkan http/https supaya tidak ada skema aneh seperti javascript:
		"url": "required|url|max_len:2048|regex:^https?://",
	}, validation.Filters(map[string]any{
		"url": "trim",
	}), validation.Messages(map[string]string{
		"url.required": "URL wajib diisi.",
		"url.url":      "Format URL tidak valid, contoh: https://contoh.com",
		"url.max_len":  "URL terlalu panjang (maksimal 2048 karakter).",
		"url.regex":    "URL harus diawali http:// atau https://",
	}))
	if err != nil {
		return ctx.Response().String(nethttp.StatusInternalServerError, err.Error())
	}
	var req struct {
		URL string `form:"url" json:"url"`
	}
	if err := validator.Bind(&req); err != nil {
		return ctx.Response().String(nethttp.StatusInternalServerError, err.Error())
	}
	if errs := validator.Errors(); errs != nil {
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

func (r *LinkController) render(ctx http.Context, data map[string]any) http.Response {
	data["title"] = "Goravel Link"
	data["host"] = ctx.Request().Host()

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
