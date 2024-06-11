package main

import (
	"os"

	"github.com/rs/zerolog/log"
	"github.com/rs/zerolog"

	"we-backend/pkg/config"
	"we-backend/pkg/di"
)

func init() {
	log.Logger = log.Output(zerolog.ConsoleWriter{Out: os.Stderr})  // 控制台输出、文件输出	
}

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
