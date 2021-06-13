package main

import (
	"context"
	"github.com/gorilla/mux"
	"github.com/yunginnanet/HellPot/src/config"
	"net"
	"net/http"
	"os"
	"os/signal"
	"strings"
	"syscall"
	"time"
)

const robotsTxt = "User-agent: *\r\n"

func listenOnUnixSocket(addr string, srv *http.Server) error {
	var err error
	var unixAddr *net.UnixAddr
	var unixListener *net.UnixListener
	socketPath := strings.TrimPrefix(addr, "unix:")
	unixAddr, err = net.ResolveUnixAddr("unix", socketPath)
	if err == nil {
		// Always unlink sockets before listening on them
		_ = syscall.Unlink(socketPath)
		unixListener, err = net.ListenUnix("unix", unixAddr)
		if err == nil {
			err = srv.Serve(unixListener)
		}
	}
	return err
}

func startPot() {
	addr := config.BindAddr
	port := config.BindPort
	listenOnUnixDomain := true

	// subscribe to SIGINT signals
	stopChan := make(chan os.Signal, 1)
	signal.Notify(stopChan, syscall.SIGINT, syscall.SIGTERM)

	r := mux.NewRouter()

	r.HandleFunc("/{path}", DefaultHoneypot)

	srv := &http.Server{Handler: r}
	if !strings.HasPrefix(addr, "unix:") {
		listenOnUnixDomain = false
		srv.Addr = addr + ":" + port
	}

	go func() {
		log.Info().Str("bind_addr", addr).Str("bind_port", port).
			Msg("Listening and serving HTTP...")
		// service connections
		var err error
		if listenOnUnixDomain {
			err = listenOnUnixSocket(addr, srv)
		} else {
			err = srv.ListenAndServe()
		}
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
