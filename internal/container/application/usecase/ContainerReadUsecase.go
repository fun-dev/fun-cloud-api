package usecase

import (
	"github.com/fun-dev/fun-cloud-api/internal/container/domain/models"
	"github.com/fun-dev/fun-cloud-api/internal/container/domain/repository"
	"github.com/fun-dev/fun-cloud-api/pkg/cloudauth"
	"github.com/fun-dev/fun-cloud-api/pkg/cloudutil"
)

type (
	ContainerReadUsecaseEntry struct {
		Containers []*models.Container
	}

	ContainerReadUsecaseResponse struct {
		Entry ContainerReadUsecaseEntry
	}

	ContainerReadUsecase interface {
		Execute(userID, imageName string) (resp ContainerReadUsecaseResponse, err error)
	}

	ContainerReadInteractor struct {
		cRepo repository.Repository
		aRepo cloudauth.Repository
	}
)

func NewContainerReadInteractor(cRepo repository.Repository, aRepo cloudauth.Repository) ContainerReadUsecase {
	return &ContainerReadInteractor{cRepo: cRepo, aRepo: aRepo}
}

/*
Execute
@Option: named function
@param: userID
@param imageName
*/
func (c ContainerReadInteractor) Execute(userID, imageName string) (resp ContainerReadUsecaseResponse, err error) {
	// in this application, we use userID as cloudk8s namespace.yaml
	//TODO: Get All by UserID and Get Single by ImageName
	switch imageName {
	case cloudutil.NullString:
		resp.Entry.Containers, err = c.cRepo.GetAllByUserID(userID)
		if err != nil {
			return
		}
		return
	default:
		return
	}
}
