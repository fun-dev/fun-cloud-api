package usecase

import (
	"errors"

	"github.com/fun-dev/ccms-poc/adapters/gateway/repository"
	"github.com/gin-gonic/gin"
)

type (
	// ContainerDeleteUsecase is Usecase
	ContainerDeleteUsecase interface {
		Execute(ctx *gin.Context, containerID string) error
	}
	// ContainerDeleteInteractor is Interactor
	ContainerDeleteInteractor struct {
		ContainerGateway repository.ContainerRepository
		AuthGateway      repository.AuthRepository
	}
)

// NewContainerDeleteInteractor is ...
func NewContainerDeleteInteractor(gw repository.ContainerRepository) ContainerDeleteUsecase {
	return &ContainerDeleteInteractor{ContainerGateway: gw}
}

// Execute ...
func (i *ContainerDeleteInteractor) Execute(ctx *gin.Context, containerID string) error {
	accessToken, ok := ctx.Get("Authorization")
	if !ok {
		return errors.New("token is empty")
	}
	userID, err := i.AuthGateway.GetUserIDByToken(accessToken.(string))
	if err != nil {
		return errors.New("unauthorized")
	}
	namespace := userID.Value
	err = i.ContainerGateway.DeleteByContainerID(ctx, containerID, namespace)
	if err != nil {
		return errors.New(err.Error())
	}
	return nil
}
