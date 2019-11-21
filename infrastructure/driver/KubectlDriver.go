package driver

import (
	"bytes"
	"errors"
	"fmt"
	"github.com/fun-dev/ccms/infrastructure/apperror/drivererr"
	"io/ioutil"
	"log"
	"os"
	"os/exec"

	"k8s.io/apimachinery/pkg/runtime/serializer/json"
	"k8s.io/client-go/kubernetes/scheme"

	apps "k8s.io/api/apps/v1"
	core "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/runtime"
	serialize "k8s.io/apimachinery/pkg/runtime/serializer"
)

var (
	_binaryPath             = os.Getenv("KUBECTL_BINARY_PATH")
	_deploymentManifestPath = os.Getenv("KUBECTL_DEPLOYMENT_MANIFEST_PATH")
	_kubeconfigPath         = os.Getenv("KUBECTL_CONFIG_PATH")
	// --- For Test --- //
	_currentDir, _ = os.Getwd()
	_defaultBinaryPath             = "/snap/bin/kubectl"
	_defaultKubeconfigPath         = _currentDir + "/kubeconfig"
	_defaultDeploymentManifestPath = "./manifest/pod-deployment-template.yaml"
)

const (
	KubectlOptionApply    = "KUBECTL_OPTION_APPLY"
	KubectlOptionDelete   = "KUBECTL_OPTION_DELETE"
	UseDeploymentManifest = "USE_DEPLOYMENT_MANIFEST"
)

// IKubectlDriver is
type IKubectlDriver interface {
	Execute(option string, yaml string, namespace string) error
	DeserializeYamlToObject(option string, targetObject runtime.Object) (interface{}, error)
	DecodeObjectToYaml(targetObject runtime.Object) (string, error)
	Init()
}

// KubectlDriver is
type KubectlDriver struct {
	apps.Deployment
	core.Service
	core.PersistentVolume
	core.PersistentVolumeClaim

	BinaryPath             string
	DeploymentManifestPath string
	KubeconfigPath         string
}

func NewKubectlDriver() IKubectlDriver {
	// --- create non value instance on memory --- //
	result := &KubectlDriver{}
	// --- load data  --- //
	result.Init()
	return result
}

func NewTestKubectlDriver() IKubectlDriver {
	result := &KubectlDriver{}
	result.DeploymentManifestPath = _defaultDeploymentManifestPath
	result.KubeconfigPath = _defaultKubeconfigPath
	result.BinaryPath = _defaultBinaryPath
	return result
}

func (d *KubectlDriver) Init() {
	d.BinaryPath = _binaryPath
	d.DeploymentManifestPath = _deploymentManifestPath
	d.KubeconfigPath = _kubeconfigPath
}

/*
Execute is executing kubectl command
@param option : kubectl 'apply' or 'delete'
@param namespace : kubectl '-n'
*/
func (d *KubectlDriver) Execute(option string, yaml string, namespace string) error {
	// --- gen yaml file --- //
	file, err := os.Create("tmp.yaml")
	if err != nil {
		return err
	}
	defer file.Close()
	defer func() {
		if err := os.Remove("tmp.yaml"); err != nil {
			fmt.Println(err)
		}
	}()
	file.Write(([]byte)(yaml))
	// --------------------- //
	var cmd *exec.Cmd
	var out bytes.Buffer
	var stderr bytes.Buffer
	// --- check kubectl command "Apply" or "Delete" --- //
	switch option {
	case KubectlOptionApply:
		args := []string{"--kubeconfig", d.KubeconfigPath, "apply", "-f", _currentDir + "/tmp.yaml", "-n", namespace}
		cmd = exec.Command(d.BinaryPath, args...)
	case KubectlOptionDelete:
		args := []string{"--kubeconfig", d.KubeconfigPath, "delete", "-f", _currentDir + "/tmp.yaml", "-n", namespace}
		cmd = exec.Command(d.BinaryPath, args...)
	default:
		return drivererr.OptionCanNotBeFoundOnKubectl
	}
	log.Printf("[debug]: generate command: %s", cmd.Args)
	log.Printf("[debug]: current directory is %s", _currentDir)
	cmd.Stdout = &out
	cmd.Stderr = &stderr
	// --- Uber Style Guide: Reduce Scope of Variables --- //
	if err := cmd.Run(); err != nil {
		return fmt.Errorf("[error]: call kubectl: %w", err)
	}
	return nil
}

/*
DeserializeYamlToObject is generating object from yaml file
@param: option is kind of yaml manifest example: deployment and service
@targetObject is struct resource example: Deployment and Service
*/
func (d KubectlDriver) DeserializeYamlToObject(option string, targetObject runtime.Object) (interface{}, error) {
	var yaml []byte
	var err error
	scheme := runtime.NewScheme()
	codecFactory := serialize.NewCodecFactory(scheme)
	deserializer := codecFactory.UniversalDeserializer()

	switch option {
	case UseDeploymentManifest:
		yaml, err = ioutil.ReadFile(d.DeploymentManifestPath)
		if err != nil {
			return nil, err
		}
	default:
		return nil, errors.New("enough option")
	}
	object, _, err := deserializer.Decode(yaml, nil, targetObject)
	if err != nil {
		return nil, err
	}
	return object, nil
}

// DecodeObjectToYaml is
func (d KubectlDriver) DecodeObjectToYaml(targetObject runtime.Object) (string, error) {
	serializer := json.NewSerializerWithOptions(json.DefaultMetaFactory, scheme.Scheme, scheme.Scheme, json.SerializerOptions{Yaml: true})
	var buffer bytes.Buffer
	if err := serializer.Encode(targetObject, &buffer); err != nil {
		return "", err
	}
	return buffer.String(), nil
}
