package kubectl

import (
	"flag"
	"path/filepath"

	"github.com/fun-dev/ccms/domain/container/value"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/clientcmd"
	"k8s.io/client-go/util/homedir"
)

type IKubectlDriver interface {
	GetByNamespace(namespace value.K8SNamespace) error
	Apply(manifest value.K8SManifest) error
	Delete(manifest value.K8SManifest) error
}

type KubectlDriver struct {
	Clientset kubernetes.Clientset
}

func (d *KubectlDriver) init() error {
	var kubeconfig *string
	if home := homedir.HomeDir(); home != "" {
		kubeconfig = flag.String("kubeconfig", filepath.Join(home, ".kube", "config"), "(optional) absolute path to the kubeconfig file")
	} else {
		kubeconfig = flag.String("kubeconfig", "", "absolute path to the kubeconfig file")
	}
	flag.Parse()

	config, err := clientcmd.BuildConfigFromFlags("", *kubeconfig)
	if err != nil {
		return err
	}
	d.Clientset, err := kubernetes.NewForConfig(config)
	if err != nil {
		return err
	}
	return nil
}
