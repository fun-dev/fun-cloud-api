package gateway

import (
	"github.com/fun-dev/fun-cloud-api/internal/directory/models"
	"github.com/fun-dev/fun-cloud-api/internal/directory/repository"
	"github.com/fun-dev/fun-cloud-api/pkg/cloudk8s"
)

type ContainerDirectoryGateway struct {}

func (c ContainerDirectoryGateway) List(dirPath, containerName, namespace string) (*models.Data, error) {
	listCmdResult, err := cloudk8s.ExecuteListCmdOnKubectl(dirPath, containerName, namespace)
	if err != nil {
		return nil, err
	}
	return models.ParseLsCmdData(listCmdResult, dirPath)
}

func NewContainerDirectoryGateway() repository.ContainerDirectoryRepository {
	return &ContainerDirectoryGateway{}
}