package main

import (
	"os"
	"os/signal"
	"syscall"

	"github.com/rs/zerolog"

	"github.com/yunginnanet/HellPot/internal/config"
	"github.com/yunginnanet/HellPot/internal/extra"
	"github.com/yunginnanet/HellPot/internal/http"
)

var (
	log     zerolog.Logger
	version string // set by linker
)

func init() {
	if version != "" {
		config.Version = version[1:]
	}
	config.Init()
	if config.BannerOnly {
		extra.Banner()
		os.Exit(0)
	}
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
		log.Error().Err(http.Serve()).Msg("HTTP error")
	}()

	<-stopChan // wait for SIGINT
	log.Warn().Msg("Shutting down server...")

}
