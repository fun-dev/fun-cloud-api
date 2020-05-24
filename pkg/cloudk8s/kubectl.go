package cloudk8s

import (
	"encoding/base64"
	"errors"
	"github.com/fun-dev/fun-cloud-api/pkg/cloudutil"
	"os"
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

// --- DATA TRANSFER --- //
func TransferDataToContainerOnKubectl(base64Data, fileNameWithExtension, containerID, namespace, copyToPath string) error {
	dec, err := base64.StdEncoding.DecodeString(base64Data)
	if err != nil {
		return err
	}
	tmpDataName, err := cloudutil.CreateTmpData(string(dec), fileNameWithExtension)
	if err != nil {
		return err
	}
	defer os.Remove(tmpDataName)
	_, err = exec.Command(_kubectlBinaryPath,
		"-n", namespace, "cp", tmpDataName, containerID+":"+copyToPath,
	).Output()
	if err != nil {
		return err
	}
	return nil
}

func TransferDataFromContainerOnKubectlByBase64(containerID, namespace, copyFromPath string) (string, error) {
	catResult, err := exec.Command(_kubectlBinaryPath,
		"-n", namespace, "exec", containerID, "cat", copyFromPath ,
	).Output()
	if err != nil {
		return "", err
	}
	return base64.StdEncoding.EncodeToString(catResult), nil
}
