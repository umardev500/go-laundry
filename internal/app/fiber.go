package app

import (
	"github.com/gofiber/fiber/v2"
	"github.com/umardev500/go-laundry/internal/module/auth"
	"github.com/umardev500/go-laundry/internal/module/plan"
	"github.com/umardev500/go-laundry/internal/module/registration"
	"github.com/umardev500/go-laundry/internal/module/role"
	"github.com/umardev500/go-laundry/internal/module/subscription"
	"github.com/umardev500/go-laundry/internal/module/user"
)

func NewFiberApp(
	authHandler *auth.Handler,
	userHandler *user.Handler,
	registrationHandler *registration.Handler,
	roleHandler *role.Handler,
	subscriptionHandler *subscription.Handler,
	planHandler *plan.Handler,
) *fiber.App {
	app := fiber.New(fiber.Config{
		DisableStartupMessage: true,
	})

	api := app.Group("/api")

	authHandler.SetupRoutes(api)
	userHandler.SetupRoutes(api)
	registrationHandler.SetupRoutes(api)
	roleHandler.SetupRoutes(api)
	subscriptionHandler.SetupRoutes(api)
	planHandler.SetupRoutes(api)

	return app
}
