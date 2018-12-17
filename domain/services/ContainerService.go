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

func (repo containerService) GetContainersByUserID(UserID int) ([]models.Container, error) {
	container := models.Container{
		ID:          999,
		ImageID:     9999,
		ConnectInfo: "hogeConnection",
		Status:      "hogehoge",
	}
	return []models.Container{container}, nil
}
