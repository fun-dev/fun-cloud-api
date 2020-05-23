package cloudutil

import (
	"io/ioutil"
	"os"
)

func CreateTmpManifest(manifestStr string) (string, error) {
	tmpFile, err := ioutil.TempFile(os.TempDir(), "prefix-")
	if err != nil {
		return "", err
	}
	if _, err = tmpFile.Write([]byte(manifestStr)); err != nil {
		return "", err
	}
	return tmpFile.Name(), nil
}