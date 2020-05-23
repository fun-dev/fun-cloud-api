package cloudk8s

import (
	"errors"
	"github.com/fun-dev/fun-cloud-api/pkg/cloudutil"
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

var (
	ConfigMapYamlKey = "CONFIG_MAP_YAML_KEY"
	//
	ErrCreateConfigMapOnKubeapiClient   = errors.New("failed create config map on kubeapi client")
	ErrRetrieveConfigMapOnKubeapiClient = errors.New("failed retrieve config map on kubeapi client")
	ErrConfigMapHasNoValue              = errors.New("failed config map has no value")
)

func SaveManifestOnKubeAPIClient(configMapName, manifestStr string, namespace string) error {
	configMapData := make(map[string]string, 0)
	configMapData[ConfigMapYamlKey] = manifestStr

	configMap := corev1.ConfigMap{
		TypeMeta: metav1.TypeMeta{
			Kind:       "ConfigMap",
			APIVersion: "v1",
		},
		ObjectMeta: metav1.ObjectMeta{
			Name:      configMapName,
			Namespace: namespace,
		},
		Data: configMapData,
	}
	_, err := _kubeapiClient.CoreV1().ConfigMaps(namespace).Create(&configMap)
	if err != nil {
		return ErrCreateConfigMapOnKubeapiClient
	}
	return nil
}

func GetManifestOnKubeAPIClient(configMapName, namespace string) (string, error) {
	configMap, err := _kubeapiClient.CoreV1().ConfigMaps(namespace).Get(configMapName, metav1.GetOptions{})
	if err != nil {
		return "", ErrRetrieveConfigMapOnKubeapiClient
	}
	result := configMap.Data[ConfigMapYamlKey]
	if result != cloudutil.NullString {
		return "", ErrConfigMapHasNoValue
	}
	return result, nil
}

func DeleteManifestOnKubeAPIClient(configMapName, namespace string) error {
	err := _kubeapiClient.CoreV1().ConfigMaps(namespace).Delete(configMapName, &metav1.DeleteOptions{})
	if err != nil {
		return err
	}
	return nil
}

