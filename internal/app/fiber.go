package app

import (
	"github.com/gofiber/fiber/v2"
	"github.com/umardev500/go-laundry/internal/module/auth"
	"github.com/umardev500/go-laundry/internal/module/user"
)

func NewFiberApp(
	authHandler *auth.Handler,
	userHandler *user.Handler,
) *fiber.App {
	app := fiber.New(fiber.Config{
		DisableStartupMessage: true,
	})

	api := app.Group("/api")

	authHandler.SetupRoutes(api)
	userHandler.SetupRoutes(api)

	return app
}
