package main

import (
	"HellPot/src/config"
	"context"
	"io"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

const robotsTxt = "User-agent: *\r\n"

func startPot() {
	var paths string
	addr := config.BindAddr
	port := config.BindPort

	// subscribe to SIGINT signals
	stopChan := make(chan os.Signal, 1)
	signal.Notify(stopChan, syscall.SIGINT, syscall.SIGTERM)

	for _, p := range config.Paths {
		http.HandleFunc(p, func(w http.ResponseWriter, r *http.Request) {
			DefaultHoneypot(w, r)
		})
		paths = paths + "Disallow: " + p + "\r\n"
	}

	http.HandleFunc("/robots.txt", func(w http.ResponseWriter, r *http.Request) {
		if _, err := io.WriteString(w, robotsTxt+paths+"\r\n"); err != nil {
			log.Error().Err(err).Msg("SERVE_ROBOTS_ERROR")
		}
	})

	srv := &http.Server{Addr: addr + ":" + port, Handler: http.DefaultServeMux}

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
