package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/carlmjohnson/heffalump/heff"
)

const usage = `Usage of heffalump:

heffalump [opts]

	heffalump serves an endless HTTP honeypot

`

func main() {
	flag.Usage = func() {
		fmt.Fprintf(os.Stderr, usage)
		flag.PrintDefaults()
	}
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/" {
			http.NotFound(w, r)
			return
		}
		heff.DefaultHoneypot(w, r)
	})

	log.Fatal(http.ListenAndServe(":"+port, nil))

}
