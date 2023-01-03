package http

import (
	"fmt"
	"strings"

	"github.com/valyala/fasthttp"

	"github.com/yunginnanet/HellPot/internal/config"
)

func robotsTXT(ctx *fasthttp.RequestCtx) {
	slog := log.With().
		Str("USERAGENT", string(ctx.UserAgent())).
		Str("REMOTE_ADDR", remoteAddr).
		Interface("URL", string(ctx.RequestURI())).Logger()
	paths := &strings.Builder{}
	paths.WriteString("User-agent: *\r\n")
	for _, p := range config.Paths {
		paths.WriteString("Disallow: ")
		paths.WriteString(p)
		paths.WriteString("\r\n")
	}
	paths.WriteString("\r\n")

	slog.Debug().
		Strs("PATHS", config.Paths).
		Msg("SERVE_ROBOTS")

	if _, err := fmt.Fprintf(ctx, paths.String()); err != nil {
		slog.Error().Err(err).Msg("SERVE_ROBOTS_ERROR")
	}
}
