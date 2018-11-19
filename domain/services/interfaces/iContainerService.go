package interfaces

import (
	"github.com/fun-dev/cloud-api/domain/models"
)

type IContainerService interface {
	GetContainersByUserID(UserID int) ([]models.Container, error)
}
