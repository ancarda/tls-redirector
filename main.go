package main

import (
	"log"
	"net/http"
	"os"

	"git.sr.ht/~ancarda/tls-redirector/socketactivation"
)

const (
	acmeChallengeURLPrefix = "/.well-known/acme-challenge/"
	version                = "2.2"
	defaultPort            = "80"
)

func listenTCP(portNumber string) {
	log.Fatal(http.ListenAndServe(":"+portNumber, nil))
}

func main() {
	if len(os.Args) > 1 {
		handleCliArgs()
	}

	acmeChallengeDir := os.Getenv("ACME_CHALLENGE_DIR")
	if acmeChallengeDir != "" {
		if _, err := os.Stat(acmeChallengeDir); os.IsNotExist(err) {
			log.Fatalf("fatal: ACME HTTP challenge directory not found: %s",
				acmeChallengeDir)
		}
	}

	// Setup the server
	http.Handle("/", newApp(acmeChallengeDir))

	// If PORT is specified, take that as authoritative
	port := os.Getenv("PORT")
	if port == "systemd" {
		log.Fatal(socketactivation.Serve())
	}

	if port != "" {
		listenTCP(port)
	}

	// Try to listen using systemd socket activation
	if socketactivation.Enabled {
		switch socketactivation.CountListeners() {
		case 0:
			listenTCP(defaultPort) // fallback

		case 1:
			log.Fatal(socketactivation.Serve())

		default:
			log.Fatal("systemd socket activation - pass zero or one sockets")
		}
	}

	// Default to listening on port 80
	listenTCP(defaultPort)
}
