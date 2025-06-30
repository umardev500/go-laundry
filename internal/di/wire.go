//go:build wireinject
// +build wireinject

package di

import (
	"context"

	"github.com/gofiber/fiber/v2"
	"github.com/google/wire"
	"github.com/umardev500/go-laundry/internal/app"
	"github.com/umardev500/go-laundry/internal/config"
	"github.com/umardev500/go-laundry/internal/ent"
	"github.com/umardev500/go-laundry/internal/handler/http"
	"github.com/umardev500/go-laundry/internal/repository"
	"github.com/umardev500/go-laundry/internal/service"
)

var UserSet = wire.NewSet(
	http.NewUserHandler,
	service.NewUserService,
	repository.NewUserRepository,
)

func ProvideContext() context.Context {
	return context.Background()
}

var AppSet = wire.NewSet(
	ProvideContext,
	config.LoadDatabaseConfig,
	UserSet,
	app.ProvideFiberApp,
	ProvideEntClient,
)

func InitializeFiberApp() *fiber.App {
	wire.Build(AppSet)
	return nil
}

func GetEntClient() *ent.Client {
	wire.Build(AppSet)
	return nil
}
