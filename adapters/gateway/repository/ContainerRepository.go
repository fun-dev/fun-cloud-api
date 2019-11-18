package repository

import (
	"context"
	"github.com/fun-dev/ccms-poc/domain/container"
)

// ContainerRepository is interface
type ContainerRepository interface {
	// --- External --- //
	GetAllByUserID(ctx *context.Context, id, namespace string) ([]container.Container, error)
	Create(ctx *context.Context, imageName, namespace string) error
	DeleteByContainerID(ctx *context.Context, id, namespace string) error
	// --- Internal --- //
	GetDeploymentManifestByContainerID(ctx *context.Context, id string) (string, error)
}