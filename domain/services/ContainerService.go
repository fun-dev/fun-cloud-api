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

func (repo containerService) GetContainersByUniqueUserID(uniqueUserID string) ([]models.Container, error) {
	container := models.Container{
		Id:          999,
		ImageId:     9999,
		ConnectInfo: "hogeConnection",
		Status:      "hogehoge",
	}
	return []models.Container{container}, nil
}

func (repo containerService) CreateContainer(uniqueUserID, imageName string) (models.Container, error) {
	container, err := infrastructure.ContainerRepo.CreateContainer(uniqueUserID, imageName)
	if err != nil {
		return models.Container{}, err
	}
	return container, nil
}

func (repo containerService) DeleteContainer(uniqueUserID string, containerID int64) error {
	err := infrastructure.ContainerRepo.DeleteContainer(uniqueUserID, containerID)
	if err != nil {
		return err
	}
	return nil
}
