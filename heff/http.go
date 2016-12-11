package heff

import (
	"io"
	"log"
	"net/http"
	"sync"
)

var pool sync.Pool

func getBuffer() []byte {
	x := pool.Get()
	if buf, ok := x.([]byte); ok {
		return buf
	} else {
		return make([]byte, 100*1<<10)
	}
}

func putBuffer(buf []byte) {
	pool.Put(buf)
}

func Honeypot(w http.ResponseWriter, r *http.Request) {
	log.Printf("Started writing: %v", r.URL)
	buf := getBuffer()
	defer putBuffer(buf)
	io.WriteString(w, "<HTML>\n<BODY>\n")
	n, err := io.CopyBuffer(w, DefaultMarkovMap, buf)
	log.Printf("Wrote: %d (%v)", n, err)
}
