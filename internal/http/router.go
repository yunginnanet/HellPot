package http

import (
	"bufio"
	"fmt"
	"net/http"
	"os"
	"runtime"
	"strings"
	"time"

	"github.com/fasthttp/router"
	"github.com/rs/zerolog"
	"github.com/valyala/fasthttp"

	"github.com/yunginnanet/HellPot/heffalump"
	"github.com/yunginnanet/HellPot/internal/config"
)

var (
	log              *zerolog.Logger
	hellpotHeffalump *heffalump.Heffalump
)

func getRealRemote(ctx *fasthttp.RequestCtx) string {
	xrealip := string(ctx.Request.Header.Peek(config.HeaderName))
	if len(xrealip) > 0 {
		return xrealip
	}
	return ctx.RemoteIP().String()
}

func hellPot(ctx *fasthttp.RequestCtx) {
	path, pok := ctx.UserValue("path").(string)
	if len(path) < 1 || !pok {
		path = "/"
	}

	remoteAddr := getRealRemote(ctx)

	slog := log.With().
		Str("USERAGENT", string(ctx.UserAgent())).
		Str("REMOTE_ADDR", remoteAddr).
		Interface("URL", string(ctx.RequestURI())).Logger()

	for _, denied := range config.UseragentBlacklistMatchers {
		if strings.Contains(string(ctx.UserAgent()), denied) {
			slog.Trace().Msg("Ignoring useragent")
			ctx.Error("Not found", http.StatusNotFound)
			return
		}
	}

	if config.Trace {
		slog = slog.With().Str("caller", path).Logger()
	}

	slog.Info().Msg("NEW")

	s := time.Now()
	var n int64

	ctx.SetBodyStreamWriter(func(bw *bufio.Writer) {
		var err error
		var wn int64

		for {
			wn, err = hellpotHeffalump.WriteHell(bw)
			n += wn
			if err != nil {
				slog.Trace().Err(err).Msg("END_ON_ERR")
				break
			}
		}

		slog.Info().
			Int64("BYTES", n).
			Dur("DURATION", time.Since(s)).
			Msg("FINISH")
	})
}

func getSrv(r *router.Router) fasthttp.Server {
	if !config.RestrictConcurrency {
		config.MaxWorkers = fasthttp.DefaultConcurrency
	}

	log = config.GetLogger()

	return fasthttp.Server{
		// User defined server name
		// Likely not useful if behind a reverse proxy without additional configuration of the proxy server.
		Name: config.FakeServerName,

		/*
			from fasthttp docs: "By default request read timeout is unlimited."
			My thinking here is avoiding some sort of weird oversized GET query just in case.
		*/
		ReadTimeout:        5 * time.Second,
		MaxRequestBodySize: 1 * 1024 * 1024,

		// Help curb abuse of HellPot (we've always needed this badly)
		MaxConnsPerIP:      10,
		MaxRequestsPerConn: 2,
		Concurrency:        config.MaxWorkers,

		// only accept GET requests
		GetOnly: true,

		// we don't care if a request ends up being handled by a different handler (in fact it probably will)
		KeepHijackedConns: true,

		CloseOnShutdown: true,

		// No need to keepalive, our response is a sort of keep-alive ;)
		DisableKeepalive: true,

		Handler: r.Handler,
		Logger:  log,
	}
}

// Serve starts our HTTP server and request router
func Serve() error {
	log = config.GetLogger()

	if config.UseCustomHeffalump {
		content, err := os.ReadFile(config.BookFilename)
		if err != nil {
			panic(err)
		}
		// Wasteful, but only done once at startup
		src := string(content)
		log.Info().Msgf("Using custom book file '%s'", config.BookFilename)

		if len(src) < 1 {
			panic("book file was empty!")
		}

		markovMap := heffalump.MakeMarkovMap(strings.NewReader(src))
		hellpotHeffalump = heffalump.NewHeffalump(markovMap, heffalump.DefaultBuffSize)
	} else {
		log.Info().Msg("Using default source text")
		hellpotHeffalump = heffalump.NewDefaultHeffalump()
	}

	l := config.HTTPBind + ":" + config.HTTPPort

	r := router.New()

	if config.MakeRobots && !config.CatchAll {
		r.GET("/robots.txt", robotsTXT)
	}

	if !config.CatchAll {
		for _, p := range config.Paths {
			log.Trace().Str("caller", "router").Msgf("Add route: %s", p)
			r.GET(fmt.Sprintf("/%s", p), hellPot)
		}
	} else {
		log.Trace().Msg("Catch-All mode enabled...")
		r.GET("/{path:*}", hellPot)
	}

	srv := getSrv(r)

	//goland:noinspection GoBoolExpressions
	if !config.UseUnixSocket || runtime.GOOS == "windows" {
		log.Info().Str("caller", l).Msg("Listening and serving HTTP...")
		return srv.ListenAndServe(l)
	}

	if len(config.UnixSocketPath) < 1 {
		log.Fatal().Msg("unix_socket_path configuration directive appears to be empty")
	}

	log.Info().Str("caller", config.UnixSocketPath).Msg("Listening and serving HTTP...")
	return listenOnUnixSocket(config.UnixSocketPath, r)
}
