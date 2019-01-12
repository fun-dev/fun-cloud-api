package services

import (
	"github.com/fun-dev/cloud-api/domain/models"
	"github.com/fun-dev/cloud-api/domain/services/interfaces"
	"github.com/fun-dev/cloud-api/infrastructure"
	irepo "github.com/fun-dev/cloud-api/infrastructure/repositories/interfaces"
)

type containerService struct {
	Repo irepo.IContainerRepository
}

func NewContainerService() interfaces.IContainerService {
	return containerService{Repo: infrastructure.ContainerRepo}
}

func (srv containerService) GetContainersByUniqueUserID(uniqueUserID string) ([]models.Container, error) {
	containers, err := srv.Repo.GetContainersByNamespace(uniqueUserID)
	if err != nil {
		return nil, err
	}
	return containers, nil
}

func (srv containerService) CreateContainer(uniqueUserID, imageName string) error {
	err := infrastructure.ContainerRepo.CreateContainer(uniqueUserID, imageName)
	if err != nil {
		return err
	}
	return nil
}

func (srv containerService) DeleteContainer(uniqueUserID string, containerID int64) error {
	err := infrastructure.ContainerRepo.DeleteContainer(uniqueUserID, containerID)
	if err != nil {
		return err
	}
	return nil
}
