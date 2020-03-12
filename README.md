# TLS Redirector

[![License](https://img.shields.io/github/license/ancarda/tls-redirector.svg)](https://choosealicense.com/licenses/agpl-3.0/)
[![Build Status](https://travis-ci.org/ancarda/tls-redirector.svg?branch=master)](https://travis-ci.org/ancarda/tls-redirector)
[![Go Report Card](https://goreportcard.com/badge/github.com/ancarda/tls-redirector)](https://goreportcard.com/report/github.com/ancarda/tls-redirector)

tls-redirector is a tiny HTTP server written in Go that is designed to run on
port 80 and redirect all incoming traffic to HTTPS. It does this by emitting a
301 Permanent Redirect where the scheme is simply replaced with "https".

This is intended to separate out responsibilities from the software that
listens on port 443 and serves your website. Because you have tls-redirector
on port 80, you do not need to configure HTTP to HTTPS redirects; simplifying
configuration for Apache, nginx, or whatever web server you use.

Furthermore, most crawlers and bots will actually connect to port 80 without a
meaningful Host header. As tls-redirector cannot be configured, it can only
politely tell them to go away, whereas your primary web server likely sends
them to the "default server" which may well redirect them to your website,
opening you up to being scanned. This causes tons of noise in your log files.

Because it uses the Host header, tls-redirector is truly zero-config. Set it
up once and forget about it.

## Useful Links

* Source Code:   <https://git.sr.ht/~ancarda/tls-redirector/>
* Issue Tracker: <https://todo.sr.ht/~ancarda/tls-redirector/>
* Mailing List:  <https://lists.sr.ht/~ancarda/tls-redirector/>

## Features

* Can run as an unprivileged user by use of systemd activation sockets.
* Possible to run without any disk access, making sandboxing trivial.
* IP address traffic (usually by crawlers) is dropped.
* Can serve your .well-known/acme-challenge directory.

## Possible Caveats

* The host field is required to redirect (there's no way to configure a
  default), and visitors without one will simply see an error message.
* Must be running on a Linux distribution with systemd for now.
