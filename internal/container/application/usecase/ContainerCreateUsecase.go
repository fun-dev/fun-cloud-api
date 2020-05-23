package usecase

import (
	"github.com/fun-dev/fun-cloud-api/internal/container/domain/repository"
	"github.com/fun-dev/fun-cloud-api/pkg/cloudauth"
)

type (
	ContainerCreateUsecase interface {
		Execute(userID, imageName string) error
	}
	// ContainerDeleteInteractor is Interactor
	ContainerCreateInteractor struct {
		cRepo repository.Repository
		aRepo cloudauth.Repository
	}
)

func NewContainerCreateInteractor(cRepo repository.Repository, aRepo cloudauth.Repository) ContainerCreateUsecase {
	return &ContainerCreateInteractor{cRepo, aRepo}
}

func (c ContainerCreateInteractor) Execute(userID, imageName string) error {
	// in this application, we use userID as cloudk8s namespace.yaml
	containerID, manifest, err := c.cRepo.Create(userID, imageName)
	if err != nil {
		return err
	}
	if err := c.cRepo.SaveDeploymentManifestByContainerID(containerID, manifest); err != nil {
		return err
	}
	return nil
}
