package infrastructure

import (
	"github.com/fun-dev/cloud-api/infrastructure/repositories"
	"github.com/fun-dev/cloud-api/infrastructure/repositories/interfaces"
	"github.com/go-xorm/xorm"
)

var (
	engine        *xorm.Engine
	UserRepo      interfaces.IUserRepository
	ContainerRepo interfaces.IContainerRepository
)

func init() {
	var err error
	engine, err = repositories.NewEngine()
	if err != nil {
		panic(err)
	}

	// init user repo
	UserRepo = repositories.NewUserRepository(engine)
	// init container repo
	ContainerRepo = repositories.NewContainerRepository()
}
