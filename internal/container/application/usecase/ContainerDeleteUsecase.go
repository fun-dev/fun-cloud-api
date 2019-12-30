package usecase

import (
	"fmt"
	"github.com/fun-dev/fun-cloud-api/internal/container/domain/container"
	"github.com/fun-dev/fun-cloud-api/pkg/auth"
)

type (
	// ContainerDeleteUsecase is Usecase
	ContainerDeleteUsecase interface {
		Execute(userID, containerID string) error
	}
	// ContainerDeleteInteractor is Interactor
	ContainerDeleteInteractor struct {
		cRepo container.Repository
		aRepo auth.Repository
	}
)

// NewContainerDeleteInteractor is ...
func NewContainerDeleteInteractor(cRepo container.Repository, aRepo auth.Repository) ContainerDeleteUsecase {
	return &ContainerDeleteInteractor{cRepo, aRepo}
}

/* Execute is executing container delete usecase
@param containerID: unique string like uuid
*/
func (i ContainerDeleteInteractor) Execute(userID, containerID string) error {
	// in this application, we use userID as kubernetes namespace.yaml
	if err := i.cRepo.DeleteByContainerID(userID, containerID); err != nil {
		return fmt.Errorf("call ContainerRepo.DeleteByContainerID: %w", err)
	}
	if err := i.cRepo.DeleteDeploymentManifestByContainerID(containerID); err != nil {
		return fmt.Errorf("call DeleteDeploymentManifestByContainerID: %w", err)
	}
	return nil
}
