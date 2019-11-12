package provider

import (
	"flag"
	"github.com/fun-dev/ccms-poc/infrastructure/config"
	"github.com/fun-dev/ccms-poc/infrastructure/driver"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/clientcmd"
	"k8s.io/client-go/util/homedir"
	"log"
	"path/filepath"
)

var KubernetesProviderImpl *KubernetesProvider

func init() {
	KubernetesProviderImpl = &KubernetesProvider{}
	err := KubernetesProviderImpl.InitClient()
	if err != nil {
		log.Fatal(err)
	}
	KubernetesProviderImpl.InitKubectl()
	log.Printf("[debug] establised connection on KubernetesProvider.Init()\n")
}

type KubernetesProvider struct {
	Client  *kubernetes.Clientset
	Kubectl driver.IKubectlDriver
	Config  *config.AppVariableOnKubectl
}

func (d *KubernetesProvider) InitClient() error {
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
	d.Client, err = kubernetes.NewForConfig(config)
	if err != nil {
		return err
	}
	return nil
}

func (d *KubernetesProvider) InitKubectl() {
	d.Kubectl = &driver.KubectlDriver{Config:d.Config}
	d.Kubectl.Init()
}
