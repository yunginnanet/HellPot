package main

import (
	"HellPot/src/config"
	"github.com/gorilla/mux"
	"io"
	"net/http"
	"sync"
	"time"
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
		var inscope bool = false

		if vars["path"] == "robots.txt" {
			var paths string
			for _, p := range config.Paths {
				paths = paths + "Disallow: " + p + "\r\n"
			}

			log.Debug().
				Str("UserAgent", r.UserAgent()).
				Strs("REMOTE_ADDR", r.Header.Values("X-Real-IP")).
				Strs("PATHS", config.Paths).
				Msg("SERVE_ROBOTS")

			if _, err := io.WriteString(w, robotsTxt+paths+"\r\n"); err != nil {
				log.Error().Err(err).Msg("SERVE_ROBOTS_ERROR")
			}
			return
		}

		for _, p := range config.Paths {
			if vars["path"] == p {
				inscope = true
			}
		}

		if !inscope {
			log.Warn().
				Str("UserAgent", r.UserAgent()).
				Str("URL", r.URL.RequestURI()).
				Strs("REMOTE_ADDR", r.Header.Values("X-Real-IP")).
				Msg("Request outside of configured scope!")
			return
		}

		s := time.Now()
		log.Info().
			Str("UserAgent", r.UserAgent()).
			Interface("URL", r.URL.RequestURI()).
			Strs("REMOTE_ADDR", r.Header.Values("X-Real-IP")).
			Msg("SERVE")
		buf := getBuffer()
		defer putBuffer(buf)
		io.WriteString(w, "<HTML>\n<BODY>\n")
		n, _ := io.CopyBuffer(w, mm, buf)
		log.Info().
			Str("UserAgent", r.UserAgent()).
			Interface("URL", r.URL.RequestURI()).
			Strs("REMOTE_ADDR", r.Header.Values("X-Real-IP")).
			Int64("BYTES", n).
			Dur("DURATION", time.Since(s)).
			Msg("FINISH")
	}
}
