package main

import (
	// zerolog json logging and console pretty printing
	"github.com/rs/zerolog"
	"github.com/yunginnanet/HellPot/src/logger"

	// viper configuration engine for a toml config file
	"github.com/yunginnanet/HellPot/src/config"

	// ascii banners and other aesthetic shit
	"github.com/yunginnanet/HellPot/src/decorate"
	// bitcask embedded key/value database
	//"github.com/yunginnanet/HellPot/src/casket"
)

var log zerolog.Logger

func init() {
	// configuration engine
	config.Blueprint()

	// style points
	decorate.Banner()

	// buffered configuration engine log entries _after_ the banner
	config.PrintConfigLog()

	// initialize logger
	logger.LogInit()
	log = logger.GlobalLogger

	// initialize configuration file/engine
	zerolog.SetGlobalLevel(zerolog.InfoLevel)
	if config.Debug {
		zerolog.SetGlobalLevel(zerolog.DebugLevel)
	}

	// bitcask embedded key/value database
	//casket.Initialize()
}

func main() {
	startPot()
}
