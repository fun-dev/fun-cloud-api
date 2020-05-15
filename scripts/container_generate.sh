
CONTAINER_GW_PATH=./cmd/containergw
CONTAINER_PROTO_PATH=internal/container/adapters/controller/container_user.proto
CONTAINER_PATH=internal/container/adapters/controller

protoc -I . \
    -I $GOPATH/src/github.com/grpc-ecosystem/grpc-gateway/third_party/googleapis \
    -I $GOPATH/src/github.com/grpc-ecosystem/grpc-gateway \
    --go_out=plugins=grpc,paths=source_relative:. \
    internal/container/adapters/controller/container_user.proto
# grpc gateway
protoc -I . \
    -I $GOPATH/src/github.com/grpc-ecosystem/grpc-gateway/third_party/googleapis \
    -I $GOPATH/src/github.com/grpc-ecosystem/grpc-gateway \
    --grpc-gateway_out=logtostderr=true,paths=source_relative:. $CONTAINER_PROTO_PATH
#
protoc -I ./internal/container/adapters/controller \
    -I $GOPATH/src/github.com/grpc-ecosystem/grpc-gateway/third_party/googleapis \
    --swagger_out=logtostderr=true:$GOPATH/src/github.com/fun-dev/fun-cloud-api/docs/container $CONTAINER_PROTO_PATH