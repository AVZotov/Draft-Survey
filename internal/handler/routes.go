package handler

import "github.com/gofiber/fiber/v2"

func SetupRoutes(app *fiber.App, h *Handler) {
	app.Static("/static", "./web/static")

	app.Get("/", h.home)
	app.Get("/profile", h.profile)

	api := app.Group("/api/v1")
	api.Post("/profile", h.createProfile)
	//api.Post("newdraft", h.newDraft)

	// api.Put("/profile", h.UpdateProfile)
}
