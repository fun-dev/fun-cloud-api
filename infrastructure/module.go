package infrastructure

import (
	"github.com/fun-dev/ccms-poc/domain/container/interfaces"
	"github.com/fun-dev/ccms-poc/infrastructure/repository"
)

var (
	ContainerRepo interfaces.IContainerRepository
)

func init() {
	ContainerRepo = repository.NewContainerRepository()
}
