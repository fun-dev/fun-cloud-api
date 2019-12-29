package server

import (
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"os"
)

var (
	_port   = os.Getenv("SERVER_PORT")
	_router = gin.Default()
)

type IContainerController interface {
	Get(c *gin.Context)
	Post(c *gin.Context)
	Delete(c *gin.Context)
}

type GinDriver struct {
	ContainerCtrl IContainerController
	Router        *gin.Engine
}

func NewGinDriver(ctrl IContainerController) *GinDriver {
	result := &GinDriver{}
	result.Router = _router
	result.ContainerCtrl = ctrl
	if err := result.setupRouting(); err != nil {
		log.Fatal(err)
	}
	return result
}

func (d *GinDriver) setupRouting() error {
	// --- Health Check Endpoint --- //
	d.Router.GET("/health", func(c *gin.Context) {
		c.Status(http.StatusOK)
	})
	// --- Client should add /v1 at first on url
	v1 := d.Router.Group("/v1")
	// --- REST ENDPOINT --- //
	v1.GET("/containers")
	v1.POST("/containers")
	v1.DELETE("/containers/:id", d.ContainerCtrl.Delete)
	return nil
}

func (d *GinDriver) Run() error {
	if err := d.Router.Run(_port); err != nil {
		return err
	}
	return nil
}
