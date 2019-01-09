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

func (repo containerService) GetContainersByToken(token string) ([]models.Container, error) {
	container := models.Container{
		Id:          999,
		ImageId:     9999,
		ConnectInfo: "hogeConnection",
		Status:      "hogehoge",
	}
	return []models.Container{container}, nil
}

func (repo containerService) PostContainerByID(UserID int, ContainerID int) []models.Container {
	container := models.Container{
		Id:          UserID,
		ImageId:     ContainerID,
		ConnectInfo: "gehogehoConnection",
		Status:      "gehogeho",
	}
	return []models.Container{container}
}

func (repo containerService) PostContainerByToken(userToken, imageID string) (models.Container, error) {
	container, err := infrastructure.ContainerRepo.CreateContainer(userToken, imageID)
	if err != nil {
		return models.Container{}, err
	}
	return container, nil
}

func (repo containerService) DeleteContainerByID(userToken string, containerID int64) error {
	err := infrastructure.ContainerRepo.DeleteContainerByID(userToken, containerID)
	if err != nil {
		return err
	}
	return nil
}
