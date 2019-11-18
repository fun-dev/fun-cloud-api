package apperror

import (
	"errors"
)


type ControllerError struct {
	ContainerIDCanNotBeFoundOnParam error
}

func NewControllerError() *ControllerError {
	result := new(ControllerError)
	result.Init()
	return result
}

func (e *ControllerError) Init() {
	e.ContainerIDCanNotBeFoundOnParam = errors.New("container id can not be found on param")
}