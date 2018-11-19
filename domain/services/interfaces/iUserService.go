package interfaces

import (
	"github.com/fun-dev/cloud-api/domain/models"
)

type IUserService interface {
	Get(string) (*models.User, error)
	Add(*models.User) error
}
