package main

import (
	"os"

	"github.com/yunginnanet/HellPot/internal/http"
)

func main() {
	stopChan := make(chan os.Signal, 1)
	log, resolvedConf, realConf, err := setup(stopChan)

	if err != nil {
		println("failed to start: " + err.Error())
		os.Exit(1)
	}

	printInfo(log, resolvedConf, realConf)

	go func() {
		log.Fatal().Err(http.Serve(realConf)).Msg("HTTP error")
	}()

	sig := <-stopChan // wait for SIGINT
	log.Warn().Interface("signal_received", sig).
		Msg("Shutting down server...")
}
