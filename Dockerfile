FROM golang:latest
MAINTAINER Silver "lolooo@live.com"

ENV GOROOT /usr/local/go
ENV GOPATH /Users/luoliang/projects/go
ENV PATH $GOROOT/bin:$PATH

WORKDIR $GOPATH/src/github.com/wuleying/silver-framework
RUN pwd

ADD . $GOPATH/src/github.com/wuleying/silver-framework
RUN go build .

ENTRYPOINT  ["./silver-framework"]