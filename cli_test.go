package main

import (
	"bytes"
	"regexp"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCLIVersionFlag(t *testing.T) {
	var buf bytes.Buffer
	args := []string{"./test", "--version"}

	assert.Equal(t, 0, handleCliArgs(&buf, args, false))
	assert.Equal(t, "tls redirector 2.2\n", buf.String())
}

func TestCLIUnknownFlag(t *testing.T) {
	var buf bytes.Buffer
	args := []string{"./test", "--no-such-flag"}

	assert.Equal(t, 1, handleCliArgs(&buf, args, false))
	assert.Equal(t, "unrecognized: --no-such-flag\n", buf.String())
}

func TestCLIHelpFlag(t *testing.T) {
	var buf bytes.Buffer
	args := []string{"./test", "--help"}

	assert.Equal(t, 0, handleCliArgs(&buf, args, false))

	strings := []string{
		"Usage: ./test [OPTION]\n",
		"Compiled without support for systemd socket activation.\n",
	}
	for _, s := range strings {
		assert.Contains(t, buf.String(), s)
	}

	regexes := []string{
		"--help\\s+display this help and exit",
		"--version\\s+output version information and exit",
		"PORT\\s+- TCP/IP port number to listen on",
		"ACME_CHALLENGE_DIR\\s+- Directory to serve at /.well-known/acme-challenge",
	}
	for _, r := range regexes {
		assert.Regexp(t, regexp.MustCompile(r), buf.String())
	}
	assert.Regexp(t, regexp.MustCompile("--help\\s+display this help and exit"), buf.String())
}

func TestCLIHelpFlag_WithSocketActivation(t *testing.T) {
	var buf bytes.Buffer
	args := []string{"./test", "--help"}

	assert.Equal(t, 0, handleCliArgs(&buf, args, true))
	assert.Contains(t, buf.String(), "Compiled with support for systemd socket activation.")
}
