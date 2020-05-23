package service

import (
	"context"
	"github.com/fun-dev/fun-cloud-api/internal/container/application/usecase"
	pb "github.com/fun-dev/fun-cloud-protobuf/container/rpc"
	"github.com/golang/protobuf/ptypes/empty"
)

type ContainerService struct {
	usecase.ContainerReadUsecase
	usecase.ContainerCreateUsecase
	usecase.ContainerDeleteUsecase
}

func (c ContainerService) Create(ctx context.Context, request *pb.CreateContainerRequest) (*empty.Empty, error) {
	panic("implement me")
}

func (c ContainerService) List(ctx context.Context, request *pb.ListContainerRequest) (*pb.ListContainerResponse, error) {
	panic("implement me")
}

func (c ContainerService) Delete(ctx context.Context, request *pb.DeleteContainerRequest) (*empty.Empty, error) {
	panic("implement me")
}

func NewContainerService(
	cRed usecase.ContainerReadUsecase,
	cCre usecase.ContainerCreateUsecase,
	cDel usecase.ContainerDeleteUsecase,
) pb.ContainerServiceServer {
	return &ContainerService{
		cRed,
		cCre,
		cDel,
	}
}
