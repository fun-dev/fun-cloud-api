package cloudk8s

import (
	"errors"
	"os/exec"
)

type KubectlOption int

const (
	APPLY KubectlOption = iota
	DELETE
)

var (
	ErrExecuteKubectlBinary    = errors.New("failed execute kubectl binary")
	ErrKubectlOptionHasNoValue = errors.New("failed kubectl option has no value")
)

func ExecuteManifestOnKubectl(manifestPath string, option KubectlOption) error {
	switch option {
	case APPLY:
		_, err := exec.Command(_kubectlBinaryPath, "-f", manifestPath, "apply").Output()
		if err != nil {
			return ErrExecuteKubectlBinary
		}
	case DELETE:
		_, err := exec.Command(_kubectlBinaryPath, "-f", manifestPath, "delete").Output()
		if err != nil {
			return ErrExecuteKubectlBinary
		}
	default:
		return ErrKubectlOptionHasNoValue
	}
	return nil
}

func ExecuteListCmdOnKubectl(dirPath, containerName, namespace string) (string, error) {
	result, err := exec.Command(_kubectlBinaryPath,
		"-n", namespace,
		"exec", containerName, "--", "ls", "-p", "-1", dirPath,
	).Output()
	if err != nil {
		return "", err
	}
	return string(result), nil
}