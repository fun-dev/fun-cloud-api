package cloudk8s

import (
	"bytes"
	"errors"
	"io/ioutil"
	v1 "k8s.io/api/apps/v1"
	v12 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/runtime/serializer"
	"k8s.io/apimachinery/pkg/runtime/serializer/json"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/kubernetes/scheme"
	"k8s.io/client-go/rest"
	"k8s.io/client-go/tools/clientcmd"
	"os"
)

var (
	_kubectlBinaryPath              = os.Getenv("KUBECTL_BINARY_PATH")
	_clusterAdminKubeconfigPath     = os.Getenv("CLUSTER_ADMIN_KUBECONFIG_PATH")
	_templateDeploymentManifestPath = os.Getenv("TEMPLATE_DEPLOYMENT_MANIFEST_PATH")
	_templatePVManifestPath         = os.Getenv("TEMPLATE_PV_MANIFEST_PATH")
	_templatePVCManifestPath        = os.Getenv("TEMPLATE_PVC_MANIFEST_PATH")
	_templateServiceManifestPath    = os.Getenv("TEMPLATE_SERVICE_MANIFEST_PATH")
	//
	_configObj     *rest.Config
	_kubeapiClient *kubernetes.Clientset
	// Error
	ErrTemplateDeploymentManifestIsNotFound = errors.New("failed template deployment manifest is not found")
	ErrTemplatePVManifestIsNotFound         = errors.New("failed template pv manifest is not found")
	ErrTemplatePVCManifestIsNotFound        = errors.New("failed template pvc manifest is not found")
	ErrTemplateServiceManifestIsNotFound    = errors.New("failed template service manifest is not found")
	ErrReadTemplateManifestOptionIsNotFound = errors.New("failed read template manifest option is not found")
)

type ManifestType int

const (
	DEPLOYMENT ManifestType = iota
	PV
	PVC
	SERVICE
)

func init() {
	_configObj, _ = clientcmd.BuildConfigFromFlags("", _clusterAdminKubeconfigPath)
	_kubeapiClient, _ = kubernetes.NewForConfig(_configObj)
}

func GetTemplateDeploymentObject() (*v1.Deployment, error) {
	var result interface{}
	var err error
	result, err = _transformYamlToObject(DEPLOYMENT)
	if err != nil {
		return nil, err
	}
	castResult := result.(v1.Deployment)
	return &castResult, nil
}

func GetTemplatePVObject() (*v12.PersistentVolume, error) {
	var result interface{}
	var err error
	result, err = _transformYamlToObject(PV)
	if err != nil {
		return nil, err
	}
	castResult := result.(v12.PersistentVolume)
	return &castResult, nil
}

func GetTemplatePVCObject() (*v12.PersistentVolumeClaim, error) {
	var result interface{}
	var err error
	result, err = _transformYamlToObject(PVC)
	if err != nil {
		return nil, err
	}
	castResult := result.(v12.PersistentVolumeClaim)
	return &castResult, nil
}

func GetManifestFromObject(item runtime.Object) (string, error) {
	s := json.NewSerializerWithOptions(json.DefaultMetaFactory, scheme.Scheme, scheme.Scheme, json.SerializerOptions{Yaml: true})
	var buffer bytes.Buffer
	if err := s.Encode(item, &buffer); err != nil {
		return "", err
	}
	return buffer.String(), nil
}

func _transformYamlToObject(manifestType ManifestType) (interface{}, error) {
	s := runtime.NewScheme()
	codecFactory := serializer.NewCodecFactory(s)
	deserializer := codecFactory.UniversalDeserializer()
	var result runtime.Object
	switch manifestType {
	case DEPLOYMENT:
		yamlStr, err := _readTemplateManifest(DEPLOYMENT)
		if err != nil {
			return nil, err
		}
		result, _, err = deserializer.Decode([]byte(yamlStr), nil, &v1.Deployment{})
		if err != nil {
			return nil, err
		}
	case PV:
		yamlStr, err := _readTemplateManifest(PV)
		if err != nil {
			return nil, err
		}
		result, _, err = deserializer.Decode([]byte(yamlStr), nil, &v12.PersistentVolume{})
		if err != nil {
			return nil, err
		}
	case PVC:
		yamlStr, err := _readTemplateManifest(PVC)
		if err != nil {
			return nil, err
		}
		result, _, err = deserializer.Decode([]byte(yamlStr), nil, &v12.PersistentVolumeClaim{})
		if err != nil {
			return nil, err
		}
	case SERVICE:
		//TODO: Serviceの構造体からYamlマニフェストを生成
		return "", nil
	default:
		return nil, ErrReadTemplateManifestOptionIsNotFound
	}
	return result, nil
}

func _readTemplateManifest(manifestType ManifestType) (string, error) {
	var data []byte
	var err error
	switch manifestType {
	case DEPLOYMENT:
		data, err = ioutil.ReadFile(_templateDeploymentManifestPath)
		if err != nil {
			return "", ErrTemplateDeploymentManifestIsNotFound
		}
	case PV:
		data, err = ioutil.ReadFile(_templatePVManifestPath)
		if err != nil {
			return "", ErrTemplatePVManifestIsNotFound
		}
	case PVC:
		data, err = ioutil.ReadFile(_templatePVCManifestPath)
		if err != nil {
			return "", ErrTemplatePVCManifestIsNotFound
		}
	case SERVICE:
		data, err = ioutil.ReadFile(_templateServiceManifestPath)
		if err != nil {
			return "", ErrTemplateServiceManifestIsNotFound
		}
	default:
		return "", ErrReadTemplateManifestOptionIsNotFound
	}
	return string(data), nil
}
