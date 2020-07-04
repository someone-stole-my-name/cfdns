FROM alpine AS builder

ENV QEMU_URL https://github.com/balena-io/qemu/releases/download/v3.0.0%2Bresin/qemu-3.0.0+resin-aarch64.tar.gz
RUN apk add curl && curl -L ${QEMU_URL} | tar zxvf - -C . --strip-components 1

FROM arm64v8/golang:rc-alpine3.12

COPY --from=builder qemu-aarch64-static /usr/bin
WORKDIR /go/src/app
RUN apk add git dep
RUN git clone https://github.com/someone-stole-my-name/cfdns.git .
RUN cd src/cfdns && dep ensure && go build *.go
COPY Entrypoint /Entrypoint
ENTRYPOINT ["/Entrypoint"]
