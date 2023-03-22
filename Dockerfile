FROM golang:1.20

RUN apt-get update -y && apt-get install -y awscli git

WORKDIR /src
COPY . .

RUN go get github.com/mitchellh/gox
RUN wget -P /usr/local/bin/ https://raw.githubusercontent.com/silinternational/runny/0.2/runny && chmod a+x /usr/local/bin/runny
