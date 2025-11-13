//go:build wireinject
// +build wireinject

package app

import (
	"github.com/google/wire"
	"github.com/umardev500/laundry/internal/config"
	"github.com/umardev500/laundry/internal/db"
	"github.com/umardev500/laundry/internal/handler"
	"github.com/umardev500/laundry/internal/repository"
	"github.com/umardev500/laundry/internal/service"
)

var AppSet = wire.NewSet(
	New,
	NewServer,
	handler.Set,
	service.Set,
	repository.Set,
	NewRoutes,
	db.NewClient,
)

func Initialize(config *config.Config) *App {
	wire.Build(AppSet)
	return nil
}
