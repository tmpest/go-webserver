FROM golang

RUN go get -u github.com/tmpest/go-webserver

ENV WEBSERVER_ROOT "/go/src/github.com/tmpest/go-webserver"

RUN apk --no-cache add ca-certificates

EXPOSE 80
ENTRYPOINT /go/bin/go-webserver
