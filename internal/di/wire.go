//go:build wireinject
// +build wireinject

package di

import (
	"context"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"github.com/google/wire"
	"github.com/umardev500/go-laundry/internal/bootstrap"
	"github.com/umardev500/go-laundry/internal/config"
	"github.com/umardev500/go-laundry/internal/ent"
	"github.com/umardev500/go-laundry/internal/modules/auth"
)

func ProvideContext() context.Context {
	return context.Background()
}

func ProvideValidator() *validator.Validate {
	return validator.New()
}

var AppSet = wire.NewSet(
	auth.ProvideSet,
	bootstrap.ProvideFiberApp,
	ProvideEntClient,
	ProvideContext,
	ProvideValidator,
)

func InitializeFiberApp(cfg *config.AppConfig) *fiber.App {
	wire.Build(AppSet)

	return nil
}

func GetEntClient(cfg *config.AppConfig) *ent.Client {
	wire.Build(AppSet)
	return nil
}
