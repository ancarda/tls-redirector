FROM golang:1.15.1-alpine3.12 AS builder
WORKDIR /go/src/git.sr.ht/~ancarda/tls-redirector
RUN apk add git binutils
COPY go.* ./
RUN go mod download
COPY . ./
RUN go generate ./... && go build . && strip tls-redirector

FROM alpine:3.12.0
COPY --from=builder /go/src/git.sr.ht/~ancarda/tls-redirector/tls-redirector .
EXPOSE 80
CMD ["./tls-redirector"]
