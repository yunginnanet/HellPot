package http

import (
	"fmt"

	"github.com/valyala/fasthttp"

	"github.com/yunginnanet/HellPot/config"
)

const robotsTxt = "User-agent: *\r\n"

func robotsTXT(ctx *fasthttp.RequestCtx) {
	var paths string
	for _, p := range config.Paths {
		paths = paths + "Disallow: " + p + "\r\n"
	}

	log.Debug().
		Strs("PATHS", config.Paths).
		Msg("SERVE_ROBOTS")

	if _, err := fmt.Fprintf(ctx, robotsTxt+paths+"\r\n"); err != nil {
		log.Error().Err(err).Msg("SERVE_ROBOTS_ERROR")
	}
}
