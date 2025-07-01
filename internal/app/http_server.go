package app

import (
	"github.com/gofiber/fiber/v2"
	"github.com/umardev500/go-laundry/internal/handler/http"
)

func ProvideFiberApp(
	userHandler *http.UserHandler,
	authHandler *http.AuthHandler,
) *fiber.App {
	app := fiber.New(fiber.Config{
		DisableStartupMessage: true,
	})

	regular := app.Group("/api")
	protected := app.Group("/api")

	userHandler.Setup(regular, protected)
	authHandler.Setup(regular, protected)

	return app
}
