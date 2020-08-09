FROM golang

RUN go get -u github.com/tmpest/go-webserver

ENV WEBSERVER_ROOT "/go/src/github.com/tmpest/go-webserver"

EXPOSE 8080
ENTRYPOINT /go/bin/go-webserver
