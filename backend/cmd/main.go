package main

import (
	"github.com/esceer/due-dash/backend/cmd/setup"
	"github.com/rs/zerolog/log"
)

func main() {
	// Config
	cfg, err := setup.Config()
	if err != nil {
		log.Fatal().Msgf("Startup failed while reading config: %v", err)
	}
	setup.Logger(cfg)

	// Database
	if err = setup.RunMigrationScripts(cfg); err != nil {
		log.Fatal().Err(err).Msg("DB migration failed")
	}
	database, err := setup.ConnectToDB(cfg)
	if err != nil {
		log.Fatal().Err(err).Msg("Connecting to DB failed")
	}

	// Services
	services := setup.Services(cfg, database)

	// Http server
	server := setup.WebRouting(cfg, services)
	log.Info().Msgf("Listening on %v...", cfg.ServerAddress)
	server.Logger.Fatal(server.Start(cfg.ServerAddress))
}
