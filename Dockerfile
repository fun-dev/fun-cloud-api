# first stage - builder
FROM golang:1.12.7-stretch as builder
COPY . /ccms
WORKDIR /ccms
ENV GO111MODULE=on
RUN CGO_ENABLED=0 GOOS=linux go build -o ccms
RUN curl -LO https://storage.googleapis.com/kubernetes-release/release/v1.15.0/bin/linux/amd64/kubectl && \
    chmod +x ./kubectl

# second stage
FROM alpine:latest
WORKDIR /root/
# Copy Program Binary
COPY --from=builder /ccms/ccms .
# Copy Kubectl
COPY --from=builder /ccms/kubectl /usr/local/bin/kubectl
# Set Env
ENV KUBECTL_PATH="/usr/local/bin/kubectl"
CMD ["./ccms"]