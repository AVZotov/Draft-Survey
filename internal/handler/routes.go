package handler

import "github.com/gofiber/fiber/v2"

func SetupRoutes(app *fiber.App, h *Handler) {
	app.Static("/static", "./web/static")

	app.Get("/", h.home)
	app.Get("/profile", h.profile)

	// HTMX формы
	// api := app.Group("/api/v1")
	// api.Post("/profile", h.CreateProfile)   // создание (первый запуск)
	// api.Put("/profile", h.UpdateProfile)
}
