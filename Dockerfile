FROM golang:1.10.2-alpine3.7 AS build
RUN apk --no-cache add gcc g++ make ca-certificates
WORKDIR /go/src/github.com/griffinup/yachtsearch

COPY vendor vendor
COPY Gopkg.lock Gopkg.toml ./
COPY util util
COPY db db
COPY search search
COPY schema schema
COPY update-service update-service
COPY search-service search-service

RUN go install ./...

FROM alpine:3.7
WORKDIR /usr/bin
COPY --from=build /go/bin .
