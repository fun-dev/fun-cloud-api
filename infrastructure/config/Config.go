package config

import (
	"errors"
	"os"
)

var (
	AppVariableOnKubectlImpl *AppVariableOnKubectl
	AppVariableOnRedisImpl   *AppVariableOnRedis
)

type AppVariableOnKubectl struct {
	BinaryPath                string
	DeploymentManifestPath    string
	ServiceManifestPath       string
	PersistentManifestPath    string
	PersistentVolumeClaimPath string
}

// Load is loading config
func (k *AppVariableOnKubectl) Load() error {
	// --- Binary Path --- //
	k.BinaryPath = os.Getenv("KUBECTL_BINARY_PATH")
	if k.BinaryPath == "" {
		return errors.New("binary path can not be found")
	}
	return nil
}

type AppVariableOnRedis struct {
	Address  string
	Password string
	DBName   int
}
