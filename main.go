// Package main implements tls-redirector, a microservice that
// is designed to forward all HTTP traffic to HTTPS.
package main

import (
	"log"
	"net"
	"net/http"

	"github.com/coreos/go-systemd/activation"
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

func handle(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Server", "tls-redirector/1.0")

	// If we haven't been given a host, just abort.
	if r.Host == "" {
		http.Error(w, "tls-redirector cannot handle this request because no `host` header was sent by your browser.", http.StatusBadRequest)
		return
	}

	// If we've seen a direct request, such as an IP address,
	// we'll discard it now. Not only is it uncommon to see
	// X.509 certificates for IP addresses, it's very likely
	// to be junk traffic (scanning bots). We can save the
	// real web server some effort by dropping it now.
	if isIPAddress(r.Host) {
		http.Error(w, "tls-redirector cannot redirect IP addresses.", http.StatusBadRequest)
		return
	}

	// Overwrite the scheme to https:// and redirect.
	// Change host as well as in r.URL, it's empty.
	r.URL.Host = r.Host
	r.URL.Scheme = "https"
	http.Redirect(w, r, r.URL.String(), http.StatusMovedPermanently)
}

func main() {
	http.HandleFunc("/", handle)

	listeners, err := activation.Listeners()
	log.Println(listeners)
	if err != nil {
		log.Panicf("cannot retrieve listeners: %s", err)
	}

	if len(listeners) != 1 {
		log.Panicf("unexpected number of socket activation (%d != 1)", len(listeners))
	}

	log.Fatal(http.Serve(listeners[0], nil))
}
