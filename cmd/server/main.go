package main

import (
	"log"

	"github.com/AVZotov/draft-survey/internal/handler"
	"github.com/AVZotov/draft-survey/internal/storage"
	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()

	userStore := storage.NewUserStore("./users")

	h := handler.New(userStore)

	handler.SetupRoutes(app, h)

	log.Fatal(app.Listen(":3000"))
}
