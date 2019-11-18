package gateway

import (
	"context"
	"fmt"
	"github.com/fun-dev/ccms-poc/adapters/gateway/repository"
	"github.com/fun-dev/ccms-poc/domain/container"
	"github.com/fun-dev/ccms-poc/infrastructure/apperror/repoerr"
	"github.com/fun-dev/ccms-poc/infrastructure/driver"
	"github.com/fun-dev/ccms-poc/infrastructure/provider"
	"log"
)

// ContainerGateway is
type ContainerGateway struct {
	provider.KubernetesProvider
	driver.RedisDriver
}

// NewContainerGateway is
func NewContainerGateway(
	kubernetes provider.KubernetesProvider,
	redis driver.RedisDriver,
) repository.ContainerRepository {
	return &ContainerGateway{
		KubernetesProvider: kubernetes,
		RedisDriver:        redis,
	}
}

func (g ContainerGateway) GetAllByUserID(ctx *context.Context, id, namespace string) ([]container.Container, error) {
	panic("implement me")
}

func (g ContainerGateway) Create(ctx *context.Context, imageName, namespace string) error {
	panic("implement me")
}

// DeleteByContainerID is
func (g ContainerGateway) DeleteByContainerID(ctx *context.Context, id string, namespace string) error {
	deploymentManifest, err := g.GetDeploymentManifestByContainerID(ctx, id)
	if err != nil {
		return fmt.Errorf("call ContainerGateway.GetDeploymentManifestByContainerID: %w", err)
	}
	log.Printf("[debug] deployment manifest (%s) on ContainerGateway.DeleteByContainerID()\n", deploymentManifest)
	return g.KubernetesProvider.Kubectl.Execute(driver.KubectlOptionDelete, deploymentManifest, namespace)
}

func (g ContainerGateway) GetDeploymentManifestByContainerID(ctx *context.Context, id string) (string, error) {
	key := "deployment_" + id
	result := g.RedisDriver.Client.Get(key)
	if result.String() == "" {
		return "", repoerr.DeploymentManifestCanNotBeFound
	}
	log.Printf("[debug] result (%s) on ContainerGateway.GetDeploymentManifestByContainerID()\n", result.String())
	return result.String(), nil
}
