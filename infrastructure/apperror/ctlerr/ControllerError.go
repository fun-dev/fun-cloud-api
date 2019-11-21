package ctlerr

import (
	"errors"
)

var (
	ContainerIDCanNotBeFoundOnParam = errors.New("container id can not be found on param")
	AuthorizationIsNotFoundOnHeader = errors.New("authorization can not be found on header")
)
