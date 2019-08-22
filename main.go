package main

import (
	"log"
	"net/http"

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

	// Health Check Endpoint
	router.GET("/health", func(c *gin.Context) {
		c.Status(http.StatusOK)
	})
	// Client should add /api/v1 at first on url
	v1 := router.Group("/api/v1")

	// REST Container
	v1.GET("/containers", )
	v1.POST("/containers", )
	v1.DELETE("/containers/:id",)

	return router
}
