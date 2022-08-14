# syntax=docker/dockerfile:1

##
## Build
##
FROM golang:1.18-buster AS build

WORKDIR /app

COPY src/go.mod ./
COPY src/go.sum ./
RUN go mod download

COPY src/*.go ./

RUN go build -o /go-chatbot

##
## Deploy
##
FROM gcr.io/distroless/base-debian10
#FROM alpine:3.16

#RUN apk update && apk upgrade && apk add bash

WORKDIR /

COPY --from=build /go-chatbot usr/local/bin/go-chatbot

EXPOSE 8080

#USER nonroot:nonroot

WORKDIR /usr/local/bin

ENTRYPOINT ["/usr/local/bin/go-chatbot"]
