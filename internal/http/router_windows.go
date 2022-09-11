//go:build windows
// +build windows

package http

import (
	"errors"

	"github.com/fasthttp/router"
)

func listenOnUnixSocket(addr string, r *router.Router) error {
	return errors.New("unix sockets are not supported on Windows")
}
