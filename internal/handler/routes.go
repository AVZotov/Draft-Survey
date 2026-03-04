package handler

import "github.com/gofiber/fiber/v2"

func SetupRoutes(app *fiber.App, h *Handler) {
	app.Static("/static", "./web/static")

	app.Get("/", h.home)
	app.Get("/profile", h.profile)
	app.Get("/survey/new", h.newSurvey)
	app.Get("/survey/:id", h.getSurvey)
	app.Get("/survey/:id/draft", h.draftReadings)

	api := app.Group("/api/v1")
	api.Post("/profile", h.createProfile)
	api.Post("/survey", h.saveSurvey)
}
