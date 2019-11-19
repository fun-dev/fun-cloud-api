package usecase

import (
	"context"
	"github.com/fun-dev/ccms-poc/adapters/gateway/repository"
)

type (
	ContainerCreateUsecase interface {
		Execute(ctx *context.Context, imageName string) error
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

func (c ContainerCreateInteractor) Execute(ctx *context.Context, imageName string) error {
	panic("implement me")
}