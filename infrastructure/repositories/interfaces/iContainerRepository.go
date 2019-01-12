package interfaces

import "github.com/fun-dev/cloud-api/domain/models"

type IContainerRepository interface {
	GetContainersByNamespace(namespace string) ([]models.Container, error)
	CreateContainer(uniqueUserID, imageName string) error
	DeleteContainer(uniqueUserID string, containerID int64) error
}
