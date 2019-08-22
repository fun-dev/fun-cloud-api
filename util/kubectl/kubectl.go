package kubectl

import (
	apiv1 "k8s.io/api/core/v1"
)

type Kubectl interface {
	// [Official SDK]
	GetContainer(namespace string) (*apiv1.PodList, error)
	// [Wrapper Kubectl] Set pod name inside func
	CreateContainer(namespace, imageName string)
	DeleteContainer(namespace, podName, imageName string)
	ConfirmIsNamespaceExist(namespace string) (bool, error)
	// [The Other]
}
