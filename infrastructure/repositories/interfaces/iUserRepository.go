package interfaces

import "github.com/fun-dev/cloud-api/domain/models"

// IUserRepository is userRepository of interface
type IUserRepository interface {
	Insert(*models.User) error
	FindById(int64) (*models.User, error)
	FingByToken(string) (*models.User, error)
	Update(*models.User) error
	Delete(int64) error
}
