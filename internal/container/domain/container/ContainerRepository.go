package repository

import (
	"context"
	"github.com/fun-dev/ccms/internal/container/domain"
)

// ContainerRepository is interface
type ContainerRepository interface {
	// --- External --- //
	GetAllByUserID(ctx context.Context, id, namespace string) ([]domain.Container, error)
	Create(ctx context.Context, imageName, namespace string) error
	DeleteByContainerID(ctx context.Context, id, namespace string) error
	// --- Internal --- //
	GetDeploymentManifestByContainerID(ctx context.Context, id string) (manifest string, err error)
	SaveDeploymentManifestByContainerID(ctx context.Context, id string) error
}
