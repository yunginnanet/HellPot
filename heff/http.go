package heff

import "net/http"

func Honeypot(w http.ResponseWriter, r *http.Request) {
	for {
		Generate(w, 10000)
	}
}
