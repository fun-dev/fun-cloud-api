package domain

import (
	"github.com/fun-dev/ccms-poc/domain/container"
	"github.com/fun-dev/ccms-poc/domain/container/interfaces"
)

var (
	ContainerSrv interfaces.IContainerService
)

func init() {
	ContainerSrv = container.NewContainerService()
}
