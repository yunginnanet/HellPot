package main

import (
	"github.com/rs/zerolog"

	"github.com/yunginnanet/HellPot/config"
	"github.com/yunginnanet/HellPot/extra"
	"github.com/yunginnanet/HellPot/http"
)

var log zerolog.Logger

func init() {
	config.Init()
	log = config.StartLogger()
	extra.Banner()
	if config.Debug {
		log.Debug().Msg("debug enabled")
	}
	log.Info().Str("config", config.Filename).Msg("initialization finished")
}

func main() {
	log.Error().Err(http.Serve()).Msg("HTTP err")
}
