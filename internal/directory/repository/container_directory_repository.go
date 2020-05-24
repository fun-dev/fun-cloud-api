package repository

import "github.com/fun-dev/fun-cloud-api/internal/directory/models"

type ContainerDirectoryRepository interface {
	List(dirPath, containerName, namespace string) (*models.Data, error)
}