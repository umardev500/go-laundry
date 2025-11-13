package main

import (
	"os"

	"github.com/joho/godotenv"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"github.com/umardev500/laundry/internal/config"
	"github.com/umardev500/laundry/internal/db"
	"github.com/umardev500/laundry/internal/seeder"
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
	client := db.NewClient(cfg)

	if err := seeder.RunAll(client); err != nil {
		log.Fatal().Err(err).Msg("Failed to run seeder")
	}
}
