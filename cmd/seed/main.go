package main

import (
	"context"
	"os"

	"github.com/joho/godotenv"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"github.com/umardev500/go-laundry/internal/di"
	"github.com/umardev500/go-laundry/internal/seeds"
)

func init() {
	log.Logger = log.Output(zerolog.ConsoleWriter{Out: os.Stdout})

	err := godotenv.Load()
	if err != nil {
		panic(err)
	}
}

func main() {
	log.Info().Msg("seeding data...")

	client := di.GetEntClient()
	defer client.Close()

	// Start transaction
	tx, err := client.Tx(context.Background())
	if err != nil {
		log.Fatal().Err(err).Msg("failed to start transaction")
	}

	// Seed data
	if err := seeds.SeedMerchants(context.Background(), tx); err != nil {
		log.Fatal().Err(err).Msg("failed to seed merchants")
	}

	// Commit transaction
	if err := tx.Commit(); err != nil {
		log.Fatal().Err(err).Msg("failed to commit transaction")
	}

	log.Info().Msg("data seeded successfully")
}
