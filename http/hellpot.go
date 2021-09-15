package http

import (
	"net"
	"net/http"
	"os"
	"os/signal"
	"strings"
	"syscall"

	"github.com/yunginnanet/HellPot/config"
)

var log = config.GetLogger()

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

func listen() {
	addr := config.BindAddr
	port := config.BindPort
	listenOnUnixDomain := true

	log.Info().Str("bind_addr", addr).Str("bind_port", port).
		Msg("Listening and serving HTTP...")

	if !strings.HasPrefix(addr, "unix:") {
		listenOnUnixDomain = false
	}

/*	if listenOnUnixDomain {
		err = listenOnUnixSocket(addr, srv)
	} else {
		err = srv.ListenAndServe()
	}*/

	// log.Warn().Err(err).Msg("HTTP_STOP")
}

func StartPot() {

	// subscribe to SIGINT signals
	stopChan := make(chan os.Signal, 1)
	signal.Notify(stopChan, syscall.SIGINT, syscall.SIGTERM)

	/*r := mux.NewRouter()

	r.HandleFunc("/{path}", heffalump.DefaultHoneypot)

	srv := &http.Server{Handler: r}

*/
	go listen()

	<-stopChan // wait for SIGINT
	log.Warn().Msg("Shutting down server...")

	/*
	// shut down gracefully, but wait no longer than 5 seconds before halting
	ctx, c := context.WithTimeout(context.Background(), 5*time.Second)
	defer c()
	if err := srv.Shutdown(ctx); err != nil {
		log.Debug().Caller().Err(err)
	}
	log.Info().Msg("Server gracefully stopped")
	 */
}
