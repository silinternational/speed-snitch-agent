FROM golang:1.15
# The previous deployment (with ":latest") appears to have matched the following image version
# FROM golang:1.12.6-stretch

RUN apt-get update -y && apt-get install -y awscli git
RUN go get -u github.com/golang/dep/cmd/dep && go get github.com/mitchellh/gox
RUN wget -P /usr/local/bin/ https://raw.githubusercontent.com/silinternational/runny/0.2/runny && chmod a+x /usr/local/bin/runny

COPY ./ /go/src/github.com/silinternational/speed-snitch-agent/
WORKDIR /go/src/github.com/silinternational/speed-snitch-agent/
RUN dep ensure