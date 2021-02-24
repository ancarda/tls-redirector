FROM golang:1.16.0-alpine3.13 AS builder
WORKDIR /go/src/git.sr.ht/~ancarda/tls-redirector
RUN apk add git binutils
COPY go.* ./
RUN go mod download
COPY . ./
RUN go build . && strip tls-redirector

FROM alpine:3.12.3
RUN apk add curl
COPY --from=builder /go/src/git.sr.ht/~ancarda/tls-redirector/tls-redirector .
EXPOSE 80
CMD ["./tls-redirector"]
HEALTHCHECK --interval=10s CMD curl --fail http://localhost || exit 1
