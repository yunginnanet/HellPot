# Heffalump [![GoDoc](https://godoc.org/github.com/carlmjohnson/heffalump?status.svg)](https://godoc.org/github.com/carlmjohnson/heffalump) [![Go Report Card](https://goreportcard.com/badge/github.com/carlmjohnson/heffalump)](https://goreportcard.com/report/github.com/carlmjohnson/heffalump)
Heffalump is an endless honeypot that gives malicious bots nightmares. To use, in your robots.txt tell robots not to go to a certain URL, which heffalump is reverse proxying. Any web agent that does go to the URL will receive an endless stream of random data, which will overflow its memory and/or storage if it doesn't have a max buffer size set or at the very least severely waste its time.

The source of the honeypot data is [Once On a Time](http://www.gutenberg.org/files/27771/27771-h/27771-h.htm), one of A. A. Milne's most beloved and most public domain works.

![Exploding Heffalump](exploding-heffalump.gif)

## Installation
First install [Go](http://golang.org).

If you just want to install the binary to your current directory and don't care about the source code, run

```shell
GOBIN=$(pwd) GOPATH=$(mktemp -d) go get github.com/carlmjohnson/heffalump
```

## Usage
Usage of heffalump:

heffalump [opts]

    heffalump serves an endless HTTP honeypot

  -addr string
        Address to serve (default "127.0.0.1:8080")
  -path string
        Path to serve from. Path ending in / serves sub-paths. (default "/")
