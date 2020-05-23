package gateway

import (
	"github.com/fun-dev/fun-cloud-api/internal/container/domain/repository"
	"github.com/fun-dev/fun-cloud-api/pkg/cloudk8s"
	"github.com/fun-dev/fun-cloud-api/pkg/cloudutil"
	"github.com/fun-dev/fun-cloud-protobuf/container/rpc"
	v1 "k8s.io/api/core/v1"
	"os"
)

// ContainerGateway is
type ContainerGateway struct {}

func (c ContainerGateway) GetAllByUserID(userID, namespace string) ([]*rpc.Container, error) {
	pods, err := cloudk8s.GetPodsOnKubeAPIClient(userID)
	if err != nil {
		return nil, err
	}
	var result []*rpc.Container
	for _, pod  := range pods.Items {
		container := rpc.Container{
			Id:        pod.Name,
			ImageName: pod.Spec.Containers[0].Image,
			Url:       pod.GetSelfLink(),
			Condition: _judgePodCondition(pod.Status.Conditions[0].Status),
		}
		result = append(result, &container)
	}
	return result, nil
}

func (c ContainerGateway) Create(userID, imageName, namespace string) error {
	err := cloudk8s.IsExistNamespaceOnKubeAPIClient(namespace)
	if err != nil {
		if err == cloudk8s.ErrNamespaceIsNotFound {
			if err := cloudk8s.CreateNamespaceOnKubeAPIClient(namespace); err != nil {
				return err
			}
		}
		return err
	}
	//	// create deployment manifest for user
	appName := cloudutil.NewUUID()
	// PVC
	pvc, _ := cloudk8s.GetTemplatePVCObject()
	pvc.Name = _genPVCName(appName)
	pvc.Namespace = namespace
	pvcManifest, _ := cloudk8s.GetManifestFromObject(pvc)
	pvcManifestPath, _ := cloudutil.CreateTmpManifest(pvcManifest)
	defer os.Remove(pvcManifestPath)
	if err := cloudk8s.ExecuteManifestOnKubectl(pvcManifestPath, cloudk8s.APPLY); err != nil {
		return err
	}
	if err := cloudk8s.SaveManifestOnKubeAPIClient(pvc.Name, pvcManifest, namespace); err != nil {
		return err
	}
	// Deployment
	deployment, _ := cloudk8s.GetTemplateDeploymentObject()
	deployment.Name = _genDeploymentName(appName)
	deployment.Namespace = namespace
	deployment.Spec.Template.Spec.Containers = []v1.Container{{
		Name:                     imageName,
		Image:                    imageName,
		Command:                  nil,
		Args:                     nil,
		WorkingDir:               "",
		Ports:                    nil,
		EnvFrom:                  nil,
		Env:                      nil,
		Resources:                v1.ResourceRequirements{},
		VolumeMounts:             []v1.VolumeMount{},
		VolumeDevices:            nil,
		LivenessProbe:            nil,
		ReadinessProbe:           nil,
		StartupProbe:             nil,
		Lifecycle:                nil,
		TerminationMessagePath:   "",
		TerminationMessagePolicy: "",
		ImagePullPolicy:          "",
		SecurityContext:          nil,
		Stdin:                    false,
		StdinOnce:                false,
		TTY:                      false,
	}}
	deployment.Spec.Template.Spec.Volumes = []v1.Volume{{
		Name:         appName+"-"+"vol",
		VolumeSource: v1.VolumeSource{
			PersistentVolumeClaim: &v1.PersistentVolumeClaimVolumeSource{
				ClaimName: appName+"-"+"pvc",
				ReadOnly:  false,
			},
		},
	}}
	deploymentManifest, _ := cloudk8s.GetManifestFromObject(deployment)
	deploymentManifestPath, _ := cloudutil.CreateTmpManifest(deploymentManifest)
	defer os.Remove(deploymentManifestPath)
	if err := cloudk8s.ExecuteManifestOnKubectl(deploymentManifest, cloudk8s.APPLY); err != nil {
		return err
	}
	if err := cloudk8s.SaveManifestOnKubeAPIClient(deployment.Name, deploymentManifest, namespace); err != nil {
		return err
	}
	return nil
}

func (c ContainerGateway) DeleteByContainerID(userID, appName, namespace string) error {
	pvcName := _genPVCName(appName)
	deploymentName := _genDeploymentName(appName)
	pvcManifest, _ := cloudk8s.GetManifestOnKubeAPIClient(pvcName, namespace)
	pvcManifestPath, _ := cloudutil.CreateTmpManifest(pvcManifest)
	defer os.Remove(pvcManifestPath)
	deploymentManifest, _ := cloudk8s.GetManifestOnKubeAPIClient(deploymentName, namespace)
	deploymentManifestPath, _ := cloudutil.CreateTmpManifest(deploymentManifest)
	defer os.Remove(deploymentManifestPath)
	if err := cloudk8s.ExecuteManifestOnKubectl(pvcManifestPath, cloudk8s.DELETE); err != nil {
		return err
	}
	if err := cloudk8s.ExecuteManifestOnKubectl(deploymentManifestPath, cloudk8s.DELETE); err != nil {
		return err
	}
	if err := cloudk8s.DeleteManifestOnKubeAPIClient(pvcName, namespace); err != nil {
		return err
	}
	if err := cloudk8s.DeleteManifestOnKubeAPIClient(deploymentName, namespace); err != nil {
		return err
	}
	return nil
}

func NewContainerGateway() repository.ContainerRepository {
	return &ContainerGateway{}
}


func _judgePodCondition(condition v1.ConditionStatus) rpc.ContainerCondition {
	switch condition {
	case v1.ConditionTrue:
		return rpc.ContainerCondition_TRUE
	case v1.ConditionFalse:
		return rpc.ContainerCondition_FALSE
	case v1.ConditionUnknown:
		return rpc.ContainerCondition_UNKNOWN
	}
	return 3
}

func _genDeploymentName(name string) string {
	return name+"-"+"dep"
}

func _genPVCName(name string) string {
	return name+"-"+"pvc"
}