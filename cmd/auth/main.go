package main

import (
	"github.com/fun-dev/fun-cloud-api/internal/auth/service"
	"github.com/gin-gonic/gin"
	"log"
)

func main() {
	// Init Controller
	authMockCtrl := service.NewAuthMockController()
	authCtrl := service.NewAuthController()
	// setup routing
	router := gin.Default()
	mock := router.Group("/mock")
	{
		mock.GET("/token/validate", authMockCtrl.TokenValidate)
	}
	v1 := router.Group("/v1")
	{
		v1.GET("/token/validate", authCtrl.TokenValidate)
	}
	if err := router.Run(":3000"); err != nil {
		log.Fatalln(err)
	}
}
