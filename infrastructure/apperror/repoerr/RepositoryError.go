package apperror

import "errors"

type RepositoryError struct {
	DeploymentManifestCanNotBeFound error
}

func NewRepositoryError() *RepositoryError {
	result := new(RepositoryError)
	result.Init()
	return result
}

func (e *RepositoryError) Init() {
	e.DeploymentManifestCanNotBeFound = errors.New("deployment manifest can not be found")
}