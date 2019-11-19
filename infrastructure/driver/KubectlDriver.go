package driver

import (
	"bytes"
	"fmt"
	"github.com/fun-dev/ccms/infrastructure/apperror/drivererr"
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
	//_serviceManifestPath       string
	//_persistentManifestPath    string
	//_persistentVolumeClaimPath string
)

const (
	KubectlOptionApply  = "KUBECTL_OPTION_APPLY"
	KubectlOptionDelete = "KUBECTL_OPTION_DELETE"
)

// IKubectlDriver is
type IKubectlDriver interface {
	Execute(option string, json string, namespace string) error
	DeserializeYamlToObject(filePath string, targetObject runtime.Object) (interface{}, error)
	DecodeObjectToYaml(targetObject runtime.Object) (string, error)
	Init()
}

// KubectlDriver is
type KubectlDriver struct {
	apps.Deployment
	core.Service
	core.PersistentVolume
	core.PersistentVolumeClaim

	BinaryPath string
}

func NewKubectlDriver() IKubectlDriver {
	// --- create non value instance on memory --- //
	result := &KubectlDriver{}
	// --- load data  --- //
	result.Init()
	return result
}

func (d *KubectlDriver) Init() {
	d.BinaryPath = _binaryPath
}

// Execute = execute kubectl command
func (d *KubectlDriver) Execute(option string, json string, namespace string) error {
	args := []string{"-y", json, "-n", namespace}
	var cmd *exec.Cmd
	var out bytes.Buffer
	var stderr bytes.Buffer
	// --- Check kubectl command "Apply" or "Delete" --- //
	switch option {
	case KubectlOptionApply:
		cmd = exec.Command("apply", args...)
	case KubectlOptionDelete:
		cmd = exec.Command("delete", args...)
	default:
		return drivererr.OptionCanNotBeFoundOnKubectl
	}
	cmd.Stdout = &out
	cmd.Stderr = &stderr
	// --- Uber Style Guide: Reduce Scope of Variables --- //
	if err := cmd.Run(); err != nil {
		return fmt.Errorf("execute kubectl: %w", err)
	}
	return nil
}

// DeserializeYamlToObject is
func (d *KubectlDriver) DeserializeYamlToObject(filePath string, targetObject runtime.Object) (interface{}, error) {
	scheme := runtime.NewScheme()
	codecFactory := serialize.NewCodecFactory(scheme)
	deserializer := codecFactory.UniversalDeserializer()
	yaml, err := ioutil.ReadFile(filePath)
	if err != nil {
		return nil, err
	}
	object, _, err := deserializer.Decode(yaml, nil, targetObject)
	if err != nil {
		return nil, err
	}
	return object, nil
}

// DecodeObjectToYaml is
func (d *KubectlDriver) DecodeObjectToYaml(targetObject runtime.Object) (string, error) {
	serializer := json.NewYAMLSerializer(json.DefaultMetaFactory, scheme.Scheme, scheme.Scheme)
	var buffer bytes.Buffer
	if err := serializer.Encode(targetObject, &buffer); err != nil {
		return "", err
	}
	return buffer.String(), nil
}
