package http

import (
	"fmt"
	"net"
	"syscall"
	"time"

	"github.com/fasthttp/router"
	"github.com/rs/zerolog"
	"github.com/valyala/fasthttp"

	"github.com/yunginnanet/HellPot/config"
	"github.com/yunginnanet/HellPot/heffalump"
)

var log zerolog.Logger

func getRealRemote(ctx *fasthttp.RequestCtx) string {
	xrealip := string(ctx.Request.Header.Peek("X-Real-IP"))
	if len(xrealip) > 0 {
		return xrealip
	}
	return ctx.RemoteIP().String()
}

func scopeCheck(ctx *fasthttp.RequestCtx) {
	var inscope = false
	reqpath := ctx.QueryArgs().Peek("path")

	remoteAddr := getRealRemote(ctx)

	slog := log.With().
		Str("UserAgent", string(ctx.UserAgent())).
		Str("REMOTE_ADDR", remoteAddr).
		Interface("URL", string(ctx.RequestURI())).Logger()

	for _, p := range config.Paths {
		if string(reqpath) == p {
			inscope = true
			break
		}
	}

	if !inscope {
		slog.Warn().Msg("Request outside of configured scope!")
		ctx.Error("", 404)
		return
	}

	s := time.Now()

	n := heffalump.DefaultHeffalump.WriteHell(ctx.Response.BodyWriter())

	slog.Info().
		Int64("BYTES", n).
		Dur("DURATION", time.Since(s)).
		Msg("FINISH")

}

func listenOnUnixSocket(addr string, r *router.Router) error {
	var err error
	var unixAddr *net.UnixAddr
	var unixListener *net.UnixListener
	unixAddr, err = net.ResolveUnixAddr("unix", addr)
	if err == nil {
		// Always unlink sockets before listening on them
		if err2 := syscall.Unlink(addr); err2 != nil {
			panic(err2)
		}
		unixListener, err = net.ListenUnix("unix", unixAddr)
		if err == nil {
			err = fasthttp.Serve(unixListener, r.Handler)
		}
	}
	return err
}

// Serve starts our HTTP server and request router
func Serve() error {
	log = config.GetLogger()

	l := fmt.Sprintf("%s:%s", config.BindAddr, config.BindPort)

	r := router.New()
	r.GET("/robots.txt", robotsTXT)
	r.GET("/{path}", scopeCheck)

	if !config.RestrictConcurrency {
		config.MaxWorkers = fasthttp.DefaultConcurrency
	}

	srv := fasthttp.Server{
		// User defined server name
		// Likely not useful if behind a reverse proxy without additional configuration of the proxy server.
		Name: config.FakeServerName,

		/*
		from fasthttp docs: "By default request read timeout is unlimited."
		My thinking here is avoiding some sort of weird oversized GET query just in case.
		*/
		ReadTimeout: 5 * time.Second,
		MaxRequestBodySize: 1 * 1024 * 1024,

		// Help curb abuse of HellPot (we've always needed this badly)
		MaxConnsPerIP: 1,
		MaxRequestsPerConn: 2,
		Concurrency: config.MaxWorkers,

		// only accept GET requests
		GetOnly: true,

		// we don't care if a request ends up being handled by a different handler
		// it shouldn't for now, but this may prevent future bugs
		KeepHijackedConns: true,

		CloseOnShutdown: true,



		// No need to keepalive, our response is a sort of keep-alive ;)
		DisableKeepalive: true,



		Handler: r.Handler,
	}

	if !config.UseUnixSocket {
		log.Info().Str("caller", l).Msg("Listening and serving HTTP...")
		return srv.ListenAndServe(l)
	}

	if len(config.UnixSocketPath) < 1 {
		log.Fatal().Msg("unix_socket_path configuration directive appears to be empty")
	}

	log.Info().Str("caller", config.UnixSocketPath).Msg("Listening and serving HTTP...")
	return listenOnUnixSocket(config.UnixSocketPath, r)
}
