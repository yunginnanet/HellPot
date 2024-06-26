package main

import (
	"os"
	"strconv"

	"github.com/yunginnanet/HellPot/internal/config"
	"github.com/yunginnanet/HellPot/internal/http"
	"github.com/yunginnanet/HellPot/internal/version"
)

var (
	runningConfig *config.Parameters
)

func main() {
	stopChan := make(chan os.Signal, 1)
	log, _, resolvedConf, err := setup(stopChan)

	if err != nil {
		println("failed to start: " + err.Error())
		os.Exit(1)
	}

	log.Info().Msg("🔥 Starting HellPot 🔥")
	log.Info().Msg("Version: " + version.Version)
	log.Info().Msg("PID: " + strconv.Itoa(os.Getpid()))
	log.Info().Msg("Using config file: " + resolvedConf)
	if runningConfig.UsingDefaults {
		log.Warn().Msg("Using default configuration")
	}
	if runningConfig.Logger.RSyslog != "" {
		log.Info().Msg("Logging to syslog: " + runningConfig.Logger.RSyslog)
	}
	if runningConfig.Logger.ActiveLogFileName != "" {
		log.Info().Msg("Logging to file: " + runningConfig.Logger.ActiveLogFileName)
	}
	if runningConfig.Logger.DockerLogging &&
		runningConfig.Logger.File == "" &&
		runningConfig.Logger.Directory == "" &&
		runningConfig.Logger.RSyslog == "" {
		log.Info().Msg("Only logging to standard output")
	}

	log.Debug().Msg("Debug logging enabled")
	log.Trace().Msg("Trace logging enabled")

	go func() {
		log.Fatal().Err(http.Serve(runningConfig)).Msg("HTTP error")
	}()

	sig := <-stopChan // wait for SIGINT
	log.Warn().Interface("signal_received", sig).
		Msg("Shutting down server...")
}
