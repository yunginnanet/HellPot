package heff

import (
	"io"
	"log"
	"net/http"
	"sync"
)

var DefaultHoneypot = NewHoneypot(DefaultMarkovMap, 100*1<<10)

func NewHoneypot(mm MarkovMap, buffsize int) http.HandlerFunc {
	var pool sync.Pool

	getBuffer := func() []byte {
		x := pool.Get()
		if buf, ok := x.([]byte); ok {
			return buf
		} else {
			return make([]byte, buffsize)
		}
	}

	putBuffer := func(buf []byte) {
		pool.Put(buf)
	}

	return func(w http.ResponseWriter, r *http.Request) {
		log.Printf("Started writing: %v", r.URL)
		buf := getBuffer()
		defer putBuffer(buf)
		io.WriteString(w, "<HTML>\n<BODY>\n")
		n, err := io.CopyBuffer(w, DefaultMarkovMap, buf)
		log.Printf("Wrote: %d (%v)", n, err)
	}
}
