# build stage
FROM golang:1.13-stretch as builder
WORKDIR /fun-cloud-api
COPY . .
ENV GO111MODULE=on
RUN CGO_ENABLED=0 GOOS=linux go build -o container
RUN curl -LO https://storage.googleapis.com/kubernetes-release/release/v1.15.0/bin/linux/amd64/kubectl && \
    chmod +x ./kubectl

# runtime stage
FROM alpine:latest
WORKDIR /fun-cloud-api
COPY --from=builder /fun-cloud-api/container .
COPY --from=builder /fun-cloud-api/kubectl /usr/local/bin/kubectl
COPY --from=builder /fun-cloud-api/configs .
RUN set -a && source ./configs/container/dev.env
CMD ["/container"]