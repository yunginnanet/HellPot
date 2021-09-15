package main

import (
	"io"
	"net/http"
	"sync"
	"time"

	"github.com/gorilla/mux"

	"github.com/yunginnanet/HellPot/config"
)

// DefaultHoneypot is an http.HandlerFunc that serves random HTML from the
// DefaultMarkovMap, 100KB at a time.
var DefaultHoneypot = NewHoneypot(DefaultMarkovMap, 100*1<<10)

// NewHoneypot creates an http.HandlerFunc from a MarkovMap
func NewHoneypot(mm MarkovMap, buffsize int) http.HandlerFunc {
	var pool sync.Pool

	getBuffer := func() []byte {
		x := pool.Get()
		if buf, ok := x.([]byte); ok {
			return buf
		}
		return make([]byte, buffsize)
	}

	putBuffer := func(buf []byte) {
		pool.Put(buf)
	}

	return func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		var inscope = false
		var remoteAddr string

		remoteAddr = r.RemoteAddr

		if len(r.Header.Values("X-Real-IP")) > 0 {
			remoteAddr = r.Header.Values("X-Real-IP")[0]
		}

		slog := log.With().
			Str("UserAgent", r.UserAgent()).
			Str("REMOTE_ADDR", remoteAddr).
			Interface("URL", r.URL.RequestURI()).Logger()

		if vars["path"] == "robots.txt" {
			var paths string
			for _, p := range config.Paths {
				paths = paths + "Disallow: " + p + "\r\n"
			}

			slog.Debug().
				Strs("PATHS", config.Paths).
				Msg("SERVE_ROBOTS")

			// Not writing this header broke HellPot in ~go1.17
			w.WriteHeader(200)

			if _, err := io.WriteString(w, robotsTxt+paths+"\r\n"); err != nil {
				slog.Error().Err(err).Msg("SERVE_ROBOTS_ERROR")
			}
			return
		}

		for _, p := range config.Paths {
			if vars["path"] == p {
				inscope = true
			}
		}

		if !inscope {
			w.WriteHeader(404)
			slog.Warn().Msg("Request outside of configured scope!")
			return
		}

		s := time.Now()
		slog.Info().Msg("SERVE")

		buf := getBuffer()
		defer putBuffer(buf)

		if _, err := io.WriteString(w, "<HTML>\n<BODY>\n"); err != nil {
			slog.Debug().Caller().Err(err).Msg("WriteString_fail")
		}

		var n int64
		var err error
		if n, err = io.CopyBuffer(w, mm, buf); err != nil {
			slog.Debug().Caller().Err(err).Msg("CopyBuffer_fail")
		}

		slog.Info().
			Int64("BYTES", n).
			Dur("DURATION", time.Since(s)).
			Msg("FINISH")
	}
}
