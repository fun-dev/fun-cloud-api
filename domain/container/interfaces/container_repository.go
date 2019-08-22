package interfaces

type IContainerRepository interface {
	SaveDeploymentManifest(namespace, podName, yaml string) (bool, error)
	SaveServiceManifest(namespace, podName, yaml string) (bool, error)
	DeleteDeploymentManifest(namespace, podName string) (bool, error)
	DeleteServiceManifest(namespace, podName string) (bool, error)
}