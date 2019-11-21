package gateway

import (
	"context"
	"fmt"
	"github.com/fun-dev/ccms/adapters/gateway/repository"
	"github.com/fun-dev/ccms/domain"
	"github.com/fun-dev/ccms/infrastructure/apperror/repoerr"
	"github.com/fun-dev/ccms/infrastructure/driver"
	"github.com/fun-dev/ccms/infrastructure/provider"
	"log"
)

// ContainerGateway is
type ContainerGateway struct {
	provider.KubernetesProvider
	driver.RedisDriver
}

// NewContainerGateway is
func NewContainerGateway(kubernetes provider.KubernetesProvider, redis driver.RedisDriver) repository.ContainerRepository {
	return &ContainerGateway{
		kubernetes,
		redis,
	}
}

func (g ContainerGateway) GetAllByUserID(ctx context.Context, id, namespace string) ([]domain.Container, error) {
	panic("implement me")
}

func (g ContainerGateway) Create(ctx context.Context, imageName, namespace string) error {
	g.KubernetesProvider.Kubectl.DeserializeYamlToObject()
	panic("implement me")
}

// DeleteByContainerID is
func (g ContainerGateway) DeleteByContainerID(ctx context.Context, id string, namespace string) error {
	deploymentManifest, err := g.GetDeploymentManifestByContainerID(ctx, id)
	if err != nil {
		return fmt.Errorf("call ContainerGateway.GetDeploymentManifestByContainerID: %w", err)
	}
	log.Printf("[debug] deployment manifest (%s) on ContainerGateway.DeleteByContainerID()\n", deploymentManifest)
	return g.KubernetesProvider.Kubectl.Execute(driver.KubectlOptionDelete, deploymentManifest, namespace)
}

func (g ContainerGateway) GetDeploymentManifestByContainerID(ctx context.Context, id string) (string, error) {
	key := "deployment_" + id
	deploymentManifest := g.RedisDriver.Client.Get(key).String()
	if deploymentManifest == "" {
		return "", repoerr.DeploymentManifestCanNotBeFound
	}
	log.Printf("[debug] result (%s) on ContainerGateway.GetDeploymentManifestByContainerID()\n", deploymentManifest)
	return deploymentManifest, nil
}
