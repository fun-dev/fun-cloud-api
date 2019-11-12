package gateway

import (
	"errors"
	"github.com/fun-dev/ccms-poc/adapters/gateway/repository"
	"github.com/fun-dev/ccms-poc/infrastructure/config"
	"github.com/fun-dev/ccms-poc/infrastructure/driver"
	"github.com/fun-dev/ccms-poc/infrastructure/provider"
	"github.com/gin-gonic/gin"
	"log"
)

// ContainerGateway is
type ContainerGateway struct {
	KubernetesProvider *provider.KubernetesProvider
	RedisDriver        *driver.RedisDriver
	Config             *config.AppVariableOnKubectl
}

// NewContainerGateway is
func NewContainerGateway(
	kubernetes *provider.KubernetesProvider,
	redis *driver.RedisDriver,
	config *config.AppVariableOnKubectl,
) repository.ContainerRepository {
	return &ContainerGateway{
		KubernetesProvider: kubernetes,
		RedisDriver:        redis,
		Config:             config,
	}
}

// DeleteByContainerID is
func (g *ContainerGateway) DeleteByContainerID(ctx *gin.Context, id string, namespace string) error {
	deploymentManifest, err := g.GetDeploymentManifestByContainerID(ctx, id)
	if err != nil {
		return err
	}
	log.Printf("[debug] deployment manifest (%s) on ContainerGateway.DeleteByContainerID()\n", deploymentManifest)
	return g.KubernetesProvider.Kubectl.Execute(driver.KubectlOptionDelete, deploymentManifest, namespace)
}

func (g *ContainerGateway) GetDeploymentManifestByContainerID(ctx *gin.Context, id string) (string, error) {
	key := "deployment_" + id
	result := g.RedisDriver.Client.Get(key)
	if result.String() == "" {
		return "", errors.New("not found")
	}
	log.Printf("[debug] result (%s) on ContainerGateway.GetDeploymentManifestByContainerID()\n", result.String())
	return result.String(), nil
}
