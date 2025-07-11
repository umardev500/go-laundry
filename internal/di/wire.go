//go:build wireinject
// +build wireinject

package di

import (
	"github.com/gofiber/fiber/v2"
	"github.com/google/wire"
	"github.com/umardev500/go-laundry/internal/bootstrap"
	"github.com/umardev500/go-laundry/internal/config"
	"github.com/umardev500/go-laundry/internal/modules/auth"
)

var AppSet = wire.NewSet(
	auth.ProvideSet,
	bootstrap.ProvideFiberApp,
)

func InitializeFiberApp(cfg *config.AppConfig) *fiber.App {
	wire.Build(AppSet)

	return nil
}
