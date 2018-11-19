package interfaces

import "github.com/fun-dev/cloud-api/infrastructure/dbmodels"

type IContainerRepository interface {
	GetContainersByNamespace(namespace string) []dbmodels.Container
}
