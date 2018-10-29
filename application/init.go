package application

import (
	"github.com/fun-dev/cloud-api/application/controllers"
	"github.com/fun-dev/cloud-api/application/controllers/interfaces"
)

var (
	UserController interfaces.IUserController
)

func init() {
	UserController = controllers.NewUserController()
}
