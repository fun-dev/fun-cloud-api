package usecase

import (
	"context"
	"github.com/fun-dev/fun-cloud-api/pkg/auth"
	"github.com/fun-dev/fun-cloud-api/internal/container/domain/container"
)

type (
	ContainerReadUsecaseEntry struct {
		Containers []*container.Container
	}

	ContainerReadUsecaseResponse struct {
		Entry ContainerReadUsecaseEntry
	}

	ContainerReadUsecase interface {
		Execute(ctx context.Context, userID, imageName string) (resp ContainerReadUsecaseResponse, err error)
	}

	ContainerReadInteractor struct {
		cRepo container.Repository
		aRepo auth.AuthRepository
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
func (c ContainerReadInteractor) Execute(ctx context.Context, userID, imageName string) (resp ContainerReadUsecaseResponse, err error) {
	// in this application, we use userID as kubernetes namespace.yaml
	namespace := userID
	resp.Entry.Containers, err = c.cRepo.GetAllByUserID(ctx, userID, namespace)
	if err != nil {
		return
	}
	return
}
