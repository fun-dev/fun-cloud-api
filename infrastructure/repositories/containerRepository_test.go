package repositories

import (
	"fmt"
	"testing"
)

const (
	uniqueUserID = "108898913206828981410"
)

func Test_containerRepository_GetContainersByNamespace(t *testing.T) {
	repo := NewContainerRepository()
	containers, err := repo.GetContainersByNamespace("default")
	if err != nil {
		t.Fatal(err)
	}
	fmt.Printf("%+v\n", containers)
}

func Test_containerRepository_CreateContainer(t *testing.T) {
	repo := NewContainerRepository()
	err := repo.CreateContainer(uniqueUserID, "nginx:latest")
	if err != nil {
		t.Fatal(err)
	}
}

// func Test_containerRepository_DeleteContainer(t *testing.T) {
// 	repo := NewContainerRepository()
// 	err := repo.DeleteContainer(uniqueUserID, "20190112201521")
// 	if err != nil {
// 		t.Fatal(err)
// 	}
// }
