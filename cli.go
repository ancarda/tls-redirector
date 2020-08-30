package main

import (
	"fmt"
	"os"

	"git.sr.ht/~ancarda/tls-redirector/socketactivation"
)

func handleCliArgs() {
	switch os.Args[1] {
	case "--version":
		fmt.Println("tls redirector", version)

	case "--help":
		fmt.Println("Usage:", os.Args[0], "[OPTION]")
		fmt.Println("A tiny service for port 80 that rewrites URLs to HTTPS.")
		fmt.Println("")
		if socketactivation.Enabled {
			fmt.Println("Compiled with support for systemd socket activation.")
		} else {
			fmt.Println("Compiled without support for systemd socket activation.")
		}
		fmt.Println("")
		fmt.Println("Options")
		fmt.Println("    --help     display this help and exit")
		fmt.Println("    --version  output version information and exit")
		fmt.Println("")
		fmt.Println("Environmental Variables")
		fmt.Println("    PORT               - TCP/IP port number to listen on.")
		fmt.Println("    ACME_CHALLENGE_DIR - Directory to serve at", acmeChallengeURLPrefix)
		fmt.Println("")
		fmt.Println("For more help, please refer to the README. Available online:")
		fmt.Println("    https://git.sr.ht/~ancarda/tls-redirector/tree/master/README.md")

	default:
		fmt.Println("unrecognized:", os.Args[1])
		os.Exit(1)
	}

	os.Exit(0)
}
