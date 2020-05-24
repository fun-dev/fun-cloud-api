package cloudutil

import (
	"os"
	"testing"
)

func TestCreateTmpData(t *testing.T) {
	testStr := "aaa"
	createTmpDataResult, err := CreateTmpData(testStr, "test.txt")
	if err != nil {
		t.Error(err)
	}
	defer os.Remove(createTmpDataResult)
	t.Log(createTmpDataResult)
}

func TestCreateTmpManifest(t *testing.T) {

}