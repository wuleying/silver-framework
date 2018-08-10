FROM golang:latest
MAINTAINER Silver "lolooo@live.com"

WORKDIR $GOPATH/src/github.com/wuleying/silver-framework
RUN pwd

ADD . $GOPATH/src/github.com/wuleying/silver-framework
RUN go build .

ENTRYPOINT  ["./silver-framework"]