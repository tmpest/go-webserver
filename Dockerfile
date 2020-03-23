FROM golang

RUN go get -u github.com/tmpest/go-webserver

# Enable module management
# May or may not be necessary 
# ENV GO111MODULE=on
# ENV GOFLAGS=-mod=vendor

# COPY /go/src/github.com/tmpest/go-webserver /go/src/github.com/tmpest/go-webserver

# RUN go build go-webserver/cmd/webserver/webserver.go
ENV WEBSERVER_ROOT "/go/src/github.com/tmpest/go-webserver"
# ENTRYPOINT /go/bin/go-webserver
ENTRYPOINT ping google.com
# EXPOSE 8080gi