package socketactivation

const (
	// Enabled allows consumers to check if support was enabled at compile
	// time.
	Enabled bool = enabled

	// ErrNotEnabled is returned when a caller attempts to call a function
	// that requires support for socket activation.
	ErrNotEnabled = err("socket activation was not enabled at compile time")
)

// CountListeners determines how many sockets have been passed.
//
// When Enabled = false, this function always returns zero (0).
func CountListeners() int {
	return countListeners()
}

// Serve will attempt to serve HTTP on the passed in socket.
//
// When Enabled = false, this function always returns ErrNotEnabled.
func Serve() error {
	return serve()
}
