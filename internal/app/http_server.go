package app

import (
	"github.com/gofiber/fiber/v2"
	"github.com/umardev500/go-laundry/internal/handler/http"
)

func ProvideFiberApp(
	userHandler *http.UserHandler,
	authHandler *http.AuthHandler,
	merchantHandler *http.MerchantHandler,
) *fiber.App {
	app := fiber.New(fiber.Config{
		DisableStartupMessage: true,
	})

	api := app.Group("api")

	userHandler.Setup(api.Group("users"))
	authHandler.Setup(api.Group("auth"))
	merchantHandler.Setup(api.Group("merchants"))

	return app
}
