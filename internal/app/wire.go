//go:build wireinject
// +build wireinject

package app

import (
	"github.com/google/wire"
	"github.com/umardev500/go-laundry/internal/config"
	"github.com/umardev500/go-laundry/internal/db"
	"github.com/umardev500/go-laundry/internal/module/auth"
	"github.com/umardev500/go-laundry/internal/module/user"
	"github.com/umardev500/go-laundry/internal/seed"
	"github.com/umardev500/go-laundry/internal/types"
)

func InitApp(cfg *config.Config) (*App, error) {
	wire.Build(
		db.NewEntClient,
		NewFiberApp,
		NewApp,
	)
	return nil, nil
}

func InitSeeders(cfg *config.Config) ([]types.Seeder, error) {
	wire.Build(
		db.NewEntClient,
		user.ProviderSet,
		auth.ProviderSet,
		seed.ProvideSeeders,
	)
	return nil, nil
}
