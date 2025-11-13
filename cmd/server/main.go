package main

import (
	"context"
	"os"
	"os/signal"
	"syscall"

	"github.com/joho/godotenv"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"github.com/umardev500/laundry/internal/app"
	"github.com/umardev500/laundry/internal/config"
	"github.com/umardev500/laundry/internal/db"
)

func init() {
	if err := godotenv.Load(); err != nil {
		log.Error().Err(err).Msg("Error loading .env file")
	}

	if os.Getenv("APP_ENV") == "development" {
		log.Logger = log.Output(zerolog.ConsoleWriter{Out: os.Stdout})
	} else {
		log.Logger = log.Output(os.Stdout)
	}
}

func main() {
	cfg := config.LoadConfig()
	app := app.Initialize(cfg)
	client := db.NewClient(cfg)

	// Run auto migration tool.
	log.Info().Msg("Running auto migration")
	if err := client.Client().Schema.Create(context.Background()); err != nil {
		log.Fatal().Err(err).Msg("Failed to run auto migration")
	}

	go func() {
		if err := app.Run(); err != nil {
			log.Fatal().Err(err).Msg("Server stopped")
		}
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit

	log.Info().Msg("Shutting down server...")
}
