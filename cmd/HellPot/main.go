package main

import (
	"github.com/rs/zerolog"

	"github.com/yunginnanet/HellPot/config"
	"github.com/yunginnanet/HellPot/extra"
)

var log zerolog.Logger

func init() {
	config.Init()
	log = config.Logger()
	extra.Banner()
	if config.Debug {
		log.Debug().Msg("debug enabled")
	}
	log.Info().Str("file", config.Filename).Msg("current config")
	log.Info().Str("file", config.CurrentLogFile).Msg("current log")
}

func main() {
	startPot()
}
