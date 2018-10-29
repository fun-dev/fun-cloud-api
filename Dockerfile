FROM golang:1.11.1-alpine3.8 AS build-env
COPY ./ /go/src/github.com/fun-dev/cloud-api/
WORKDIR /go/src/github.com/fun-dev/cloud-api
RUN go build -o server main.go

FROM alpine:latest
RUN apk add --no-cache --update ca-certificates
COPY --from=build-env /go/src/github.com/fun-dev/cloud-api/server /usr/local/bin/server
ENV SQL_USER root
ENV SQL_PASS pass
ENV SQL_HOST 127.0.0.1
ENV SQL_PORT 3306
ENV SQL_DB prac
ENTRYPOINT ["/usr/local/bin/server"]