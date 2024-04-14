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

	switch config.DockerLogging {
	case true:
		config.CurrentLogFile = "/dev/stdout"
		config.NoColor = true
		log = config.StartLogger(false, os.Stdout)
	default:
		log = config.StartLogger(true)
	}

	extra.Banner()

	log.Info().Str("caller", "config").Str("file", config.Filename).Msg(config.Filename)
	log.Info().Str("caller", "logger").Msg(config.CurrentLogFile)
	log.Debug().Str("caller", "logger").Msg("debug enabled")
	log.Trace().Str("caller", "logger").Msg("trace enabled")

}

func main() {
	stopChan := make(chan os.Signal, 1)
	signal.Notify(stopChan, syscall.SIGINT, syscall.SIGTERM)

	go func() {
		log.Fatal().Err(http.Serve()).Msg("HTTP error")
	}()

	<-stopChan // wait for SIGINT
	log.Warn().Msg("Shutting down server...")

}
