package container

import (
	"context"
)

// ContainerRepository is interface
type Repository interface {
	// --- External --- //
	GetAllByUserID(ctx context.Context, id, namespace string) ([]*Container, error)
	Create(ctx context.Context, id, imageName, namespace string) (manifest string, err error)
	DeleteByContainerID(ctx context.Context, id, namespace string) error
	GetDeploymentManifestByContainerID(ctx context.Context, id string) (manifest string, err error)
	SaveDeploymentManifestByContainerID(ctx context.Context, userID, id, yaml string) error
}
