package main

import (
	"context"
	"os"
	"strconv"

	"github.com/joho/godotenv"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"github.com/umardev500/go-laundry/internal/config"
	"github.com/umardev500/go-laundry/internal/di"
)

func init() {
	log.Logger = log.Output(zerolog.ConsoleWriter{Out: os.Stdout})

	err := godotenv.Load()
	if err != nil {
		panic(err)
	}
}

func main() {
	cfg := config.Load()
	client := di.GetEntClient(cfg)
	app := di.InitializeFiberApp(cfg)

	port := ":" + strconv.Itoa(cfg.Port)

	// Migrate database
	if err := client.Schema.Create(context.Background()); err != nil {
		log.Fatal().Err(err).Msg("failed creating schema resources")
	}

	log.Info().Msgf("Listening on port %s", port)

	app.Listen(port)
}
