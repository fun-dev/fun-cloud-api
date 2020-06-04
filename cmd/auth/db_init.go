package main

import (
	"github.com/fun-dev/fun-cloud-api/internal/auth/model"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"log"
	"os"
)

var (
	_dbType = "mysql"
)

func main() {
	// ex: root:root@/cloudauth?charset=utf8&parseTime=True&loc=Local
	connectStr := os.Args[1]
	db, err := gorm.Open(_dbType, connectStr)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	db.LogMode(true)

	if err := db.AutoMigrate(&model.Scope{}, &model.Role{}).Error; err != nil {
		log.Fatal(err)
	}
}