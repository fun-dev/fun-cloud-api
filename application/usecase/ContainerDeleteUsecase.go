package usecase

import (
	"context"
	"fmt"
	"github.com/fun-dev/ccms-poc/infrastructure/apperror/usecaseerr"

	"github.com/fun-dev/ccms-poc/adapters/gateway/repository"
)

type (
	// ContainerDeleteUsecase is Usecase
	ContainerDeleteUsecase interface {
		Execute(ctx *context.Context, containerID string) error
	}
	// ContainerDeleteInteractor is Interactor
	ContainerDeleteInteractor struct {
		ContainerRepo repository.ContainerRepository
		AuthRepo      repository.AuthRepository
	}
)

// NewContainerDeleteInteractor is ...
func NewContainerDeleteInteractor(containerRepo repository.ContainerRepository, authRepo repository.AuthRepository) ContainerDeleteUsecase {
	return &ContainerDeleteInteractor{ContainerRepo: containerRepo, AuthRepo:authRepo}
}

// Execute ...
func (i ContainerDeleteInteractor) Execute(ctx *context.Context, containerID string) error {
	accessToken, ok := ctx.Get("Authorization")
	if !ok {
		return usecaseerr.AuthorizationIsNotFoundOnParam
	}
	userID, err := i.AuthRepo.GetUserIDByToken(accessToken.(string))
	if err != nil {
		return usecaseerr.UserIDCanNotBeFoundOnStore
	}
	namespace := userID.Value
	if err = i.ContainerRepo.DeleteByContainerID(ctx., containerID, namespace); err != nil {
		return fmt.Errorf("call ContainerRepo.DeleteByContainerID: %w", err)
	}
	return nil
}
