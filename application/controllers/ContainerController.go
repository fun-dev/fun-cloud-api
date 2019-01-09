package controllers

import (
	"net/http"

	"github.com/fun-dev/cloud-api/application/controllers/interfaces"
	"github.com/fun-dev/cloud-api/application/viewmodels"
	"github.com/fun-dev/cloud-api/domain"
	isrv "github.com/fun-dev/cloud-api/domain/services/interfaces"
	"github.com/gin-gonic/gin"
)

type ContainerController struct {
	Srv isrv.IContainerService
}

func NewContainerController() interfaces.IContainerController {
	return ContainerController{
		Srv: domain.ContainerSrv,
	}
}

func (ctrl ContainerController) Get(c *gin.Context) {
	userToken := c.GetHeader("Authorization")
	containers, err := ctrl.Srv.GetContainersByToken(userToken)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"err": err.Error()})
		return
	}
	c.JSON(http.StatusOK, containers)
}

func (ctrl ContainerController) Post(c *gin.Context) {
	containerImage := viewmodels.ContainerImage{}
	userToken := c.GetHeader("Authorization")
	c.BindJSON(&containerImage)
	containers, err := ctrl.Srv.PostContainerByToken(userToken, containerImage.ImageId)
	if err != nil {
		c.JSON(500, gin.H{
			"err": err.Error(),
		})
	}
	c.JSON(http.StatusCreated, containers)
}

func (ctrl ContainerController) Delete(c *gin.Context) {
	userToken := c.GetHeader("Authorization")
	err := ctrl.Srv.DeleteContainerByID(userToken, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"err": err.Error()})
		return
	}
	c.Status(http.StatusNoContent)
}
