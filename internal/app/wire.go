//go:build wireinject
// +build wireinject

package app

import (
	"github.com/google/wire"
	"github.com/umardev500/go-laundry/internal/config"
	"github.com/umardev500/go-laundry/internal/db"
	"github.com/umardev500/go-laundry/internal/module/auth"
	"github.com/umardev500/go-laundry/internal/module/feature"
	"github.com/umardev500/go-laundry/internal/module/permission"
	"github.com/umardev500/go-laundry/internal/module/plan"
	"github.com/umardev500/go-laundry/internal/module/registration"
	"github.com/umardev500/go-laundry/internal/module/role"
	"github.com/umardev500/go-laundry/internal/module/subscription"
	"github.com/umardev500/go-laundry/internal/module/tenant"
	"github.com/umardev500/go-laundry/internal/module/user"
	"github.com/umardev500/go-laundry/internal/seed"
	"github.com/umardev500/go-laundry/internal/types"
	"github.com/umardev500/go-laundry/pkg/email"
	"github.com/umardev500/go-laundry/pkg/validator"
)

var AppSet = wire.NewSet(
	db.NewEntClient,
	db.NewRedisClient,
	validator.New,
	user.ProviderSet,
	auth.ProviderSet,
	feature.ProviderSet,
	registration.ProviderSet,
	tenant.ProviderSet,
	role.ProviderSet,
	plan.ProviderSet,
	permission.ProviderSet,
	email.NewClient,
	subscription.ProviderSet,
	seed.ProvideSeeders,
)

func InitApp(cfg *config.Config) (*App, error) {
	wire.Build(
		AppSet,
		NewFiberApp,
		NewApp,
	)
	return nil, nil
}

func InitSeeders(cfg *config.Config) ([]types.Seeder, error) {
	wire.Build(
		AppSet,
	)
	return nil, nil
}
