package cloudutil

import (
	"io/ioutil"
	"os"
)

func CreateTmpManifest(manifestStr string) (string, error) {
	tmpFile, err := ioutil.TempFile(os.TempDir(), "manifest-*.yaml")
	if err != nil {
		return "", err
	}
	if _, err = tmpFile.Write([]byte(manifestStr)); err != nil {
		return "", err
	}
	return tmpFile.Name(), nil
}

func CreateTmpData(base64Decoded string, fileName string) (string, error) {
	tmpFile, err := ioutil.TempFile(os.TempDir(), "*." + fileName)
	if err != nil {
		return "", err
	}
	if _, err = tmpFile.Write([]byte(base64Decoded)); err != nil {
		return "", err
	}
	return tmpFile.Name(), nil
}
