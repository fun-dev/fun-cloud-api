package driver

import (
	"bytes"
	"errors"
	"github.com/fun-dev/ccms-poc/infrastructure/config"
	"io/ioutil"
	"os/exec"

	"k8s.io/apimachinery/pkg/runtime/serializer/json"
	"k8s.io/client-go/kubernetes/scheme"

	apps "k8s.io/api/apps/v1"
	core "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/runtime/serializer"
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
	BinaryPath       string
	DeploymentObject apps.Deployment
	SvcObject        core.Service
	PVObject         core.PersistentVolume
	PVCObject        core.PersistentVolumeClaim
	Config           *config.AppVariableOnKubectl
}

func (d *KubectlDriver) Init() {
	d.BinaryPath = d.Config.BinaryPath
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
		return errors.New("option is not found")
	}
	cmd.Stdout = &out
	cmd.Stderr = &stderr
	err := cmd.Run()
	if err != nil {
		return errors.New(stderr.String())
	}
	return nil
}

// DeserializeYamlToObject is
func (d *KubectlDriver) DeserializeYamlToObject(filePath string, targetObject runtime.Object) (interface{}, error) {
	scheme := runtime.NewScheme()
	codecFactory := serializer.NewCodecFactory(scheme)
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
	err := serializer.Encode(targetObject, &buffer)
	if err != nil {
		return "", err
	}
	return buffer.String(), nil
}
