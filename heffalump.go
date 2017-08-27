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

	path := flag.String("path", "/", `Path to serve from. Path ending in / serves sub-paths.`)
	flag.Parse()

	http.HandleFunc(*path, heff.DefaultHoneypot)

	log.Fatal(http.ListenAndServe(":"+port, nil))

}
