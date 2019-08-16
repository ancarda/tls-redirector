# ancarda/tls-redirector

_A tiny service for port 80 that rewrites URLs to HTTPS_

[![License](https://img.shields.io/github/license/ancarda/tls-redirector.svg)](https://choosealicense.com/licenses/mit/)
[![Build Status](https://travis-ci.org/ancarda/tls-redirector.svg?branch=master)](https://travis-ci.org/ancarda/tls-redirector)
[![Go Report Card](https://goreportcard.com/badge/github.com/ancarda/tls-redirector)](https://goreportcard.com/report/github.com/ancarda/tls-redirector)

tls-redirector is a tiny HTTP server that is designed to run on
port 80 and redirect all incoming traffic to HTTPS. It does this
by emitting a 301 Permanent Redirect where the scheme is simply
replaced with "https".

### Features

 * Can run as an unprivileged user by use of systemd activation sockets.
 * Possible to run without any disk access, making sandboxing trivial.
 * IP address traffic (usually by crawlers) is dropped.
 * Can serve your .well-known/acme-challenge directory.

### Possible Caveats

 * The `host` field is required to redirect (there's no way to configure a
   default), and visitors without one will simply see an error message.
 * Must be running on a Linux distribution with systemd for now.
