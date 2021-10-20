# syntax=docker/dockerfile:1

FROM golang:1.14-alpine

WORKDIR /app

COPY go.mod ./
COPY go.sum ./
RUN go mod download

COPY *.go ./

RUN go build -o /go-blog-api

EXPOSE 8080

CMD [ "/go-blog-api" ]
