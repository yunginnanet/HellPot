package http

import (
	"fmt"

	"github.com/yunginnanet/common/pool"
	"github.com/valyala/fasthttp"
)

var strs = pool.NewStringFactory()

func robotsTXT(ctx *fasthttp.RequestCtx) {
	config := runningConfig.HTTP.Router
	slog := log.With().
		Str("USERAGENT", string(ctx.UserAgent())).
		Str("REMOTE_ADDR", getRealRemote(ctx)).
		Interface("URL", string(ctx.RequestURI())).Logger()
	pathBuf := strs.Get()
	pathBuf.MustWriteString("User-agent: *\r\n")
	for _, p := range config.Paths {
		pathBuf.MustWriteString("Disallow: ")
		pathBuf.MustWriteString(p)
		pathBuf.MustWriteString("\r\n")
	}
	pathBuf.MustWriteString("\r\n")
	paths := pathBuf.String()
	strs.MustPut(pathBuf)

	slog.Debug().
		Strs("PATHS", config.Paths).
		Msg("SERVE_ROBOTS")

	if _, err := fmt.Fprintf(ctx, paths); err != nil {
		slog.Error().Err(err).Msg("SERVE_ROBOTS_ERROR")
	}
}
