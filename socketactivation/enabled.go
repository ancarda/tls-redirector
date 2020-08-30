// +build systemd

package socketactivation

import (
	"fmt"
	"net"
	"net/http"

	"github.com/coreos/go-systemd/activation"
)

const (
	enabled              bool = true
	maximumListenerCount int  = 1
)

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

	if len(listeners) != maximumListenerCount {
		return fmt.Errorf("expected exactly %d listener, have %d",
			maximumListenerCount, len(listeners))
	}

	return http.Serve(listeners[0], nil)
}
