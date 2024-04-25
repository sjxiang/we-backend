package main

import (
	"github.com/rs/zerolog/log"

	"we-backend/pkg/config"
	"we-backend/pkg/di"
)


func main() {
	
	cfg, err := config.LoadConfig()
	if err != nil {
		log.Fatal().Err(err).Msg("Failed to load the config")
	}

	server, err := di.InitializeApi(cfg)
	if err != nil {
		log.Fatal().Err(err).Msg("Failed to initialize the api")
	}

	if err := server.Start(); err != nil {
		log.Fatal().Err(err).Msg("Failed to start server")
	}
}
