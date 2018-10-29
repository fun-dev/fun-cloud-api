package main

import (
	"github.com/fun-dev/cloud-api/infrastructure/dbmodels"
	_ "github.com/go-sql-driver/mysql"
	"github.com/go-xorm/xorm"
)

const (
	connectionString = "root:199507620@tcp(127.0.0.1:3306)/prac?charset=utf8&parseTime=True"
)

func main() {
	engine, err := xorm.NewEngine("mysql", connectionString)
	if err != nil {
		panic(err)
	}
	defer engine.Close()

	if err := engine.Sync2(new(dbmodels.User)); err != nil {
		panic(err)
	}
}
