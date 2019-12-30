package usecase

import (
	"github.com/fun-dev/fun-cloud-api/internal/container/domain/container"
	"github.com/fun-dev/fun-cloud-api/pkg/auth"
	"github.com/fun-dev/fun-cloud-api/pkg/term"
)

type (
	ContainerReadUsecaseEntry struct {
		Containers []*container.Container
	}

	ContainerReadUsecaseResponse struct {
		Entry ContainerReadUsecaseEntry
	}

	ContainerReadUsecase interface {
		Execute(userID, imageName string) (resp ContainerReadUsecaseResponse, err error)
	}

	ContainerReadInteractor struct {
		cRepo container.Repository
		aRepo auth.Repository
	}
)

func NewContainerReadInteractor(cRepo container.Repository, aRepo auth.Repository) ContainerReadUsecase {
	return &ContainerReadInteractor{cRepo:cRepo,aRepo:aRepo}
}

/*
Execute
@Option: named function
@param: userID
@param imageName
*/
func (c ContainerReadInteractor) Execute(userID, imageName string) (resp ContainerReadUsecaseResponse, err error) {
	// in this application, we use userID as kubernetes namespace.yaml
	//TODO: Get All by UserID and Get Single by ImageName
	switch imageName {
	case term.NullString:
		resp.Entry.Containers, err = c.cRepo.GetAllByUserID(userID)
		if err != nil {
			return
		}
		return
	default:
		return
	}
}
