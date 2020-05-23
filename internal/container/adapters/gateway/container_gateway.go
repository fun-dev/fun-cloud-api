package gateway

import (
	"context"
	"errors"
	"fmt"
	"github.com/fun-dev/fun-cloud-api/internal/container/domain/models"
	"github.com/fun-dev/fun-cloud-api/internal/container/domain/repository"
	"github.com/fun-dev/fun-cloud-api/pkg/cloudk8s"
	"github.com/fun-dev/fun-cloud-api/pkg/cloudstore"
	"github.com/fun-dev/fun-cloud-api/pkg/cloudutil"
	"go.mongodb.org/mongo-driver/bson"
	core "k8s.io/api/core/v1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"log"
	"time"
)

var (
	_collectionName = "manifest"
)

// ContainerGateway is
type ContainerGateway struct {
	K8SProvider cloudk8s.IK8SProvider
	Mongo       cloudstore.IMongoDriver
}

func NewContainerGateway(k cloudk8s.IK8SProvider, m cloudstore.IMongoDriver) repository.Repository {
	return &ContainerGateway{
		k,
		m,
	}
}

func (g ContainerGateway) GetAllByUserID(userID string) ([]*models.Container, error) {
	deploymentList, err := g.K8SProvider.Client().AppsV1().Deployments(userID).List(v1.ListOptions{})
	if err != nil {
		return nil, err
	}
	if len(deploymentList.Items) == cloudutil.HasNoItem {
		return nil, errors.New("error: you don't have any containers")
	}
	resultContainers := make([]*models.Container, len(deploymentList.Items))
	for i, d := range deploymentList.Items {
		podContainer := d.Spec.Template.Spec.Containers[0]
		resultContainer := &models.Container{
			UID:         podContainer.Name,
			ImageName:   podContainer.Image,
			ConnectInfo: d.SelfLink, // TODO: fix deprecated
			Status:      d.Status.String(),
		}
		resultContainers[i] = resultContainer
	}
	return resultContainers, nil
}

func (g ContainerGateway) Create(userID, imageName string) (containerID string, manifest string, err error) {
	//TODO: implements imageName Validation
	//TODO: update deployment manifest on store or create new deployment
	ns, _ := g.K8SProvider.Client().CoreV1().Namespaces().Get(userID, v1.GetOptions{})
	log.Printf("info: namespace is %v\n", ns.ObjectMeta.Name)
	if ns.ObjectMeta.Name == cloudutil.NullString {
		log.Printf("info: namespace is not found, next create namespace\n")
		_, err = g.K8SProvider.Client().CoreV1().Namespaces().Create(&core.Namespace{
			TypeMeta: v1.TypeMeta{},
			ObjectMeta: v1.ObjectMeta{
				Name: userID,
			},
			Spec:   core.NamespaceSpec{},
			Status: core.NamespaceStatus{},
		})
		if err != nil {
			return cloudutil.NullString, cloudutil.NullString, err
		}
	}
	// create deployment manifest for user
	containerID = cloudutil.NewUUID()
	object, _ := g.K8SProvider.Manifest().NewDeploymentObject()
	object.Name = containerID
	object.Namespace = userID
	object.Spec.Template.Spec.Containers = []core.Container{{Name: containerID, Image: imageName}}
	manifest, _ = g.K8SProvider.Manifest().TransformObjectToYaml(object)
	_, ok := g.K8SProvider.Kubectl().Apply(manifest)
	if !ok {
		return cloudutil.NullString, cloudutil.NullString, fmt.Errorf("")
	}
	return
}

func (g ContainerGateway) DeleteByContainerID(userID, containerID string) error {
	//TODO: user is exist
	deploymentManifest, err := g.GetDeploymentManifestByContainerID(containerID)
	if err != nil {
		return fmt.Errorf("call ContainerGateway.GetDeploymentManifestByContainerID: %w", err)
	}
	cloudutil.Logf("[debug] deployment manifest (%s) on ContainerGateway.DeleteByContainerID()\n", deploymentManifest)
	result, ok := g.K8SProvider.Kubectl().Delete(deploymentManifest)
	if !ok {
		cloudutil.Logf("info: kubectl apply result %v\n", result)
		return fmt.Errorf("call ContainerGateway.DeleteByContainerID: %w", err)
	}
	return nil
}

func (g ContainerGateway) GetDeploymentManifestByContainerID(containerID string) (manifest string, err error) {
	ctx, _ := context.WithTimeout(context.Background(), 5*time.Second)
	var doc bson.M
	if err = g.Mongo.DB().Collection(_collectionName).FindOne(ctx, bson.M{"container_id": containerID}).Decode(&doc); err != nil {
		return cloudutil.NullString, err
	}
	return doc["manifest"].(string), nil
}

func (g ContainerGateway) SaveDeploymentManifestByContainerID(containerID, manifest string) error {
	ctx, _ := context.WithTimeout(context.Background(), 5*time.Second)
	if _, err := g.Mongo.DB().Collection(_collectionName).InsertOne(ctx, bson.M{
		"container_id": containerID,
		"manifest":     manifest,
	}); err != nil {
		return err
	}
	return nil
}

func (g ContainerGateway) DeleteDeploymentManifestByContainerID(containerID string) error {
	ctx, _ := context.WithTimeout(context.Background(), 5*time.Second)
	if _, err := g.Mongo.DB().Collection(_collectionName).DeleteOne(ctx, bson.M{
		"container_id": containerID,
	}); err != nil {
		return err
	}
	return nil
}
