package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/fun-dev/cloud-api/application"
	"github.com/fun-dev/cloud-api/config"
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
	// routerの生成
	router := gin.Default()

	// healthチェック用のエンドポイント作成
	router.GET("/health", func(c *gin.Context) {
		c.Status(http.StatusOK)
	})

	// 各種コントローラのインスタンス取得
	usrCtrl := application.UserController
	imgCtrl := application.ImageController
	containerCtrl := application.ContainerController

	// JWTTokenを確認するmiddlewareの登録
	router.Use(middleware.TokenAuthMiddleware())

	// routerグループの作成
	v1 := router.Group("/api/v1")

	// users api
	v1.GET("/users", usrCtrl.Get)
	v1.POST("/users", usrCtrl.Create)
	v1.PUT("/users", usrCtrl.Update)

	// images api
	v1.GET("/images", imgCtrl.Get)

	// containers api
	v1.GET("/containers", containerCtrl.Get)
	v1.POST("/containers", containerCtrl.Post)
	v1.DELETE("/containers/:id", containerCtrl.Delete)

	return router
}

// 初回起動時に呼ぶ関数
// 構造体からテーブル定義を自動生成
func migrate() {
	connectionString := getConnectionString()

	// DB接続
	engine, err := xorm.NewEngine("mysql", connectionString)
	if err != nil {
		panic(err)
	}
	// 接続を閉じるようにしておく
	defer engine.Close()

	// 構造体からテーブル生成
	if err := engine.Sync2(new(dbmodels.User)); err != nil {
		panic(err)
	}
}

// 接続文字列を取得する関数
func getConnectionString() string {
	return fmt.Sprintf(
		"%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True",
		config.GetSQLUser(),
		config.GetSQLPass(),
		config.GetSQLHost(),
		config.GetSQLPort(),
		config.GetSQLDB(),
	)
}
