package controller

import (
	"github.com/fun-dev/ccms-poc/application/usecase"
	"github.com/fun-dev/ccms-poc/infrastructure/apperror/ctlerr"
	"github.com/gin-gonic/gin"
	"net/http"
)

// ContainerController is
type (
	ContainerController struct {
		usecase.ContainerCreateUsecase
		usecase.ContainerDeleteUsecase
	}
	PostRequest struct {
		ImageName string `json:"image_name"`
	}
)

func NewContainerController(cCre usecase.ContainerCreateUsecase, cDel usecase.ContainerDeleteUsecase) *ContainerController {
	return &ContainerController{
		cCre,
		cDel,
	}
}

// Post is
// Header: key is Authorization
// BODY: {"image_name": "nginx:latest"}
func (cc ContainerController) Post(c *gin.Context) {
	var json PostRequest
	if err := c.ShouldBindJSON(&json); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if err := cc.ContainerCreateUsecase.Execute(c, json.ImageName); err != nil{
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, gin.H{"message": "success on creating container"})
	return
}

// Delete is
func (cc ContainerController) Delete(c *gin.Context) {
	containerID := c.Param("container_id")
	if containerID == "" {
		c.JSON(http.StatusBadRequest, gin.H{"message": ctlerr.ContainerIDCanNotBeFoundOnParam.Error()})
		return
	}
	if err := cc.ContainerDeleteUsecase.Execute(c, containerID); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, gin.H{"message": "success on deleting container"})
	return
}
