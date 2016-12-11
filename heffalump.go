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

heffalump [<network address> [<path>]]

	heffalump serves an endless HTTP honeypot

	<network address> defaults to ":8080".

	<path> defaults to "/". Paths ending in "/" will match all sub-pathes.
`

func main() {
	flag.Usage = func() {
		fmt.Fprintf(os.Stderr, usage)
	}
	flag.Parse()

	addr := flag.Arg(0)
	if addr == "" {
		addr = ":8080"
	}

	path := flag.Arg(1)
	if path == "" {
		path = "/"
	}

	http.HandleFunc(path, heff.Honeypot)

	log.Fatal(http.ListenAndServe(addr, nil))

}
