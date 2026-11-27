package requests

import (
	"github.com/goravel/framework/contracts/http"
	"github.com/goravel/framework/contracts/validation"
)

type StoreLinkRequest struct {
	URL string `form:"url" json:"url"`
}

func (r *StoreLinkRequest) Authorize(ctx http.Context) error {
	return nil
}

func (r *StoreLinkRequest) Filters(ctx http.Context) map[string]any {
	return map[string]any{
		"url": "trim",
	}
}

func (r *StoreLinkRequest) Rules(ctx http.Context) map[string]any {
	return map[string]any{
		// regex: hanya izinkan http/https supaya tidak ada skema aneh seperti javascript:
		"url": "required|url|max_len:2048|regex:^https?://",
	}
}

func (r *StoreLinkRequest) Messages(ctx http.Context) map[string]string {
	return map[string]string{
		"url.required": "URL wajib diisi.",
		"url.url":      "Format URL tidak valid, contoh: https://contoh.com",
		"url.max_len":  "URL terlalu panjang (maksimal 2048 karakter).",
		"url.regex":    "URL harus diawali http:// atau https://",
	}
}

func (r *StoreLinkRequest) Attributes(ctx http.Context) map[string]string {
	return map[string]string{}
}

func (r *StoreLinkRequest) PrepareForValidation(ctx http.Context, data validation.Data) error {
	return nil
}
