// +build systemd

package socketactivation

import (
	"net"
	"net/http"

	"github.com/coreos/go-systemd/activation"
)

const enabled = true

var (
	listeners []net.Listener = nil
	initerr   error          = nil
)

func init() {
	listeners, initerr = activation.Listeners()
}

func countListeners() int {
	return len(listeners)
}

func serve() error {
	if initerr != nil {
		return initerr
	}

	return http.Serve(listeners[0], nil)
}
