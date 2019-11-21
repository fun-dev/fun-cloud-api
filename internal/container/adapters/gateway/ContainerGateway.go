package gateway

import (
	"context"
	"errors"
	"fmt"
	"github.com/fun-dev/fun-cloud-api/internal/container/domain/container"
	"github.com/fun-dev/fun-cloud-api/internal/container/infrastructure/apperror/repoerr"
	"github.com/fun-dev/fun-cloud-api/pkg/kubernetes"
	"github.com/fun-dev/fun-cloud-api/pkg/kubernetes/kubectl"
	"github.com/fun-dev/fun-cloud-api/pkg/logging"
	"github.com/fun-dev/fun-cloud-api/pkg/redis"
	"github.com/fun-dev/fun-cloud-api/pkg/term"
	apps "k8s.io/api/apps/v1"
	core "k8s.io/api/core/v1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// ContainerGateway is
type ContainerGateway struct {
	K8SProvider kubernetes.IK8SProvider
	Redis redis.Driver
}
// NewContainerGateway is
func NewContainerGateway(k kubernetes.IK8SProvider, r redis.Driver) container.Repository {
	return &ContainerGateway{
		k,
		r,
	}
}

// DeleteByContainerID is
func (g ContainerGateway) DeleteByContainerID(ctx context.Context, id string, namespace string) error {
	deploymentManifest, err := g.GetDeploymentManifestByContainerID(ctx, id)
	if err != nil {
		return fmt.Errorf("call ContainerGateway.GetDeploymentManifestByContainerID: %w", err)
	}
	logging.Logf("[debug] deployment manifest (%s) on ContainerGateway.DeleteByContainerID()\n", deploymentManifest)
	return g.K8SProvider.Kubectl().Execute(kubectl.Delete, deploymentManifest, namespace)
}

func (g ContainerGateway) GetDeploymentManifestByContainerID(ctx context.Context, id string) (manifest string, err error) {
	key := "deployment_" + id
	// TODO: adapt getter
	manifest = g.Redis.Client.Get(key).String()
	if manifest == term.NullString {
		return "", repoerr.DeploymentManifestCanNotBeFound
	}
	logging.Logf("info: ContainerGateway.GetDeploymentManifestByContainerID result is ", manifest)
	return
}

func (g ContainerGateway) GetAllByUserID(ctx context.Context, id, namespace string) ([]*container.Container, error) {
	deploymentList, err := g.K8SProvider.Client().AppsV1().Deployments(namespace).List(v1.ListOptions{})
	if err != nil {
		return nil, err
	}
	if len(deploymentList.Items) == term.HasNoItem {
		return nil, errors.New("error: you don't have any containers")
	}
	resultContainers := make([]*container.Container, len(deploymentList.Items))
	for i, d := range deploymentList.Items {
		podContainer := d.Spec.Template.Spec.Containers[0]
		resultContainer := &container.Container{
			UID: podContainer.Name,
			ImageName: podContainer.Image,
			ConnectInfo: d.SelfLink, // TODO: fix deprecated
			Status: d.Status.String(),
		}
		resultContainers[i] = resultContainer
	}
	return resultContainers, nil
}

func (g ContainerGateway) Create(ctx context.Context, id, imageName, namespace string) (manifest string, err error) {
	object, _ := g.K8SProvider.Kubectl().DeserializeYamlToObject(kubectl.UseDeploymentManifest, &apps.Deployment{})
	deployment := object.(*apps.Deployment)
	podContainer := core.Container{
		Name: id,
		Image: imageName,
	}
	deployment.Spec.Template.Spec.Containers[0] = podContainer
	manifest, _ = g.K8SProvider.Kubectl().DecodeObjectToYaml(deployment)
	err = g.K8SProvider.Kubectl().Execute(kubectl.Apply, manifest, namespace)
	if err != nil {
		return
	}
	return
}

func (g ContainerGateway) SaveDeploymentManifestByContainerID(ctx context.Context, userID, id, yaml string) error {
	key := userID + "_" + "id"
	err := g.Redis.Client.Append(key, yaml).Err()
	if err != nil {
		return err
	}
	return nil
}