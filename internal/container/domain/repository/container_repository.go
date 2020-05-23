package repository

import (
	"github.com/fun-dev/fun-cloud-protobuf/container/rpc"
)

// ContainerRepository is interface
type ContainerRepository interface {
	// Kubernetes
	GetAllByUserID(userID, namespace string) ([]*rpc.Container, error)
	Create(userID, imageName, namespace string) error
	DeleteByContainerID(userID, containerID, namespace string) error
}
