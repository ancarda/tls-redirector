package main

import (
	"net"
)

func isIPAddress(host string) bool {
	if net.ParseIP(host) != nil {
		return true
	}

	if host, _, err := net.SplitHostPort(host); err == nil {
		if net.ParseIP(host) != nil {
			return true
		}
	}

	return false
}
