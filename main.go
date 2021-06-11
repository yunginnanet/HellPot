package main

import (
	// zerolog json logging and console pretty printing
	"HellPot/src/logger"
	"github.com/rs/zerolog"

	// viper configuration engine for a toml config file
	"HellPot/src/config"

	// ascii banners and other aesthetic shit
	"HellPot/src/decorate"

	// bitcask embedded key/value database
	//"HellPot/src/casket"

)

var log zerolog.Logger

//   TODO:
//// optional bitcask database initialization
//// fix default config file writing

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
