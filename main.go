package main

import (
	"log"
	"net/http"

	"github.com/fun-dev/cloud-api/application"
	"github.com/gin-gonic/gin"
)

func main() {
	router := setupRouter()
	if err := router.Run(":8080"); err != nil {
		log.Fatal(err)
	}
}

func setupRouter() *gin.Engine {
	router := gin.Default()
	router.GET("/health", func(c *gin.Context) {
		c.Status(http.StatusOK)
	})

	usrCtrl := application.UserController

	router.GET("/user/:id", usrCtrl.Get)
	router.POST("/user", usrCtrl.Create)
	router.PUT("/user/:id", usrCtrl.Update)
	router.DELETE("/user/:id", usrCtrl.Delete)

	router.GET("/containers", application.ContainerController.Get)
	router.POST("/containers/user/:id", application.ContainerController.Post)
	return router
}
