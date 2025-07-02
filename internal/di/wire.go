//go:build wireinject
// +build wireinject

package di

import (
	"context"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"github.com/google/wire"
	"github.com/umardev500/go-laundry/internal/app"
	"github.com/umardev500/go-laundry/internal/config"
	"github.com/umardev500/go-laundry/internal/ent"
	"github.com/umardev500/go-laundry/internal/handler/http"
	"github.com/umardev500/go-laundry/internal/repository"
	"github.com/umardev500/go-laundry/internal/service"
	"github.com/umardev500/go-laundry/internal/usecase"
	"github.com/umardev500/go-laundry/pkg/transaction"
)

var UserSet = wire.NewSet(
	http.NewUserHandler,
	service.NewUserService,
	repository.NewUserRepository,
)

var AuthSet = wire.NewSet(
	http.NewAuthHandler,
	service.NewAuthService,
)

var MerchantSet = wire.NewSet(
	http.NewMerchantHandler,
	usecase.NewMerchantRegisterUsecase,
	repository.NewMerchantRepository,
)

func ProvideContext() context.Context {
	return context.Background()
}

func ProvideValidator() *validator.Validate {
	return validator.New()
}

var AppSet = wire.NewSet(
	ProvideContext,
	ProvideValidator,
	config.LoadDatabaseConfig,
	AuthSet,
	UserSet,
	MerchantSet,
	app.ProvideFiberApp,
	ProvideEntClient,
	transaction.NewTransactionManager,
)

func InitializeFiberApp() *fiber.App {
	wire.Build(AppSet)
	return nil
}

func GetEntClient() *ent.Client {
	wire.Build(AppSet)
	return nil
}
