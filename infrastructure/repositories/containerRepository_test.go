package repositories

import (
	"fmt"
	"testing"
)

func Test_containerRepository_GetContainersByNamespace(t *testing.T) {
	repo := NewContainerRepository()
	containers, err := repo.GetContainersByNamespace("default")
	if err != nil {
		t.Fatal(err)
	}
	fmt.Printf("%+v\n", containers)
}
