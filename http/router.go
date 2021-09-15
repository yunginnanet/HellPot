package http

import (
	"time"

	"github.com/fasthttp/router"
	"github.com/valyala/fasthttp"

	"github.com/yunginnanet/HellPot/config"
	"github.com/yunginnanet/HellPot/heffalump"
)

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

// Serve starts our HTTP server and request router
func Serve() error {
	r := router.New()
	r.GET("/robots.txt", robotsTXT)
	r.GET("/{path}", scopeCheck)
	return fasthttp.ListenAndServe(":8080", r.Handler)
}
