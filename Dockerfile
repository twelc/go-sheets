FROM golang:1.21.6

WORKDIR /usr/src/app

COPY . .

RUN go mod tidy

RUN go build github.com/twelc/go-sheets