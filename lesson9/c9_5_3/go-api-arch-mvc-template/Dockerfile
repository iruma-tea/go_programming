FROM golang:1.23.6-alpine3.21

RUN apk --no-cache add curl

WORKDIR /go/src/web

COPY .. .

RUN go mod download

ENV GO111MODULE=on

RUN go build main.go

ENTRYPOINT ./main