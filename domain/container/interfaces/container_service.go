package interfaces


type IContainerService interface {
	GetContainersByUniqueUserID(uniqueUserID string) (string, error)
	CreateContainer(uniqueUserID, imageName string) error
	DeleteContainer(uniqueUserID, containerID string) error
}
