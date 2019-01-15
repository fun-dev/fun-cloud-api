FROM golang:1.11.1-alpine3.8 AS build-env
COPY ./ /go/src/github.com/fun-dev/cloud-api/
WORKDIR /go/src/github.com/fun-dev/cloud-api
RUN go build -o server main.go

FROM alpine:latest
RUN apk add --no-cache --update ca-certificates
COPY --from=build-env /go/src/github.com/fun-dev/cloud-api/server /usr/local/bin/server
COPY ./kubeconfig /opt/kubeconfig
ENV SQL_USER root
ENV SQL_PASS pass
ENV SQL_HOST 127.0.0.1
ENV SQL_PORT 3306
ENV SQL_DB prac
ENV K8S_CONFIG_PATH /opt/kubeconfig
ENV K8S_IP 35.232.60.106
ENV PROXY_ADDR 35.211.84.216:8081
ENV GOOGLE_CLIENT_ID=43129496828-cubp85rjtc5su7mlslpcg6enadreb7gk.apps.googleusercontent.com
ENV GOOGLE_CLIENT_SECRET=j7l-0co59KBVdDlnjhFtA-oa
ENV GOOGLE_TOKEN_VALIDATE=https://www.googleapis.com/oauth2/v3/tokeninfo?id_token=

EXPOSE 8080
ENTRYPOINT ["/usr/local/bin/server"]