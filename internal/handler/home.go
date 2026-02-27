package handler

import (
	"errors"
	"os"

	"github.com/AVZotov/draft-survey/web"
	"github.com/gofiber/fiber/v2"
)

func (h *Handler) home(c *fiber.Ctx) error {
	_, err := h.userRepository.Get()

	if err == nil {
		// TODO: Need to make dashboard
		return c.SendString("TODO: dashboard")
	}

	if errors.Is(err, os.ErrNotExist) {
		return h.profile(c)
	}
	// If file corrupted or no access show profile and banner with warning
	c.Locals("banner", web.BannerFileCorrupted)
	return h.profile(c)
}
