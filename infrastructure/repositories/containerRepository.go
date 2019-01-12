package repositories

import (
	"fmt"

	corev1 "k8s.io/api/core/v1"

	"github.com/fun-dev/cloud-api/config"
	"github.com/fun-dev/cloud-api/domain/models"
	"github.com/fun-dev/cloud-api/infrastructure/repositories/interfaces"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/clientcmd"
)

type containerRepository struct{}

func NewContainerRepository() interfaces.IContainerRepository {
	return containerRepository{}
}

func (repo containerRepository) GetContainersByNamespace(namespace string) ([]models.Container, error) {
	// kubeconfigのパス取得
	kubeConfigPath := config.GetKubeConfigPath()

	// kubeconfig内の現在のコンテキストを利用する
	config, err := clientcmd.BuildConfigFromFlags("", kubeConfigPath)
	if err != nil {
		return nil, err
	}

	// クライアントセットを作成
	clientset, err := kubernetes.NewForConfig(config)
	if err != nil {
		return nil, err
	}

	// podの一覧取得
	pods, err := clientset.CoreV1().Pods(namespace).List(metav1.ListOptions{})
	if err != nil {
		return nil, err
	}

	// 返却用スライスの生成
	returnContainers := make([]models.Container, 0, len(pods.Items))

	for _, pod := range pods.Items {
		// pod内のコンテナ一覧取得
		containers := pod.Spec.Containers

		// コンテナが1つであることを期待する
		if len(containers) != 1 {
			continue
		}
		container := containers[0]

		item := models.Container{
			UID:         string(pod.GetUID()),
			ImageName:   container.Image,
			ConnectInfo: getWebSocketPath(pod.GetSelfLink()),
			Status:      getPodState(pod),
		}

		returnContainers = append(returnContainers, item)
	}

	return returnContainers, nil
}

func (repo containerRepository) CreateContainer(uniqueUserID, imageID string) (models.Container, error) {
	// TODO youtangai コンテナ作成の処理を記述
	return models.Container{}, nil
}

func (repo containerRepository) DeleteContainer(uniqueUserID string, containerID int64) error {
	// TODO youtangai コンテナ削除の処理を記述する
	return nil
}

func getWebSocketPath(selfLink string) string {
	return fmt.Sprintf("ws://%s%s/exec?container=my-nginx&stdin=1&stdout=1&stderr=1&tty=1&command=/bin/bash", config.GetKubeIP(), selfLink)
}

func getPodState(pod corev1.Pod) string {
	state := pod.Status.ContainerStatuses[0].State
	if state.Waiting != nil {
		return "creating"
	}
	if state.Running != nil {
		return "running"
	}
	if state.Terminated != nil {
		return "halted"
	}
	return ""
}
