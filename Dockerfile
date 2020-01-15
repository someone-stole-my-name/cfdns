FROM golang:alpine

WORKDIR /go/src/app

RUN apk add git dep

RUN git clone https://github.com/someone-stole-my-name/cfdns.git .
RUN dep ensure && go build *.go

ENTRYPOINT [ "/go/src/app/cfdns" ]

