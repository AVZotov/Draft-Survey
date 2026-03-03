package handler

import (
	"strconv"

	"github.com/gofiber/fiber/v2"
)

func parseFloat(c *fiber.Ctx, field string) float64 {
	v := c.FormValue(field)
	if v == "" {
		return 0
	}
	f, _ := strconv.ParseFloat(v, 64)
	return f
}

func parseInt(c *fiber.Ctx, field string) int {
	v := c.FormValue(field)
	if v == "" {
		return 0
	}
	i, _ := strconv.Atoi(v)
	return i
}
