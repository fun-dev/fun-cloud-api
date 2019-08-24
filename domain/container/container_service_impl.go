package container

import (
	"github.com/fun-dev/ccms-poc/domain/container/interfaces"
	"github.com/fun-dev/ccms-poc/util/config"
	"github.com/fun-dev/ccms-poc/util/kubectl"
)

type containerService struct {
	Repo    interfaces.IContainerRepository
	Kubectl kubectl.IKubectl
}

func NewContainerService(repo interfaces.IContainerRepository, kubectl kubectl.IKubectl) interfaces.IContainerService {
	return &containerService{
		Repo:    repo,
		Kubectl: kubectl,
	}
}

func (c *containerService) GetContainersByUserID(userID, podName string) ([]*Container, error) {
	targetNamespace := userID
	pods, err := c.Kubectl.GetContainer(targetNamespace)
	if err != nil {
		return nil, err
	}
	result := make([]*Container, len(pods.Items))
	topContainer := 0
	for index, pod := range pods.Items {
		result[index] = &Container{
			UID:         targetNamespace,
			ImageName:   pod.Spec.Containers[topContainer].Image,
			PodName:     pod.Name,
			ConnectInfo: config.Ext.CreateK8SWebsocketProxyPath(pod.ObjectMeta.SelfLink),
			Status:      pod.Status.ContainerStatuses[topContainer].State.String(),
		}
	}
	return result, nil
}

func (c containerService) CreateContainer(userID, imageName string) error {
	targetNamespace := userID
	err := c.Kubectl.CreateContainer(targetNamespace, imageName)
	if err != nil {
		return err
	}
	return nil
}

func (c containerService) DeleteContainer(userID, podName string) error {
	targetNamespace := userID
	err := c.Kubectl.DeleteContainer(targetNamespace, podName)
	if err != nil {
		return err
	}
	return nil
}
