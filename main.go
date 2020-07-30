// Package main implements tls-redirector, a microservice that
// is designed to forward all HTTP traffic to HTTPS.
package main

import (
	"log"
	"net/http"
	"os"
)

const (
	acmeChallengeUrlPrefix = "/.well-known/acme-challenge/"
	version                = "2.1"
)

var acmeChallengeDir string

func listenTCP(portNumber string) {
	log.Fatal(http.ListenAndServe(":"+portNumber, nil))
}

func main() {
	if len(os.Args) > 1 {
		handleCliArgs()
	}

	acmeChallengeDir = os.Getenv("ACME_CHALLENGE_DIR")
	if acmeChallengeDir != "" {
		if _, err := os.Stat(acmeChallengeDir); os.IsNotExist(err) {
			log.Fatalf("fatal: ACME HTTP challenge directory not found: %s",
				acmeChallengeDir)
		}
	}

	// If PORT is specified, take that as authoritative
	port := os.Getenv("PORT")
	if port == "systemd" {
		log.Fatal(systemdServe())
	}

	if port != "" {
		listenTCP(port)
	}

	// Try to listen using systemd socket activation
	if systemdEnabled() {
		switch systemdCountListeners() {
		case 0:
			listenTCP("80") // fallback

		case 1:
			log.Fatal(systemdServe())

		default:
			log.Fatal("systemd socket activation - pass zero or one sockets")
		}
	}

	// Default to listening on port 80
	listenTCP("80")
}
