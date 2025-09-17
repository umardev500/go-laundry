package app

import (
	"context"

	"github.com/gofiber/fiber/v2"
	"github.com/rs/zerolog/log"
	"github.com/umardev500/go-laundry/internal/config"
	"github.com/umardev500/go-laundry/internal/db"
)

type App struct {
	cfg       *config.Config
	entClient *db.Client
	fiber     *fiber.App
}

func NewApp(cfg *config.Config, client *db.Client, fiber *fiber.App) *App {
	return &App{
		cfg:       cfg,
		entClient: client,
		fiber:     fiber,
	}
}

func (a *App) Run() error {
	addr := ":" + a.cfg.Server.Port
	log.Printf("🚀 Server running on %s", addr)
	return a.fiber.Listen(addr)
}

func (a *App) Shutdown(ctx context.Context) error {
	log.Info().Msg("Shutting down server")
	if err := a.entClient.Client.Close(); err != nil {
		return err
	}
	return a.fiber.Shutdown()
}
