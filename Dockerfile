# syntax=docker/dockerfile:1

##
## Build
##
FROM golang:1.17-buster AS build

WORKDIR /app

COPY go.mod ./
COPY go.sum ./
RUN go mod download

ADD . ./

RUN go get
RUN go build -o /owl-of-athena

##
## Deploy
##
FROM gcr.io/distroless/base-debian10

WORKDIR /

COPY --from=build /owl-of-athena /owl-of-athena

EXPOSE 80

USER nonroot:nonroot

ENTRYPOINT ["/owl-of-athena"]
