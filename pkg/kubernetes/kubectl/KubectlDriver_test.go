package kubectl

import (
	apps "k8s.io/api/apps/v1"
	"log"
	"testing"
)

func TestKubectlDriver_DeserializeYamlToObject(t *testing.T) {
	driver := NewTestKubectlDriver()
	object, err := driver.DeserializeYamlToObject(UseDeploymentManifest, &apps.Deployment{})
	if err != nil {
		panic(err)
	}
	deployment := object.(*apps.Deployment)
	actual := deployment.Name
	expected := "test"
	log.Printf("deployment name: %v", deployment.Name)
	if actual != expected {
		t.Errorf("got: %v\nwant: %v", actual, expected)
	}
}

func TestKubectlDriver_DecodeObjectToYaml(t *testing.T) {
	driver := NewTestKubectlDriver()
	object, err := driver.DeserializeYamlToObject(UseDeploymentManifest, &apps.Deployment{})
	if err != nil {
		panic(err)
	}
	deployment := object.(*apps.Deployment)
	yaml, err := driver.DecodeObjectToYaml(deployment)
	if err != nil {
		panic(err)
	}
	log.Printf("%s", yaml)
}

func TestKubectlDriver_ExecuteApply(t *testing.T) {
	driver := NewTestKubectlDriver()
	object, err := driver.DeserializeYamlToObject(UseDeploymentManifest, &apps.Deployment{})
	if err != nil {
		log.Fatal(err)
	}
	deployment := object.(*apps.Deployment)
	//deployment.Name = "tarako-chan"
	yaml, err := driver.DecodeObjectToYaml(deployment)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("yaml: %s", yaml)
	err = driver.Execute(KubectlOptionApply, yaml, "default")
	if err != nil {
		log.Fatal(err)
	}
}

func TestKubectlDriver_ExecuteDelete(t *testing.T) {
	driver := NewTestKubectlDriver()
	object, err := driver.DeserializeYamlToObject(UseDeploymentManifest, &apps.Deployment{})
	if err != nil {
		log.Fatal(err)
	}
	deployment := object.(*apps.Deployment)
	//deployment.Name = "tarako-chan"
	yaml, err := driver.DecodeObjectToYaml(deployment)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("yaml: %s", yaml)
	err = driver.Execute(KubectlOptionDelete, yaml, "default")
	if err != nil {
		log.Fatal(err)
	}
}