// +build !systemd

package socketactivation

const enabled bool = false

func countListeners() int {
	return 0
}

func serve() error {
	return ErrNotEnabled
}
