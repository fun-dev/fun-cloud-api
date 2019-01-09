package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/fun-dev/cloud-api/application"
	"github.com/fun-dev/cloud-api/config"
	_ "github.com/fun-dev/cloud-api/infrastructure"
	"github.com/fun-dev/cloud-api/infrastructure/dbmodels"
	"github.com/fun-dev/cloud-api/middleware"
	"github.com/gin-gonic/gin"
	"github.com/go-xorm/xorm"
)

func main() {
	router := setupRouter()
	//migrate()
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
	router.Use(middleware.TokenAuthMiddleware())
	router.GET("/users", usrCtrl.Get)
	router.POST("/users", usrCtrl.Create)
	router.PUT("/users", usrCtrl.Update)
	imgCtrl := application.ImageController
	router.GET("/images", imgCtrl.Get)
	return router
}

func migrate() {
	connectionString := fmt.Sprintf(
		"%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True",
		config.GetSQLUser(),
		config.GetSQLPass(),
		config.GetSQLHost(),
		config.GetSQLPort(),
		config.GetSQLDB(),
	)
	engine, err := xorm.NewEngine("mysql", connectionString)
	if err != nil {
		panic(err)
	}
	defer engine.Close()

	if err := engine.Sync2(new(dbmodels.User)); err != nil {
		panic(err)
	}
}
