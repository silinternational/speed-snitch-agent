FROM golang:latest

RUN apt-get update -y && apt-get install -y awscli git
RUN go get -u github.com/golang/dep/cmd/dep
RUN go get github.com/mitchellh/gox
RUN wget -P /usr/local/bin/ https://raw.githubusercontent.com/silinternational/runny/0.2/runny && chmod a+x /usr/local/bin/runny