FROM golang:1.17.1-alpine3.14 AS builder
WORKDIR /go/src/git.sr.ht/~ancarda/tls-redirector
RUN apk add binutils
COPY . ./
RUN go build . && strip tls-redirector

FROM alpine:3.14.2
RUN apk add curl
COPY --from=builder /go/src/git.sr.ht/~ancarda/tls-redirector/tls-redirector .
EXPOSE 80
CMD ["./tls-redirector"]
HEALTHCHECK --interval=10s CMD curl --fail http://localhost || exit 1
