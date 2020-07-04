FROM amd64/golang:rc-alpine3.12

WORKDIR /go/src/app
RUN apk add git dep
RUN git clone https://github.com/someone-stole-my-name/cfdns.git .
RUN cd src/cfdns && dep ensure && go build *.go
COPY Entrypoint /Entrypoint
ENTRYPOINT ["/Entrypoint"]
