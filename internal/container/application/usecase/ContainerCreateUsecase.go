package usecase

import (
	"context"
	"github.com/fun-dev/fun-cloud-api/pkg/auth"
	"github.com/fun-dev/fun-cloud-api/internal/container/domain/container"
	"github.com/fun-dev/fun-cloud-api/pkg/uuid"
)

type (
	ContainerCreateUsecase interface {
		Execute(ctx context.Context, userID, imageName string) error
	}
	// ContainerDeleteInteractor is Interactor
	ContainerCreateInteractor struct {
		cRepo container.Repository
		aRepo auth.AuthRepository
	}
)

func NewContainerCreateInteractor(cRepo container.Repository, aRepo auth.AuthRepository) ContainerCreateUsecase {
	return &ContainerCreateInteractor{cRepo, aRepo}
}

func (c ContainerCreateInteractor) Execute(ctx context.Context, userID, imageName string) error {
	// in this application, we use userID as kubernetes namespace.yaml
	containerID := uuid.NewUUID()
	namespace := userID
	manifest, err := c.cRepo.Create(ctx, containerID, imageName, namespace)
	if err != nil {
		return err
	}
	if err := c.cRepo.SaveDeploymentManifestByContainerID(ctx, containerID, manifest); err != nil {
		return err
	}
	return nil
}
