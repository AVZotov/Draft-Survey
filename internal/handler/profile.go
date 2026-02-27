package handler

import (
	"github.com/AVZotov/draft-survey/internal/handler/tadaptor"
	"github.com/AVZotov/draft-survey/web"
	"github.com/AVZotov/draft-survey/web/components"
	"github.com/gofiber/fiber/v2"
)

func (h *Handler) profile(c *fiber.Ctx) error {
	BannerProps, _ := c.Locals("banner").(components.BannerProps)
	component := web.Profile(BannerProps)
	return tadaptor.Render(c, component)
}
