package routes

import (
	"github.com/goravel/framework/contracts/http"

	"goravel/app/facades"
)

func Web() {
	facades.Route().Get("/", func(ctx http.Context) http.Response {
		return ctx.Response().View().Make("home.tmpl", map[string]any{
			"title": "Goravel Link",
		})
	})
}
