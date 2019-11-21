package driver

import (
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

var (
	_defaultPort = ":3000"
	_router      = gin.Default()
)

type IContainerController interface {
	Delete(c *gin.Context)
}

type GinDriver struct {
	ContainerCtrl IContainerController
	Router        *gin.Engine
}

func NewGinDriver() *GinDriver {
	result := &GinDriver{}
	result.Router = _router
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
	// --- Client should add /api/v1 at first on url
	v1 := d.Router.Group("/api/v1")
	// --- REST ENDPOINT --- //
	v1.GET("/containers")
	v1.POST("/containers")
	v1.DELETE("/containers/:id", d.ContainerCtrl.Delete)
	return nil
}

func (d *GinDriver) Run() error {
	if err := d.Router.Run(_defaultPort); err != nil {
		return err
	}
	return nil
}
