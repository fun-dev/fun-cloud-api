package application

import (
	"github.com/fun-dev/cloud-api/application/controllers"
	"github.com/fun-dev/cloud-api/application/controllers/interfaces"
)

var (
	UserController  interfaces.IUserController
	ImageController controllers.ImageController
)

func init() {
	UserController = controllers.NewUserController()
	ImageController = controllers.ImageController{}
}
