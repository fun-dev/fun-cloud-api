package config

import "github.com/kelseyhightower/envconfig"

type externalVariables struct {
	KubectlPath    string `envconfig:"K8S_KUBECTL_PATH"`
	KubeConfigPath string `envconfig:"K8S_CONFIG_PATH"`
	KubeIP         string `envconfig:"K8S_IP"`
	ProxyAddress   string `envconfig:"PROXY_ADDRESS"`
}

var (
	Ext externalVariables
)

func init() {
	envconfig.MustProcess("", &Ext)
}

func (ext *externalVariables) CreateK8SWebsocketProxyPath() (string, error) {
	return "", nil
}