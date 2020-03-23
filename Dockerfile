FROM golang

RUN go get -u github.com/tmpest/go-webserver

# Enable module management
# May or may not be necessary 
# ENV GO111MODULE=on
# ENV GOFLAGS=-mod=vendor

# COPY go-webserver go-webserver

# RUN go build go-webserver/cmd/webserver/webserver.go

ENTRYPOINT /go/bin/go-webserver

EXPOSE 8080