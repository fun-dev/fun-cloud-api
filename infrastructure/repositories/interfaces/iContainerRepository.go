package interfaces

import "github.com/fun-dev/cloud-api/domain/models"

type IContainerRepository interface {
	GetContainersByNamespace(namespace string) ([]models.Container, error)
	CreateContainer(userToken, imageID string) (models.Container, error)
	DeleteContainerByID(userToken string, containerID int64) error
}
