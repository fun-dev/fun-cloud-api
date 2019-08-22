package domain

import (
	"github.com/fun-dev/cloud-api/domain/services"
	"github.com/fun-dev/cloud-api/domain/services/interfaces"
)

var (
	UserSrv      interfaces.IUserService
	ContainerSrv interfaces.IContainerService
)

func init() {
	var err error

	if err != nil {
		panic(err)
	}

	UserSrv = services.NewUserService()
	ContainerSrv = services.NewContainerService()
}
