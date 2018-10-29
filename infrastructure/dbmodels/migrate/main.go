package main

import (
	"fmt"

	"github.com/fun-dev/cloud-api/config"
	"github.com/fun-dev/cloud-api/infrastructure/dbmodels"
	_ "github.com/go-sql-driver/mysql"
	"github.com/go-xorm/xorm"
)

func main() {
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
