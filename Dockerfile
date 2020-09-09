FROM alpine
#FROM golang:alpine

RUN apk add gcc

COPY ./main /app

WORKDIR /app

ENTRYPOINT ["main"]