package heff

import (
	"io"
	"log"
	"net/http"
)

func Honeypot(w http.ResponseWriter, r *http.Request) {
	log.Printf("Started writing: %v", r.URL)
	buf := make([]byte, 100*1<<10)
	io.WriteString(w, "<HTML>\n<BODY>\n")
	n, err := io.CopyBuffer(w, DefaultMarkovMap, buf)
	log.Printf("Wrote: %d (%v)", n, err)
}
