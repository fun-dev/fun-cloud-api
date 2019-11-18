package controller

import (
	"github.com/fun-dev/ccms-poc/application/usecase"
	"github.com/fun-dev/ccms-poc/infrastructure/apperror/ctlerr"
	"github.com/gin-gonic/gin"
	"net/http"
)

// ContainerController is
type ContainerController struct {
	ContainerDelete usecase.ContainerDeleteUsecase
}

func NewContainerController(containerDelete usecase.ContainerDeleteUsecase) *ContainerController {
	return &ContainerController{
		ContainerDelete: containerDelete,
	}
}

// Delete is
func (cc *ContainerController) Delete(c *gin.Context) {
	containerID := c.Param("container_id")
	if containerID == "" {
		c.JSON(http.StatusBadRequest, gin.H{"message": ctlerr.ContainerIDCanNotBeFoundOnParam.Error()})
		return
	}
	if err := cc.ContainerDelete.Execute(c, containerID); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, gin.H{"message": "success on creating container"})
	return
}
