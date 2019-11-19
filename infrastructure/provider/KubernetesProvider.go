package provider

import (
	"flag"
	"github.com/fun-dev/ccms-poc/infrastructure/driver"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/clientcmd"
	"k8s.io/client-go/util/homedir"
	"log"
	"path/filepath"
)

type KubernetesProvider struct {
	Client  *kubernetes.Clientset
	Kubectl driver.IKubectlDriver
}

func NewKubernetesProvider(k driver.IKubectlDriver) *KubernetesProvider {
	result := &KubernetesProvider{}
	result.Kubectl = k
	if err := result.InitClient(); err != nil {
		log.Fatal(err)
	}
	return result
}

func (d *KubernetesProvider) InitClient() error {
	var kubeconfig *string
	if home := homedir.HomeDir(); home != "" {
		kubeconfig = flag.String("kubeconfig", filepath.Join(home, ".kube", "config"), "(optional) absolute path to the kubeconfig file")
	} else {
		kubeconfig = flag.String("kubeconfig", "", "absolute path to the kubeconfig file")
	}
	flag.Parse()
	buildConfig, err := clientcmd.BuildConfigFromFlags("", *kubeconfig)
	if err != nil {
		return err
	}
	if d.Client, err = kubernetes.NewForConfig(buildConfig); err != nil {
		return err
	}
	return nil
}
