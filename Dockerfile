FROM golang:1.15 AS builder

ENV GO111MODULE=on

RUN useradd -ms /bin/bash movieposter

WORKDIR /go/src/app

COPY go.mod .
COPY go.sum .

RUN go mod download

COPY . .

RUN go get github.com/pilu/fresh

RUN chown -R movieposter:movieposter /go/src/app

WORKDIR /go/src/app/api
USER movieposter

EXPOSE 8080
CMD ["fresh", "/go/src/app"]
