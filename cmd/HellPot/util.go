package main

import (
	"os"
	"strconv"

	"github.com/rs/zerolog"

	"github.com/yunginnanet/HellPot/internal/config"
	"github.com/yunginnanet/HellPot/internal/version"
)

func printInfo(log zerolog.Logger, resolvedConf string, realConfig *config.Parameters) {
	log.Info().Msg("🔥 Starting HellPot 🔥")
	if realConfig.UsingDefaults {
		log.Warn().Msg("Using default configuration! Please edit the configuration file to suit your needs.")
	}
	log.Info().Msg("Version: " + version.Version)
	log.Info().Msg("PID: " + strconv.Itoa(os.Getpid()))
	log.Info().Msg("Using config file: " + resolvedConf)
	if realConfig.Logger.RSyslog != "" {
		log.Info().Msg("Logging to syslog: " + realConfig.Logger.RSyslog)
	}
	if realConfig.Logger.ActiveLogFileName != "" {
		log.Info().Msg("Logging to file: " + realConfig.Logger.ActiveLogFileName)
	}
	if realConfig.Logger.DockerLogging &&
		realConfig.Logger.File == "" &&
		realConfig.Logger.Directory == "" &&
		realConfig.Logger.RSyslog == "" {
		log.Info().Msg("Only logging to standard output")
	}
	log.Debug().Msg("Debug logging enabled")
	log.Trace().Msg("Trace logging enabled")
}
