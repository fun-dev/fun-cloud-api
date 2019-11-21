package kubernetes

import (
	"flag"
	"github.com/fun-dev/fun-cloud-api/pkg/kubernetes/kubectl"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/clientcmd"
	"k8s.io/client-go/util/homedir"
	"log"
	"path/filepath"
)

// Kubernetes Wrapper Class
type (
	IK8SProvider interface {
		Client() kubernetes.Clientset
		Kubectl() kubectl.IKubectlDriver
	}
	Provider struct {
		client  *kubernetes.Clientset
		kubectl kubectl.IKubectlDriver
	}
)

func NewKubernetesProvider() *Provider {
	result := &Provider{}
	// --- inject object --- //
	result.kubectl = kubectl.NewKubectlDriver()
	if err := result.InitClient(); err != nil {
		log.Fatal(err)
	}
	return result
}

func (p *Provider) InitClient() error {
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
	if p.client, err = kubernetes.NewForConfig(buildConfig); err != nil {
		return err
	}
	return nil
}

func (p Provider) Client() kubernetes.Clientset {
	return *p.client
}

func (p Provider) Kubectl() kubectl.IKubectlDriver {
	return p.kubectl
}
