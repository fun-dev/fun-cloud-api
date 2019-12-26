package kubectl

import (
	"bytes"
	"errors"
	"fmt"
	"github.com/fun-dev/fun-cloud-api/pkg/logging"
	"io/ioutil"
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
	_currentDir, _                 = os.Getwd()
	_defaultBinaryPath             = "/snap/bin/kubectl"
	_defaultKubeconfigPath         = _currentDir + "/kubeconfig"
	_defaultDeploymentManifestPath = "./manifest/pod-deployment-template.yaml"
)

const (
	Apply                 = "KUBECTL_OPTION_APPLY"
	Delete                = "KUBECTL_OPTION_DELETE"
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
type Driver struct {
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
	result := &Driver{}
	// --- load data  --- //
	result.Init()
	return result
}

func NewTestKubectlDriver() IKubectlDriver {
	result := &Driver{}
	result.DeploymentManifestPath = _defaultDeploymentManifestPath
	result.KubeconfigPath = _defaultKubeconfigPath
	result.BinaryPath = _defaultBinaryPath
	return result
}

func (d *Driver) Init() {
	d.BinaryPath = _binaryPath
	d.DeploymentManifestPath = _deploymentManifestPath
	d.KubeconfigPath = _kubeconfigPath
}

/*
Execute is executing kubectl command
@param option : kubectl 'apply' or 'delete'
@param namespace.yaml : kubectl '-n'
*/
func (d *Driver) Execute(option string, yaml string, namespace string) error {
	// --- gen yaml file --- //
	file, err := os.Create("tmp.yaml")
	if err != nil {
		return err
	}
	// TODO: implement error handling
	defer file.Close()
	defer func() {
		if err := os.Remove("tmp.yaml"); err != nil {
			fmt.Println(err)
		}
	}()
	// TODO: implement error handling
	file.Write(([]byte)(yaml))
	// --------------------- //
	var cmd *exec.Cmd
	var out bytes.Buffer
	var stderr bytes.Buffer
	// --- check kubectl command "Apply" or "Delete" --- //
	switch option {
	case Apply:
		args := []string{"--kubeconfig", d.KubeconfigPath, "apply", "-f", _currentDir + "/tmp.yaml", "-n", namespace}
		cmd = exec.Command(d.BinaryPath, args...)
	case Delete:
		args := []string{"--kubeconfig", d.KubeconfigPath, "delete", "-f", _currentDir + "/tmp.yaml", "-n", namespace}
		cmd = exec.Command(d.BinaryPath, args...)
	default:
		return OptionCanNotBeFoundOnKubectl
	}
	logging.Logf("debug: generate command: %s", cmd.Args)
	logging.Logf("debug: current directory is %s", _currentDir)
	cmd.Stdout = &out
	cmd.Stderr = &stderr
	// --- Uber Style Guide: Reduce Scope of Variables --- //
	if err := cmd.Run(); err != nil {
		return fmt.Errorf("error: call kubectl: %w", err)
	}
	return nil
}

/*
DeserializeYamlToObject is generating object from yaml file
@param: option is kind of yaml manifest example: deployment and service
@targetObject is struct resource example: Deployment and Service
*/
func (d Driver) DeserializeYamlToObject(option string, targetObject runtime.Object) (interface{}, error) {
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
func (d Driver) DecodeObjectToYaml(targetObject runtime.Object) (string, error) {
	serializer := NewSerializerWithOptions(DefaultMetaFactory, scheme.Scheme, scheme.Scheme, SerializerOptions{Yaml: true})
	var buffer bytes.Buffer
	if err := serializer.Encode(targetObject, &buffer); err != nil {
		return "", err
	}
	return buffer.String(), nil
}
