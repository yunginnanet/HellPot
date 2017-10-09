package main

import (
	"context"
	"flag"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/carlmjohnson/heffalump/heff"
)

const usage = `Usage of heffalump:

heffalump [opts]

	heffalump serves an endless HTTP honeypot

`

const robotsTxt = "User-agent: *\r\nDisallow: /\r\n"

func main() {
	flag.Usage = func() {
		fmt.Fprintf(os.Stderr, usage)
		flag.PrintDefaults()
	}
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	// subscribe to SIGINT signals
	stopChan := make(chan os.Signal, 1)
	signal.Notify(stopChan, syscall.SIGINT, syscall.SIGTERM)

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/" {
			http.NotFound(w, r)
			return
		}
		heff.DefaultHoneypot(w, r)
	})

	http.HandleFunc("/robots.txt", func(w http.ResponseWriter, r *http.Request) {
		if _, err := io.WriteString(w, robotsTxt); err != nil {
			log.Printf("error serving robots.txt: %v", err)
		}
	})

	srv := &http.Server{Addr: ":" + port, Handler: http.DefaultServeMux}

	go func() {
		// service connections
		err := srv.ListenAndServe()
		log.Printf("Finished listening: %v\n", err)
	}()

	<-stopChan // wait for SIGINT
	log.Println("Shutting down server...")

	// shut down gracefully, but wait no longer than 5 seconds before halting
	ctx, c := context.WithTimeout(context.Background(), 5*time.Second)
	defer c()
	srv.Shutdown(ctx)

	log.Println("Server gracefully stopped")
}
