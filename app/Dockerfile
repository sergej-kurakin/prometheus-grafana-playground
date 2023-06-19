# syntax=docker/dockerfile:1


##
## Build
##

FROM golang:1.20-alpine AS build

WORKDIR /app

COPY go.mod ./
COPY go.sum ./

RUN go mod download

COPY *.go ./

RUN go build -o /app/http-server

##
## Deploy
##

FROM alpine

WORKDIR /

COPY --from=build /app/http-server /http-server

EXPOSE 8090/tcp

ENTRYPOINT ["/http-server"]