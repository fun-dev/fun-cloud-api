package container

// ContainerRepository is interface
type Repository interface {
	// Kubernetes
	GetAllByUserID(userID string) ([]*Container, error)
	Create(userID, imageName string) (containerID string, manifest string, err error)
	DeleteByContainerID(userID, containerID string) error
	// Store
	GetDeploymentManifestByContainerID(containerID string) (manifest string, err error)
	SaveDeploymentManifestByContainerID(containerID, manifest string) error
	DeleteDeploymentManifestByContainerID(containerID string) error
}
