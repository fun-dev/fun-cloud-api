package usecase

import (
	"context"
	"github.com/fun-dev/fun-cloud-api/internal/container/domain/auth"
	"github.com/fun-dev/fun-cloud-api/internal/container/domain/container"
)

type (
	ContainerCreateUsecase interface {
		Execute(ctx context.Context, userID, imageName string) error
	}
	// ContainerDeleteInteractor is Interactor
	ContainerCreateInteractor struct {
		container.ContainerRepository
		auth.AuthRepository
	}
)

func NewContainerCreateInteractor(cRepo container.ContainerRepository, aRepo auth.AuthRepository) ContainerCreateUsecase {
	return &ContainerCreateInteractor{cRepo, aRepo}
}

func (c ContainerCreateInteractor) Execute(ctx context.Context, userID, imageName string) error {
	// in this application, we use userID as kubernetes namespace.yaml
	namespace := userID
	if err := c.ContainerRepository.Create(ctx, imageName, namespace); err != nil {
		return err
	}
	return nil
}
