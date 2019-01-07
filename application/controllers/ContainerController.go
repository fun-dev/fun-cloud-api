package controllers

import (
	"net/http"

	"github.com/fun-dev/cloud-api/application/controllers/interfaces"
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
	containers, err := ctrl.Srv.GetContainersByUserID(999)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"err": err.Error()})
		return
	}
	c.JSON(http.StatusOK, containers)
}

func (ctrl ContainerController) Post(c *gin.Context) {
	containers := ctrl.Srv.PostContainerByID(999, 9998)
	c.JSON(http.StatusCreated, containers)
}
