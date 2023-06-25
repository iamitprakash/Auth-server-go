# syntax=docker/dockerfile:1
FROM golang:1.16-alpine

WORKDIR /server

COPY go.mod ./
COPY go.sum ./

RUN go mod download
COPY ./ ./

RUN go build -o /Auth-server-go

EXPOSE 9090

CMD [ "/Auth-server-go" ]