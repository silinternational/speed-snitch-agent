FROM golang:latest

RUN apt-get update -y && apt-get install -y git
RUN go get -u github.com/golang/dep/cmd/dep
RUN go get github.com/mitchellh/gox