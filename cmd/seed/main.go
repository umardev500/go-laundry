package main

import (
	"context"
	"fmt"
	"os"

	"github.com/joho/godotenv"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"github.com/umardev500/go-laundry/internal/app"
	"github.com/umardev500/go-laundry/internal/config"
	"github.com/umardev500/go-laundry/internal/db"
	"github.com/umardev500/go-laundry/internal/seed"
)

func init() {
	log.Logger = log.Output(zerolog.ConsoleWriter{Out: os.Stdout})

	err := godotenv.Load()
	if err != nil {
		panic(err)
	}
}

func main() {
	cfg := config.LoadConfig("./config/config.yml")
	client := db.NewEntClient(cfg).Client
	defer client.Close()

	seeders, err := app.InitSeeders(cfg)
	if err != nil {
		log.Fatal().Err(err).Msg("Failed to initialize seeders")
	}

	fmt.Println(seeders)

	if err := seed.Run(context.Background(), seeders); err != nil {
		log.Fatal().Err(err).Msg("Failed to seed database")
	}
}
