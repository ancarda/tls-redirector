# TLS Redirector

[![License](https://img.shields.io/badge/license-AGPL--3.0--or--later-teal)](https://choosealicense.com/licenses/agpl-3.0/)
[![builds.sr.ht status](https://builds.sr.ht/~ancarda/tls-redirector.svg)](https://builds.sr.ht/~ancarda/tls-redirector)
[![Go Report Card](https://goreportcard.com/badge/github.com/ancarda/tls-redirector)](https://goreportcard.com/report/github.com/ancarda/tls-redirector)
[![Docker](https://img.shields.io/docker/pulls/ancarda/tls-redirector)](https://hub.docker.com/r/ancarda/tls-redirector)

TLS Redirector is a tiny HTTP server written in Go that is designed to run on
port 80 and redirect all incoming traffic to HTTPS. It does this by emitting a
301 Permanent Redirect where the scheme is simply replaced with "https".

This is intended to separate out responsibilities from the software that
listens on port 443 and serves your website. Because you have TLS Redirector
on port 80, you do not need to configure HTTP to HTTPS redirects; simplifying
configuration for Apache, nginx, or whatever web server you use.

Furthermore, most crawlers and bots will actually connect to port 80 without a
meaningful Host header. As TLS Redirector cannot be configured, it can only
politely tell them to go away, whereas your primary web server likely sends
them to the "default server" which may well redirect them to your website,
opening you up to being scanned. This causes tons of noise in your log files.

Because it uses the Host header, TLS Redirector is truly zero-config. Set it
up once and forget about it.

## Useful Links

* Source Code:   <https://git.sr.ht/~ancarda/tls-redirector/>
* Issue Tracker: <https://todo.sr.ht/~ancarda/tls-redirector/>
* Mailing List:  <https://lists.sr.ht/~ancarda/tls-redirector/>

## Features

* Can run as an unprivileged user by use of systemd activation sockets.
* Possible to run without any disk access, making sandboxing trivial.
* IP address traffic (usually by crawlers) is dropped.
* Can serve your `.well-known/acme-challenge` directory.

## Quick Run

    docker run -d --user nobody --restart=unless-stopped --env PORT=1234 \
        -p 80:1234 --name tls-redirector ancarda/tls-redirector

## How To Build

If you want systemd socket activation, you need to compile this way:

    go build -tags systemd

Building without that tag will produce a binary that only has TCP/IP support.

## Possible Caveats

* The `Host` header is required to redirect as there's no way to configure a
  default. Visitors without one will simply see an error message. This message
  is specified in the program source code and cannot be configured at runtime.

* Only a single ACME challenge directory can be served, as the `Host` header
  is ignored, so if you have multiple domains or servers on the same machine,
  you may want to:
    * Consider using DNS based ACME challenges.
    * Store all your HTTP based ACME challenges in the same directory.

## Configuration

Behavior may be configured through the following environmental variables:

* `PORT`. TCP/IP port number to listen on. If not specified, port 80 is used
  OR systemd socket activation is detected and used.

  You can force TLS Redirector to use socket activation with `PORT=systemd`.

* `ACME_CHALLENGE_DIR` Path to a directory on disk to serve at the
  path `/.well-known/acme-challenge`. All files are served as `text/plain` and
  is intended to provide support for HTTP based ACME challenges if necessary.

  Setting this to `/tmp` means TLS Redirector will look for files in that
  directory. This differs slightly from the `--webroot` command of EFF Certbot
  because Certbot expects to be at the root, and TLS Redirector does not.
  Therefore, when you setup Certbot, if you have the following:

  `certbot --webroot /var/www`

  You should set TLS Redirector to:

  `ACME_CHALLENGE_DIR=/var/www/.well-known/acme-challenge`
