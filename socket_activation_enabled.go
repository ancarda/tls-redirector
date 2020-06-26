// +build systemd

package main

import (
	"fmt"
	"net"
	"net/http"

	"github.com/coreos/go-systemd/activation"
)

var (
	saListeners []net.Listener = nil
	saErr       error          = nil
)

func init() {
	saListeners, saErr = activation.Listeners()
}

func systemdEnabled() bool {
	return true
}

func systemdCanServe() bool {
	if saErr == nil && len(saListeners) > 0 {
		return true
	}
	return false
}

func systemdServe() error {
	if saErr != nil {
		return saErr
	}

	if len(saListeners) != 1 {
		return fmt.Errorf("expected exactly 1 socket, have %d", len(saListeners))
	}

	return http.Serve(saListeners[0], nil)
}
