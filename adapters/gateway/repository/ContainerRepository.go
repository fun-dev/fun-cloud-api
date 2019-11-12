package repository

import "github.com/gin-gonic/gin"

// ContainerRepository is interface
type ContainerRepository interface {
	GetDeploymentManifestByContainerID(ctx *gin.Context, id string) (string, error)
	DeleteByContainerID(ctx *gin.Context, id string, namespace string) error
}
