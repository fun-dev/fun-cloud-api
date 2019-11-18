package repoerr

import "errors"

type RepositoryError struct {
	DeploymentManifestCanNotBeFound error
}

var (
	DeploymentManifestCanNotBeFound = errors.New("deployment manifest can not be found")
)