package handler

import (
	"net/http"

	"github.com/AVZotov/draft-survey/internal/handler/tadaptor"
	"github.com/AVZotov/draft-survey/internal/types"
	"github.com/AVZotov/draft-survey/web"
	"github.com/AVZotov/draft-survey/web/components"
	"github.com/gofiber/fiber/v2"
)

func (h *Handler) profile(c *fiber.Ctx) error {
	BannerProps, _ := c.Locals("banner").(components.BannerProps)
	component := web.Profile(BannerProps)
	return tadaptor.Render(c, component)
}

func (h *Handler) createProfile(c *fiber.Ctx) error {
	//TODO: Validation of user
	user := new(types.User)
	if err := c.BodyParser(user); err != nil {
		//ToDo. Warn Banner with general issue
		return err
	}

	if err := h.userRepository.Save(user); err != nil {
		if err := c.BodyParser(user); err != nil {
			//ToDo. Warn Banner with general issue
			return err
		}
	}

	c.Set("HX-Redirect", "/")
	return c.SendStatus(http.StatusOK)
}
