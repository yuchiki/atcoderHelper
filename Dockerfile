FROM golang:1.14 AS builder
WORKDIR /src
COPY go.mod go.sum Makefile .git ./
COPY cmd cmd
COPY internal internal
RUN go mod download
RUN make install

FROM alpine:latest
COPY --from=builder /go/bin/ach /usr/local/bin/
# https://stackoverflow.com/questions/34729748/installed-go-binary-not-found-in-path-on-alpine-linux-docker
RUN mkdir /lib64 && ln -s /lib/libc.musl-x86_64.so.1 /lib64/ld-linux-x86-64.so.2
RUN ach version
