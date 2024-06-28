package main

import (
	"os"

	"github.com/rs/zerolog"

	"github.com/yunginnanet/HellPot/internal/config"
	"github.com/yunginnanet/HellPot/internal/logger"
	"github.com/yunginnanet/HellPot/internal/version"
)

func printInfo(log *logger.Log, resolvedConf string, realConfig *config.Parameters) {
	log.Info().
		Str("version", version.Version).
		Int("pid", os.Getpid()).
		Str("config_file", resolvedConf).
		Str("log_file", log.Config.ActiveLogFileName).
		Msg("🔥 Starting HellPot 🔥")
	if realConfig.UsingDefaults {
		log.Warn().Msg("Using default configuration! Please edit the configuration file to suit your needs.")
	}

	var eventer func() *zerolog.Event

	if realConfig.Logger.Debug && !realConfig.Logger.Trace {
		eventer = log.Debug
	}
	if realConfig.Logger.Trace {
		eventer = log.Trace
	}

	if eventer == nil {
		return
	}

	eventer().Bool("docker_logging", true).
		Bool("debug", true).Bool("trace", false).
		Str("rsyslog", realConfig.Logger.RSyslog).
		Str("log_file", realConfig.Logger.ActiveLogFileName).
		Msg("logging engine started")
}
