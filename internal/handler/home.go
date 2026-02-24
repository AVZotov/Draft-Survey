package handler

import "github.com/gofiber/fiber/v2"

func (h *Handler) Home(c *fiber.Ctx) error {
	return c.SendString("TODO: check profile")
}
