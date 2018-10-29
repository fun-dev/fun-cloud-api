package interfaces

import "github.com/fun-dev/cloud-api/domain/models"

type IUserRepository interface {
	Insert(*models.User) error
	FindById(int64) (*models.User, error)
	Update(*models.User) error
	Delete(int64) error
}
