package infrastructure

import (
	"fmt"

	"github.com/fun-dev/cloud-api/config"

	"github.com/fun-dev/cloud-api/infrastructure/repositories"
	"github.com/fun-dev/cloud-api/infrastructure/repositories/interfaces"
	_ "github.com/go-sql-driver/mysql"
	"github.com/go-xorm/xorm"
)

var (
	UserRepo interfaces.IUserRepository
)

func init() {
	var err error

	UserRepo, err = initUserRepo()
	if err != nil {
		panic(err)
	}
}

func initUserRepo() (interfaces.IUserRepository, error) {
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
	return repositories.NewUserRepository(engine), nil
}
