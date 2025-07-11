package bootstrap

import (
	"github.com/gofiber/fiber/v2"
	"github.com/umardev500/go-laundry/internal/config"
	"github.com/umardev500/go-laundry/internal/modules/auth"
)

func ProvideFiberApp(
	authHandler auth.AuthHandler,
	config *config.AppConfig,
) *fiber.App {
	app := fiber.New(config.FiberConfig)

	authHandler.Setup(app.Group("/auth"))

	return app
}
