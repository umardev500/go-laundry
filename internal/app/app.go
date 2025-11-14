package app

import (
	"fmt"

	"github.com/rs/zerolog/log"
	"github.com/umardev500/laundry/internal/config"
	"github.com/umardev500/routerx"
)

type App struct {
	config *config.Config
	server *routerx.App
	routes []Route
}

func New(cfg *config.Config, srv *routerx.App, routes []Route) *App {
	return &App{
		config: cfg,
		server: srv,
		routes: routes,
	}
}

func (a *App) Run() error {
	addr := ":8080"
	if a.config.App.Port > 0 {
		addr = fmt.Sprintf(":%d", a.config.App.Port)
	}

	api := a.server.Group("/api")

	for _, route := range a.routes {
		route.Register(api)
	}

	log.Info().Msgf("Total module registered: %d", len(a.routes))
	log.Info().Msgf("Starting server on port %s", addr)
	return a.server.Listen(&addr)
}
