FROM golang:1.18.1-alpine3.15 AS builder
WORKDIR /go/src/git.sr.ht/~ancarda/tls-redirector
RUN apk add binutils git
COPY . ./
RUN go build . && strip tls-redirector

FROM alpine:3.15.4
RUN apk add curl
COPY --from=builder /go/src/git.sr.ht/~ancarda/tls-redirector/tls-redirector .
EXPOSE 80
CMD ["./tls-redirector"]
HEALTHCHECK --interval=10s CMD curl --fail http://localhost:$PORT/ || exit 1
