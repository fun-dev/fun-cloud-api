package cloudk8s

import (
	"errors"
	v1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

var (
	ErrNamespaceIsNotFound = errors.New("failed namespace is not found")
)

func IsExistNamespaceOnKubeAPIClient(namespace string) error {
	ns, err := _kubeapiClient.CoreV1().Namespaces().Get(namespace, metav1.GetOptions{})
	if err != nil {
		return err
	}
	if ns.Status.Conditions[0].Status == v1.ConditionFalse {
		return ErrNamespaceIsNotFound
	}
	return nil
}

func CreateNamespaceOnKubeAPIClient(namespace string) error {
	_, err := _kubeapiClient.CoreV1().Namespaces().Create(&v1.Namespace{
		TypeMeta:   metav1.TypeMeta{},
		ObjectMeta: metav1.ObjectMeta{
			Name: namespace,
		},
		Spec:       v1.NamespaceSpec{},
		Status:     v1.NamespaceStatus{},
	})
	if err != nil {
		return err
	}
	return nil
}