package controller

import (
	"github.com/fun-dev/ccms-poc/adapters/gateway"
	usecase "github.com/fun-dev/ccms-poc/application/usecase"
	"github.com/fun-dev/ccms-poc/infrastructure/config"
	"github.com/fun-dev/ccms-poc/infrastructure/driver"
	"github.com/fun-dev/ccms-poc/infrastructure/provider"
	"github.com/gin-gonic/gin"
	"net/http"
)

// ContainerController is
type ContainerController struct {
	ContainerDelete usecase.ContainerDeleteUsecase
}

func NewContainerController() *ContainerController {
	return &ContainerController{
		ContainerDelete: usecase.NewContainerDeleteInteractor(
			gateway.NewContainerGateway(
				provider.KubernetesProviderImpl,
				driver.RedisDriverImpl,
				config.AppVariableOnKubectlImpl,
			),
		),
	}
}

// Delete is
func (cc *ContainerController) Delete(c *gin.Context) {
	containerID := c.Param("container_id")
	err := cc.ContainerDelete.Execute(c, containerID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, gin.H{})
	return
}
