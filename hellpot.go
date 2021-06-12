package main

import (
	"context"
	"github.com/gorilla/mux"
	"github.com/yunginnanet/HellPot/src/config"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

const robotsTxt = "User-agent: *\r\n"

func startPot() {
	addr := config.BindAddr
	port := config.BindPort

	// subscribe to SIGINT signals
	stopChan := make(chan os.Signal, 1)
	signal.Notify(stopChan, syscall.SIGINT, syscall.SIGTERM)

	r := mux.NewRouter()

	r.HandleFunc("/{path}", DefaultHoneypot)

	srv := &http.Server{Addr: addr + ":" + port, Handler: r}

	go func() {
		log.Info().Str("bind_addr", addr).Str("bind_port", port).
			Msg("Listening and serving HTTP...")
		// service connections
		err := srv.ListenAndServe()
		log.Warn().Err(err).Msg("HTTP_STOP")
	}()

	<-stopChan // wait for SIGINT
	log.Warn().Msg("Shutting down server...")

	// shut down gracefully, but wait no longer than 5 seconds before halting
	ctx, c := context.WithTimeout(context.Background(), 5*time.Second)
	defer c()
	srv.Shutdown(ctx)

	log.Info().Msg("Server gracefully stopped")
}
