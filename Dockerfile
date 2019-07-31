FROM golang:1.12 as builder

WORKDIR /

COPY go.mod .
COPY go.sum .

RUN go get

COPY . .

RUN go build -o gdns .

FROM ubuntu:latest

RUN apt-get update && apt-get install -y ca-certificates

WORKDIR /usr/bin

COPY --from=builder gdns .

CMD [ "gdns" ]