package http

import (
	"io"
	"time"

	"github.com/yunginnanet/HellPot/config"
	"github.com/yunginnanet/HellPot/heffalump"
)

const robotsTxt = "User-agent: *\r\n"

func robotsTXT(w io.Writer) {
	var paths string
	for _, p := range config.Paths {
		paths = paths + "Disallow: " + p + "\r\n"
	}

	log.Debug().
		Strs("PATHS", config.Paths).
		Msg("SERVE_ROBOTS")

	if _, err := io.WriteString(w, robotsTxt+paths+"\r\n"); err != nil {
		// log.Error().Err(err).Msg("SERVE_ROBOTS_ERROR")
	}
	return
}

func hellHandle(r io.Reader, w io.Writer) {

	/*
			vars := mux.Vars(r)
		    var inscope = false

		   	var remoteAddr string
		   	remoteAddr = r.RemoteAddr

		   	if len(r.Header.Values("X-Real-IP")) > 0 {
		   		remoteAddr = r.Header.Values("X-Real-IP")[0]
		   	}

		   	log := log.With().
		   		Str("UserAgent", r.UserAgent()).
		   		Str("REMOTE_ADDR", remoteAddr).
		   		Interface("URL", r.URL.RequestURI()).Logger()

		   	if vars["path"] == "robots.txt" {
		   		robotsTXT(w)
		   		return
		   	}

		   	for _, p := range config.Paths {
		   		if vars["path"] == p {
		   			inscope = true
		   		}
		   	}

		   	if !inscope {
		   		// slog.Warn().Msg("Request outside of configured scope!")
		   		return
		   	}
	*/

	start := time.Now()
	// slog.Info().Msg("SERVE")

	n := heffalump.DefaultHeffalump.WriteHell(w)

	// slog.Info().
	log.Info().
		Int64("BYTES", n).
		Dur("DURATION", time.Since(start)).
		Msg("FINISH")
}
