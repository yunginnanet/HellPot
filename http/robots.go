package http

import (
	"fmt"
	"strings"

	"github.com/valyala/fasthttp"

	"github.com/yunginnanet/HellPot/config"
)

func robotsTXT(ctx *fasthttp.RequestCtx) {
	paths := &strings.Builder{}
	paths.WriteString("User-agent: *\r\n")
	for _, p := range config.Paths {
		paths.WriteString("Disallow: ")
		paths.WriteString(p)
		paths.WriteString("\r\n")
	}
	paths.WriteString("\r\n")

	log.Debug().
		Strs("PATHS", config.Paths).
		Msg("SERVE_ROBOTS")

	if _, err := fmt.Fprintf(ctx, paths.String()); err != nil {
		log.Error().Err(err).Msg("SERVE_ROBOTS_ERROR")
	}
}
