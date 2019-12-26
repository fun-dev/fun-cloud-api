package usecase

import (
	"context"
	"fmt"
	"github.com/fun-dev/fun-cloud-api/pkg/auth"
	"github.com/fun-dev/fun-cloud-api/internal/container/domain/container"
)

type (
	// ContainerDeleteUsecase is Usecase
	ContainerDeleteUsecase interface {
		Execute(ctx context.Context, userID, containerID string) error
	}
	// ContainerDeleteInteractor is Interactor
	ContainerDeleteInteractor struct {
		cRepo container.Repository
		aRepo auth.AuthRepository
	}
)

// NewContainerDeleteInteractor is ...
func NewContainerDeleteInteractor(cRepo container.Repository, aRepo auth.AuthRepository) ContainerDeleteUsecase {
	return &ContainerDeleteInteractor{cRepo, aRepo}
}

/* Execute is executing container delete usecase
@param containerID: unique string like uuid
*/
func (i ContainerDeleteInteractor) Execute(ctx context.Context, userID, containerID string) error {
	// in this application, we use userID as kubernetes namespace.yaml
	namespace := userID
	if err := i.cRepo.DeleteByContainerID(ctx, containerID, namespace); err != nil {
		return fmt.Errorf("call ContainerRepo.DeleteByContainerID: %w", err)
	}
	return nil
}
