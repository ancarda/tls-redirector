// Package socketactivation provides conditional support for systemd socket
// activation.
//
// By default, this package is disabled, and all functions are merely dummies.
// To enable, build like so:
//
//     go build -tags systemd
//
package socketactivation
