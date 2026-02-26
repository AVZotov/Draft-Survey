package handler

import (
	"github.com/AVZotov/draft-survey/internal/handler/tadaptor"
	"github.com/AVZotov/draft-survey/web"
	"github.com/gofiber/fiber/v2"
)

func (h *Handler) profile(c *fiber.Ctx) error {
	component := web.Profile()
	return tadaptor.Render(c, component)
}
