package application

import (
	"github.com/fun-dev/ccms-poc/application/controllers"
	"github.com/fun-dev/ccms-poc/application/controllers/interfaces"
)

var (
	ContainerCtrl interfaces.IContainerController
)

func init() {
	ContainerCtrl = controllers.NewContainerController()
}