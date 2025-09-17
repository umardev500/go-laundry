package app

import "github.com/gofiber/fiber/v2"

func NewFiberApp() *fiber.App {
	app := fiber.New(fiber.Config{
		DisableStartupMessage: true,
	})

	return app
}
