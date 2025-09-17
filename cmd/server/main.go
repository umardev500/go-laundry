package main

import (
	"context"
	"os"
	"os/signal"
	"syscall"

	"github.com/joho/godotenv"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"github.com/umardev500/go-laundry/internal/app"
	"github.com/umardev500/go-laundry/internal/config"
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
	application, err := app.InitApp(cfg)
	if err != nil {
		log.Fatal().Err(err).Msg("Failed to initialize app")
	}

	go func() {
		if err := application.Run(); err != nil {
			log.Fatal().Err(err).Msg("Failed to start server")
		}
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit

	if err := application.Shutdown(context.Background()); err != nil {
		log.Fatal().Err(err).Msg("Failed to shutdown server")
	}

}
