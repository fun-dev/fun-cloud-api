package cloudk8s

import (
	"errors"
	v1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

var (
	ErrRetrievePodsOnKubeapiClient = errors.New("failed retrieve pods on kubeapi client")
)

func GetPodsOnKubeAPIClient(namespace string) (*v1.PodList, error) {
	pods, err := _kubeapiClient.CoreV1().Pods(namespace).List(metav1.ListOptions{})
	if err != nil {
		return nil, ErrRetrievePodsOnKubeapiClient
	}
	return pods, nil
}