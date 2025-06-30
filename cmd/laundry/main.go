package main

import (
	"context"
	"os"

	"github.com/joho/godotenv"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
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
	port := os.Getenv("PORT")
	app := di.InitializeFiberApp()
	entClient := di.GetEntClient()

	// Migrate the schema
	if err := entClient.Schema.Create(context.TODO()); err != nil {
		log.Fatal().Err(err).Msg("failed creating schema resources")
	}

	log.Info().
		Str("port", port).
		Msg("server started successfully")

	app.Listen(":" + port)
}
