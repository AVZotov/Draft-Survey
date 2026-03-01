package handler

import (
	"errors"
	"os"

	"github.com/AVZotov/draft-survey/internal/handler/tadaptor"
	"github.com/AVZotov/draft-survey/web"
	"github.com/gofiber/fiber/v2"
)

func (h *Handler) home(c *fiber.Ctx) error {
	_, err := h.userRepository.Get()

	if err == nil {
		component := web.Dashboard()
		return tadaptor.Render(c, component)
	}

	if errors.Is(err, os.ErrNotExist) {
		return h.profile(c)
	}
	// If file corrupted or no access show profile and banner with warning
	c.Locals("banner", web.BannerFileCorrupted)
	return h.profile(c)
}
