FROM golang:1.14 AS builder
WORKDIR /src
COPY go.mod go.sum Makefile .git ./
COPY cmd cmd
COPY internal internal
COPY pkg pkg
RUN go mod download
RUN make install

FROM alpine:latest
RUN apk add bash
COPY --from=builder /go/bin/ach /usr/local/bin/
# https://stackoverflow.com/questions/34729748/installed-go-binary-not-found-in-path-on-alpine-linux-docker
RUN mkdir /lib64 && ln -s /lib/libc.musl-x86_64.so.1 /lib64/ld-linux-x86-64.so.2
COPY integration/sampleConfigDir /root/.ach
COPY integration/integration.sh integration.sh

RUN ach version
ENTRYPOINT [ "sh", "integration.sh" ]
