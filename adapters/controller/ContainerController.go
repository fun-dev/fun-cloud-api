package controller

import (
	"github.com/fun-dev/ccms-poc/adapters/gateway"
	"github.com/fun-dev/ccms-poc/infrastructure/config"
	"github.com/fun-dev/ccms-poc/infrastructure/driver"
	"github.com/fun-dev/ccms-poc/infrastructure/provider"
	"net/http"
	"strconv"

	usecase "github.com/fun-dev/ccms-poc/application/usecase"
	"github.com/gin-gonic/gin"
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
	castContainerID, _ := strconv.Atoi(containerID)
	err := cc.ContainerRead.Execute(c, castContainerID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, gin.H{})
	return
}
