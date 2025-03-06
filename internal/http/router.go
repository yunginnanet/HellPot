package http

import (
	"bufio"
	"fmt"
	"io"
	"net/http"
	"os"
	"runtime"
	"strconv"
	"strings"
	"time"

	"github.com/fasthttp/router"
	"github.com/valyala/fasthttp"

	"github.com/yunginnanet/HellPot/heffalump"
	"github.com/yunginnanet/HellPot/internal/config"
	"github.com/yunginnanet/HellPot/internal/logger"
)

var (
	log              *logger.Log
	hellpotHeffalump *heffalump.Heffalump
	runningConfig    *config.Parameters
)

func getRealRemote(ctx *fasthttp.RequestCtx) string {
	xrealip := string(ctx.Request.Header.Peek(runningConfig.HTTP.ProxiedIPHeader))
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

	if runningConfig.HTTP.ClientRules.MatchUseragent(ctx.UserAgent()) {
		slog.Trace().Msg("Ignoring useragent")
		ctx.Error("Not found", http.StatusNotFound)
		return
	}

	if runningConfig.Logger.Trace {
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
	if !runningConfig.Perf.ConcurrencyCap {
		runningConfig.Perf.MaxWorkers = fasthttp.DefaultConcurrency
	}

	log = runningConfig.GetLogger()

	return fasthttp.Server{
		// User defined server name
		// Likely not useful if behind a reverse proxy without additional configuration of the proxy server.
		Name: runningConfig.Liar.FakeServerName,

		/*
			from fasthttp docs: "By default request read timeout is unlimited."
										Nope.
		*/
		ReadTimeout:        5 * time.Second,
		MaxRequestBodySize: 0.5 * 1024 * 1024,

		MaxConnsPerIP:      3,
		MaxRequestsPerConn: 2,
		Concurrency:        runningConfig.Perf.MaxWorkers,

		GetOnly: !runningConfig.IdleHands.POSTMimicry,

		// we don't care if a request ends up being handled by a different handler (in fact it probably will)
		KeepHijackedConns: true,

		CloseOnShutdown: true,

		// No need to keepalive, our response is a sort of keep-alive ;)
		DisableKeepalive: true,

		Handler: r.Handler,
		Logger:  log,
	}
}

func SetupHeffalump(config *config.Parameters) error {
	switch config.Bespoke.CustomHeffalump {
	case true:
		content, err := os.ReadFile(config.Bespoke.Grimoire)
		if err != nil {
			return fmt.Errorf("failed to read grimoire file '%s': %w", config.Bespoke.Grimoire, err)
		}
		// Wasteful, but only done once at startup
		src := string(content)
		log.Info().Msgf("Using custom grimoire file '%s'", config.Bespoke.Grimoire)

		if len(src) < 1 {
			return fmt.Errorf("%w: grimoire file '%s' appears to be empty", io.EOF, config.Bespoke.Grimoire)
		}

		if !strings.Contains(src, "<") || !strings.Contains(src, ">") {
			return fmt.Errorf("%w: grimoire file '%s' does not appear to be a valid source text, missing brackets",
				io.ErrUnexpectedEOF, config.Bespoke.Grimoire)
		}

		markovMap := heffalump.MakeMarkovMap(strings.NewReader(src))
		hellpotHeffalump = heffalump.NewHeffalump(markovMap, heffalump.DefaultBuffSize)
	default:
		log.Info().Msg("Using default source text")
		hellpotHeffalump = heffalump.NewDefaultHeffalump()
	}
	if hellpotHeffalump == nil {
		panic("failed to initialize heffalump")
	}
	return nil
}

// Serve starts our HTTP server and request router
func Serve(config *config.Parameters) error {
	log = config.GetLogger()
	runningConfig = config

	if err := SetupHeffalump(config); err != nil {
		return fmt.Errorf("failed to setup heffalump: %w", err)
	}

	l := config.HTTP.Bind + ":" + strconv.Itoa(int(config.HTTP.Port))

	r := router.New()

	if config.HTTP.Router.MakeRobots && !config.HTTP.Router.CatchAll {
		r.GET("/robots.txt", robotsTXT)
	}

	if !config.HTTP.Router.CatchAll {
		for _, p := range config.HTTP.Router.Paths {
			log.Trace().Str("caller", "router").Msgf("Add route: %s", p)
			r.GET(fmt.Sprintf("/%s", p), hellPot)
		}
	} else {
		log.Trace().Msg("Catch-All mode enabled...")
		r.GET("/{path:*}", hellPot)
	}

	srv := getSrv(r)

	//goland:noinspection GoBoolExpressions
	if !config.HTTP.UnixSocket.UseUnixSocket || runtime.GOOS == "windows" {
		log.Info().Str("caller", l).Msg("Listening and serving HTTP...")
		err := srv.ListenAndServe(l)
		if err != nil {
			return fmt.Errorf("failed to start HTTP server: %w", err)
		}
	}

	if len(config.HTTP.UnixSocket.UnixSocketPath) < 1 {
		log.Fatal().Msg("unix_socket_path configuration directive appears to be empty")
	}

	log.Info().Str("caller", config.HTTP.UnixSocket.UnixSocketPath).Msg("Listening and serving HTTP...")
	listener, err := listenOnUnixSocket(config.HTTP.UnixSocket.UnixSocketPath, r)
	if err != nil {
		return fmt.Errorf("failed to start unix listener: %w", err)
	}

	err = srv.Serve(listener)
	if err != nil {
		err = fmt.Errorf("failed to serve HTTP: %w", err)
	}
	return err
}
