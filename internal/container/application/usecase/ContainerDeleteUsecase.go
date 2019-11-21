package usecase

import (
	"context"
	"fmt"
	"github.com/fun-dev/ccms/internal/container/adapters/gateway/repository"
)

type (
	// ContainerDeleteUsecase is Usecase
	ContainerDeleteUsecase interface {
		Execute(ctx context.Context, userID, containerID string) error
	}
	// ContainerDeleteInteractor is Interactor
	ContainerDeleteInteractor struct {
		ContainerRepo repository.ContainerRepository
		AuthRepo      repository.AuthRepository
	}
)

// NewContainerDeleteInteractor is ...
func NewContainerDeleteInteractor(containerRepo repository.ContainerRepository, authRepo repository.AuthRepository) ContainerDeleteUsecase {
	return &ContainerDeleteInteractor{ContainerRepo: containerRepo, AuthRepo: authRepo}
}

/* Execute is executing container delete usecase
@param containerID: unique string like uuid
*/
func (i ContainerDeleteInteractor) Execute(ctx context.Context, userID, containerID string) error {
	// in this application, we use userID as kubernetes namespace.yaml
	namespace := userID
	if err := i.ContainerRepo.DeleteByContainerID(ctx, containerID, namespace); err != nil {
		return fmt.Errorf("call ContainerRepo.DeleteByContainerID: %w", err)
	}
	return nil
}
