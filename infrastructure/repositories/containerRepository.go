package repositories

import (
	"fmt"

	appsv1 "k8s.io/api/apps/v1"
	corev1 "k8s.io/api/core/v1"

	"github.com/fun-dev/cloud-api/config"
	"github.com/fun-dev/cloud-api/domain/models"
	"github.com/fun-dev/cloud-api/infrastructure/repositories/interfaces"
	"github.com/fun-dev/cloud-api/util"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/clientcmd"
)

const (
	labelName = "pod-name"
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
			UID:         pod.ObjectMeta.Labels[labelName], //本来podが持っているuidではない
			ImageName:   container.Image,
			ConnectInfo: getWebSocketPath(pod.GetSelfLink(), container.Name),
			Status:      getPodState(pod),
		}

		returnContainers = append(returnContainers, item)
	}

	return returnContainers, nil
}

func (repo containerRepository) CreateContainer(uniqueUserID, imageName string) error {
	// TODO youtangai コンテナ作成の処理を記述
	kubeConfigPath := config.GetKubeConfigPath()

	config, err := clientcmd.BuildConfigFromFlags("", kubeConfigPath)
	if err != nil {
		return err
	}

	clientset, err := kubernetes.NewForConfig(config)
	if err != nil {
		return err
	}

	// namespaceの一覧取得
	namespaces, err := clientset.CoreV1().Namespaces().List(metav1.ListOptions{})
	if err != nil {
		return err
	}

	namespaceSlice := namespaces.Items
	currentNameSpace := &corev1.Namespace{
		ObjectMeta: metav1.ObjectMeta{
			Name: uniqueUserID,
		},
	}

	// namespaceがなければ作る
	if !isNameSpaceExist(&namespaceSlice, currentNameSpace) {
		_, err := clientset.CoreV1().Namespaces().Create(currentNameSpace)
		if err != nil {
			return err
		}
	}

	deploymentsClient := clientset.AppsV1().Deployments(uniqueUserID)

	now := util.GetNowString()

	deployment := &appsv1.Deployment{
		ObjectMeta: metav1.ObjectMeta{
			Name: now,
		},
		Spec: appsv1.DeploymentSpec{
			Replicas: util.Int32Ptr(1),
			Selector: &metav1.LabelSelector{
				MatchLabels: map[string]string{
					labelName: now,
				},
			},
			Template: corev1.PodTemplateSpec{
				ObjectMeta: metav1.ObjectMeta{
					Labels: map[string]string{
						labelName: now,
					},
				},
				Spec: corev1.PodSpec{
					Containers: []corev1.Container{
						{
							Name:  now,
							Image: imageName,
						},
					},
				},
			},
		},
	}

	fmt.Println("namespace:", uniqueUserID, "imagename:", imageName)
	_, err = deploymentsClient.Create(deployment)
	if err != nil {
		return err
	}

	return nil
}

func (repo containerRepository) DeleteContainer(uniqueUserID, containerID string) error {
	// TODO youtangai コンテナ削除の処理を記述する
	kubeConfigPath := config.GetKubeConfigPath()

	config, err := clientcmd.BuildConfigFromFlags("", kubeConfigPath)
	if err != nil {
		return err
	}

	clientset, err := kubernetes.NewForConfig(config)
	if err != nil {
		return err
	}

	deploymentsClient := clientset.AppsV1().Deployments(uniqueUserID)

	deletePolicy := metav1.DeletePropagationForeground
	err = deploymentsClient.Delete(containerID, &metav1.DeleteOptions{
		PropagationPolicy: &deletePolicy,
	})
	if err != nil {
		return err
	}

	return nil
}

func getWebSocketPath(selfLink, container string) string {
	return fmt.Sprintf("ws://%s?connect_info=wss://%s%s/exec?container=%s%%26stdin=1%%26stdout=1%%26stderr=1%%26tty=1%%26command=/bin/bash", config.GetProxyAddr(), config.GetKubeIP(), selfLink, container)
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

func isNameSpaceExist(list *[]corev1.Namespace, current *corev1.Namespace) bool {
	for _, item := range *list {
		if item.ObjectMeta.Name == current.ObjectMeta.Name {
			return true
		}
	}
	return false
}
