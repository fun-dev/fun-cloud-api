package usecase

import (
	"context"
	"github.com/fun-dev/fun-cloud-api/internal/container/domain/container"
	"github.com/fun-dev/fun-cloud-api/pkg/auth"
)

type (
	ContainerCreateUsecase interface {
		Execute(ctx context.Context, userID, imageName string) error
	}
	// ContainerDeleteInteractor is Interactor
	ContainerCreateInteractor struct {
		cRepo container.Repository
		aRepo auth.Repository
	}
)

func NewContainerCreateInteractor(cRepo container.Repository, aRepo auth.Repository) ContainerCreateUsecase {
	return &ContainerCreateInteractor{cRepo, aRepo}
}

func (c ContainerCreateInteractor) Execute(ctx context.Context, userID, imageName string) error {
	// in this application, we use userID as kubernetes namespace.yaml
	containerID, manifest, err := c.cRepo.Create(userID, imageName)
	if err != nil {
		return err
	}
	if err := c.cRepo.SaveDeploymentManifestByContainerID(containerID, manifest); err != nil {
		return err
	}
	return nil
}
