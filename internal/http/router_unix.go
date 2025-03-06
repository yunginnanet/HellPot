//go:build linux || darwin || freebsd

package http

import (
	"net"
	"os"
	"syscall"

	"github.com/fasthttp/router"
)

func listenOnUnixSocket(addr string, r *router.Router) (net.Listener, error) {
	config := runningConfig.HTTP.UnixSocket
	var err error
	var unixAddr *net.UnixAddr
	var unixListener *net.UnixListener
	unixAddr, err = net.ResolveUnixAddr("unix", addr)
	if err != nil {
		return nil, err
	}
	// Always unlink sockets before listening on them
	_ = syscall.Unlink(addr)
	// Before we set socket permissions, we want to make sure only the user HellPot is running under
	// has permission to the socket.
	oldmask := syscall.Umask(0o077)
	unixListener, err = net.ListenUnix("unix", unixAddr)
	syscall.Umask(oldmask)
	if err != nil {
		return nil, err
	}
	if err = os.Chmod(
		unixAddr.Name,
		os.FileMode(config.UnixSocketPermissions),
	); err != nil {
		return nil, err
	}

	return unixListener, nil
}
