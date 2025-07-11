package main

import (
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
	app := di.InitializeFiberApp(cfg)

	port := ":" + strconv.Itoa(cfg.Port)

	log.Info().Msgf("Listening on port %s", port)

	app.Listen(port)
}
