// +build !systemd

package main

import (
	"fmt"
)

func systemdEnabled() bool {
	return false
}

func systemdCanServe() bool {
	return false
}

func systemdServe() error {
	return fmt.Errorf("socket activation was not enabled at compile time.")
}
