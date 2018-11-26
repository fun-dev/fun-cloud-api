package infrastructure

import (
	"github.com/fun-dev/cloud-api/infrastructure/repositories"
	"github.com/fun-dev/cloud-api/infrastructure/repositories/interfaces"
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
	engine, err := repositories.NewEngine()
	if err != nil {
		return nil, err
	}
	return repositories.NewUserRepository(engine), nil
}
