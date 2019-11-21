package usecase

import (
	"context"
	"github.com/fun-dev/ccms/internal/container/adapters/gateway/repository"
	"github.com/fun-dev/ccms/internal/container/domain"
)

// input

// output

type (
	ContainerReadUsecaseEntry struct {
		Containers []domain.Container
	}

	ContainerReadUsecaseResponse struct {
		Entry ContainerReadUsecaseEntry
	}

	ContainerReadUsecase interface {
		Execute(ctx context.Context, userID, imageName string) (resp *ContainerReadUsecaseResponse, err error)
	}

	ContainerReadInteractor struct {
		repository.ContainerRepository
		repository.AuthRepository
	}
)

func NewContainerReadInteractor() ContainerReadUsecase {
	return &ContainerReadInteractor{}
}

/*
Execute
@Option: named function
@param: userID
@param imageName
*/
func (c ContainerReadInteractor) Execute(ctx context.Context, userID, imageName string) (resp *ContainerReadUsecaseResponse, err error) {
	// in this application, we use userID as kubernetes namespace.yaml
	namespace := userID
	resp.Entry.Containers, err = c.ContainerRepository.GetAllByUserID(ctx, userID, namespace)
	if err != nil {
		return
	}
	return
}
