package main

import (
	"fmt"
	"io"
)

const (
	ExitSuccess = 0
	ExitFailure = 1
)

func handleCliArgs(w io.Writer, args []string, hasSocketActivation bool) int {
	switch args[1] {
	case "--version":
		fmt.Fprintln(w, "tls redirector", version)

	case "--help":
		fmt.Fprintln(w, "Usage:", args[0], "[OPTION]")
		fmt.Fprintln(w, "A tiny service for port 80 that rewrites URLs to HTTPS.")
		fmt.Fprintln(w, "")
		if hasSocketActivation {
			fmt.Fprintln(w, "Compiled with support for systemd socket activation.")
		} else {
			fmt.Fprintln(w, "Compiled without support for systemd socket activation.")
		}
		fmt.Fprintln(w, "")
		fmt.Fprintln(w, "Options")
		fmt.Fprintln(w, "    --help     display this help and exit")
		fmt.Fprintln(w, "    --version  output version information and exit")
		fmt.Fprintln(w, "")
		fmt.Fprintln(w, "Environmental Variables")
		fmt.Fprintln(w, "    PORT               - TCP/IP port number to listen on.")
		fmt.Fprintln(w, "    ACME_CHALLENGE_DIR - Directory to serve at", acmeChallengeURLPrefix)
		fmt.Fprintln(w, "")
		fmt.Fprintln(w, "For more help, please refer to the README. Available online:")
		fmt.Fprintln(w, "    https://git.sr.ht/~ancarda/tls-redirector/tree/master/README.md")

	default:
		fmt.Fprintln(w, "unrecognized:", args[1])
		return ExitFailure
	}

	return ExitSuccess
}
