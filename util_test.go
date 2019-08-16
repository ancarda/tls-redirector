package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIsIPAddress_IPv4(t *testing.T) {
	assert.True(t, isIPAddress("127.0.0.1"), "IPv4 address")
	assert.True(t, isIPAddress("127.0.0.1:8080"), "IPv4 address with port")
}

func TestIsIPAddress_IPv6(t *testing.T) {
	assert.True(t, isIPAddress("::1"), "IPv6 address")
	assert.True(t, isIPAddress("[::1]:8080"), "IPv6 address with port")
}

func TestIsIPAddress_InvalidIPAddresses(t *testing.T) {
	assert.False(t, isIPAddress("1.2.3.4.5"))
	assert.False(t, isIPAddress("256.1.1.1"))
}

func TestIsIPAddress_RejectOtherStrings(t *testing.T) {
	assert.False(t, isIPAddress("example.com"))
	assert.False(t, isIPAddress(""))
}
