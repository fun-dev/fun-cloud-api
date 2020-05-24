package data

import (
	"context"
	"github.com/fun-dev/fun-cloud-protobuf/data/rpc"
	"github.com/golang/protobuf/ptypes/empty"
)

type TransferService struct{}

func (t TransferService) UploadToContainer(ctx context.Context, request *rpc.UploadToContainerRequest) (*empty.Empty, error) {
	panic("implement me")
}

func (t TransferService) DownloadFromContainer(ctx context.Context, request *rpc.DownloadFromContainerRequest) (*rpc.DownloadFromContainerResponse, error) {
	panic("implement me")
}

func NewTransferService() rpc.DataTransferServiceServer {
	return &TransferService{}
}