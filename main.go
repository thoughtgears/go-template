package main

import (
	"github.com/thoughtgears/go-template/internal/config"
	"github.com/thoughtgears/go-template/internal/router"

	"github.com/kelseyhightower/envconfig"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

var cfg config.Config

func init() {
	envconfig.MustProcess("", &cfg)

	zerolog.SetGlobalLevel(zerolog.InfoLevel)
	zerolog.TimeFieldFormat = zerolog.TimeFormatUnix
	zerolog.LevelFieldName = "severity"
}

func main() {
	r := router.NewRouter(&cfg)
	log.Fatal().Err(r.Run()).Msg("Failed to start the server")
}
