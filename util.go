package main

import (
	"net"
	"strings"
)

func removeSquareBrackets(s string) string {
	if !strings.HasPrefix(s, "[") {
		return s
	}

	if !strings.HasSuffix(s, "]") {
		return s
	}

	return s[1 : len(s)-1]
}

func isIPAddress(host string) bool {
	// Addresses without a port number attached can be checked as-is
	if net.ParseIP(removeSquareBrackets(host)) != nil {
		return true
	}

	// Try to strip the port number off, then check again.
	if host, _, err := net.SplitHostPort(host); err == nil {
		if net.ParseIP(removeSquareBrackets(host)) != nil {
			return true
		}
	}

	return false
}
