package interfaces

import (
	"github.com/fun-dev/cloud-api/domain/models"
)

type IUserService interface {
	Get(int64) (*models.User, error)
	Add(*models.User) error
	Update(*models.User) error
	Delete(int64) error
}