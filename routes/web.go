package routes

import (
	"goravel/app/facades"
	"goravel/app/http/controllers"
)

func Web() {
	linkController := controllers.NewLinkController()

	facades.Route().Get("/", linkController.Index)
	facades.Route().Post("/links", linkController.Store)
	facades.Route().Get("/{code}", linkController.Show)
}
