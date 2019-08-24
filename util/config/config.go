package config

import (
	"encoding/json"
	"github.com/kelseyhightower/envconfig"
	"io/ioutil"
	v1 "k8s.io/api/apps/v1"
	v12 "k8s.io/api/core/v1"
	"os"
)

type externalVariables struct {
	// Kubernetes environment
	KubectlPath    string `envconfig:"K8S_KUBECTL_PATH"`
	KubeConfigPath string `envconfig:"K8S_CONFIG_PATH"`
	KubeIP         string `envconfig:"K8S_IP"`
	ProxyAddress   string `envconfig:"PROXY_ADDRESS"`
	// Redis environment
	RedisAddress  string `envconfig:"REDIS_ADDRESS"`
	RedisPassword string `envconfig:"REDIS_PASSWORD"`
	RedisDatabase string `envconfig:"REDIS_DATABASE"`
	//
	APPFilePath string `envconfig:"APP_FILE_PATH"`
}

var (
	Ext            externalVariables
	DeployTemplate v1.Deployment
	SvcTemplate    v12.Service
	PVTemplate     v12.PersistentVolume
	PVCTemplate    v12.PersistentVolumeClaim
)

func (ext *externalVariables) CreateK8SWebsocketProxyPath() (string, error) {
	return "", nil
}

func init() {
	envconfig.MustProcess("", &Ext)
	// ---Init Deployment---
	deploymentJSON, err := readJson(Ext.APPFilePath + "/manifest/json/pod-deployment-template.json")
	if err != nil {
		panic(err)
	}
	deploymentByte, err := bindJson(deploymentJSON)
	if err != nil {
		panic(err)
	}
	err = json.Unmarshal(deploymentByte, &DeployTemplate)
	if err != nil {
		panic(err)
	}
	// ---Init Deployment---
	// ---Init Service---
	serviceJSON, err := readJson(Ext.APPFilePath + "/manifest/json/pod-service-template.json")
	if err != nil {
		panic(err)
	}
	serviceByte, err := bindJson(serviceJSON)
	if err != nil {
		panic(err)
	}
	err = json.Unmarshal(serviceByte, &SvcTemplate)
	if err != nil {
		panic(err)
	}
	// ---Init Service---
}

func readJson(filePath string) (*os.File, error) {
	jsonFile, err := os.Open(filePath)
	if err != nil {
		return nil, err
	}
	return jsonFile, nil
}

func bindJson(jsonFile *os.File) ([]byte, error) {
	byteValue, err := ioutil.ReadAll(jsonFile)
	if err != nil {
		return nil, err
	}
	return byteValue, nil
}

func ConvertToJSON(jsonModel interface{}) (*string, error){
	result, err := json.Marshal(jsonModel)
	if err != nil {
		return nil, err
	}
	strResult := string(result)
	return &strResult, nil
}