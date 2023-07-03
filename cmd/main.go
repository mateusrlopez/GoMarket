package main

import (
	"github.com/joho/godotenv"
	"github.com/mateusrlopez/go-market/internal/configurations"
	"github.com/mateusrlopez/go-market/internal/server"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

func init() {
	zerolog.TimeFieldFormat = zerolog.TimeFormatUnix
}

func main() {
	applicationConfiguration := configurations.NewApplicationConfiguration()

	if applicationConfiguration.Environment == "development" {
		if err := godotenv.Load(); err != nil {
			log.Fatal().Err(err).Msg("failed to load environment variables from .env file")
		}
	}

	if err := server.New().ListenAndServe(); err != nil {
		log.Fatal().Err(err).Msg("could not start the server")
	}
}
