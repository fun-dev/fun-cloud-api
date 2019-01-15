package application

import (
	"github.com/fun-dev/cloud-api/application/controllers"
	"github.com/fun-dev/cloud-api/application/controllers/interfaces"
)

var (
	UserController      interfaces.IUserController
	ContainerController interfaces.IContainerController
	ImageController     controllers.ImageController
)

func init() {
	UserController = controllers.NewUserController()
	ContainerController = controllers.NewContainerController()
	ImageController = controllers.ImageController{}
}
