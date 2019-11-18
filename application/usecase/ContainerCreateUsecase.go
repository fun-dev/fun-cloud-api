package usecase

import (
	"github.com/fun-dev/ccms-poc/adapters/gateway/repository"
	"github.com/gin-gonic/gin"
)

type (
	ContainerCreateUsecase interface {
		Execute(ctx *gin.Context, imageName string) error
	}
	// ContainerDeleteInteractor is Interactor
	ContainerCreateInteractor struct {
		ContainerRepo repository.ContainerRepository
		AuthRepo      repository.AuthRepository
	}
)

func NewContainerCreateInteractor() ContainerCreateUsecase {
	return &ContainerCreateInteractor{}
}

func (c ContainerCreateInteractor) Execute(ctx *gin.Context, imageName string) error {
	panic("implement me")
}