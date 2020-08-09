FROM golang

RUN go get -u github.com/tmpest/go-webserver

ENV WEBSERVER_ROOT "/go/src/github.com/tmpest/go-webserver"

RUN apt-get install apt ca-certificates

EXPOSE 80
ENTRYPOINT /go/bin/go-webserver
