package repositories

import (
	"fmt"

	"github.com/fun-dev/cloud-api/config"
	_ "github.com/go-sql-driver/mysql"
	"github.com/go-xorm/xorm"
)

// NewEngine is return New Xorm RDB Engine
func NewEngine() (*xorm.Engine, error) {
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
		return nil, err
	}
	return engine, nil
}
