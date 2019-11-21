package usecase

import (
	"context"
	"github.com/fun-dev/ccms/adapters/gateway/repository"
)

type (
	ContainerCreateUsecase interface {
		Execute(ctx context.Context, userID, imageName string) error
	}
	// ContainerDeleteInteractor is Interactor
	ContainerCreateInteractor struct {
		repository.ContainerRepository
		repository.AuthRepository
	}
)

func NewContainerCreateInteractor(cRepo repository.ContainerRepository, aRepo repository.AuthRepository) ContainerCreateUsecase {
	return &ContainerCreateInteractor{cRepo, aRepo}
}

func (c ContainerCreateInteractor) Execute(ctx context.Context, userID, imageName string) error {
	// in this application, we use userID as kubernetes namespace
	namespace := userID
	if err := c.ContainerRepository.Create(ctx, imageName, namespace); err != nil {
		return err
	}
	return nil
}
