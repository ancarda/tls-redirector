package main

import (
	"fmt"
	"html"
	"io/ioutil"
	"net/http"
	"os"
	"strings"

	"github.com/spf13/afero"
)

type app struct {
	fs               afero.Fs
	acmeChallengeDir string
}

func newApp(acd string) app {
	return app{afero.NewOsFs(), acd}
}

func (a app) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Server", "tls-redirector/"+version)
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	w.Header().Set("X-Content-Type-Options", "nosniff")

	// If we haven't been given a host, just abort.
	if r.Host == "" {
		w.WriteHeader(http.StatusBadRequest)
		w.Write(errorPage(
			"Bad Request",
			genericMessage,
			emptyHostHeader,
			version,
		))
		return
	}

	// If we've seen a direct request, such as an IP address,
	// we'll discard it now. Not only is it uncommon to see
	// X.509 certificates for IP addresses, it's very likely
	// to be junk traffic (scanning bots). We can save the
	// real web server some effort by dropping it now.
	if isIPAddress(r.Host) {
		w.WriteHeader(http.StatusBadRequest)
		w.Write(errorPage(
			"400 Bad Request",
			genericMessage,
			fmt.Sprintf(
				hostHeaderIsIPTechInfo,
				html.EscapeString(r.Host),
			),
			version,
		))
		return
	}

	// If we are serving the ACME HTTP challenges, handle that here.
	if a.acmeChallengeDir != "" {
		if strings.HasPrefix(r.URL.Path, acmeChallengeURLPrefix) {
			id := strings.TrimPrefix(r.URL.Path,
				acmeChallengeURLPrefix)
			b, err := readFile(a.fs,
				a.acmeChallengeDir+string(os.PathSeparator)+id)
			if err != nil {
				w.WriteHeader(http.StatusNotFound)
				w.Write(errorPage(
					"404 File Not Found",
					acme404Message,
					fmt.Sprintf(acme404TI,
						html.EscapeString(r.URL.Path)),
					version,
				))
				return
			}

			w.Header().Set("Content-Type", "text/plain")
			w.Write(b)
			return
		}
	}

	// Overwrite the scheme to https:// and redirect.
	// Change host as well as in r.URL, it's empty.
	r.URL.Host = r.Host
	r.URL.Scheme = "https"
	http.Redirect(w, r, r.URL.String(), http.StatusMovedPermanently)
}

func readFile(fs afero.Fs, filename string) ([]byte, error) {
	f, err := fs.Open(filename)
	if err != nil {
		return nil, err
	}
	defer f.Close()

	return ioutil.ReadAll(f)
}
