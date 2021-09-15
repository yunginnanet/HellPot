package main

import (
	"os"
	"os/signal"
	"syscall"

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
	log.Info().Str("file", config.Filename).Msg("current config")
	log.Info().Str("file", config.CurrentLogFile).Msg("current log")
}

func main() {
	stopChan := make(chan os.Signal, 1)
	signal.Notify(stopChan, syscall.SIGINT, syscall.SIGTERM)

	go func() {
		log.Error().Err(http.Serve()).Msg("HTTP err")
	}()

	<-stopChan // wait for SIGINT
	log.Warn().Msg("Shutting down server...")

}
