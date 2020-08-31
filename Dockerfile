FROM golang:1.15-alpine AS builder
WORKDIR /go/src/git.sr.ht/~ancarda/tls-redirector
COPY . ./
RUN apk add git binutils
RUN go build .
RUN strip tls-redirector

FROM alpine:latest
COPY --from=builder /go/src/git.sr.ht/~ancarda/tls-redirector/tls-redirector .
EXPOSE 80
CMD ["./tls-redirector"]
