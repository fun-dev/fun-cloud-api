package kubernetes

import (
	"github.com/tozastation/kw"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/clientcmd"
	"os"
)

// Kubernetes Wrapper Class
type (
	IK8SProvider interface {
		Client() *kubernetes.Clientset
		Kubectl() *kw.Kubectl
		Manifest() *kw.Manifest
		InitKubectl(binaryPath, kubeConfigPath string) error
		InitK8SClient() (err error)
	}
	Provider struct {
		client   *kubernetes.Clientset
		kubectl  *kw.Kubectl
		manifest *kw.Manifest
	}
)

func NewKubernetesProvider() IK8SProvider {
	result := &Provider{}
	// --- inject object --- //
	result.manifest, _ = kw.NewManifest()
	return result
}

func (p *Provider) InitKubectl(binaryPath, kubeConfigPath string) error {
	var err error
	p.kubectl, err = kw.New(binaryPath, kubeConfigPath)
	if err != nil {
		return err
	}
	return nil
}

//TODO: add kubeConfigPath in args
func (p *Provider) InitK8SClient() (err error) {
	// use the current context in kubeconfig
	kubeConfigPath := os.Getenv("KUBECONFIG")
	config, err := clientcmd.BuildConfigFromFlags("", kubeConfigPath)
	if err != nil {
		return
	}
	p.client, err = kubernetes.NewForConfig(config)
	if err != nil {
		return
	}
	return
}

func (p *Provider) Client() *kubernetes.Clientset {
	return p.client
}

func (p *Provider) Kubectl() *kw.Kubectl {
	return p.kubectl
}

func (p *Provider) Manifest() *kw.Manifest {
	return p.manifest
}
