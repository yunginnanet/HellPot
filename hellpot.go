package main

import (
	"context"
	"io"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

const robotsTxt = "User-agent: *\r\nDisallow: "

func startPot() {
	addr := os.Getenv("HONEYADDR")
	if addr == "" {
		addr = "127.0.0.1"
	}

	port := os.Getenv("HONEYPORT")
	if port == "" {
		port = "8080"
	}

	path := os.Getenv("HONEYPATH")
	if path == "" {
		path = "/wp-login.php"
	}

	// subscribe to SIGINT signals
	stopChan := make(chan os.Signal, 1)
	signal.Notify(stopChan, syscall.SIGINT, syscall.SIGTERM)

	http.HandleFunc(path, func(w http.ResponseWriter, r *http.Request) {
		DefaultHoneypot(w, r)
	})

	http.HandleFunc("/robots.txt", func(w http.ResponseWriter, r *http.Request) {
		if _, err := io.WriteString(w, robotsTxt+path+"\r\n"); err != nil {
			log.Error().Err(err).Msg("SERVE_ROBOTS_ERROR")
		}
	})

	srv := &http.Server{Addr: addr + ":" + port, Handler: http.DefaultServeMux}

	go func() {
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
