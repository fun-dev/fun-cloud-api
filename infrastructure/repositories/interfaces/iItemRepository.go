package interfaces

import "github.com/fun-dev/cloud-api/domain/models"

type IItemRepository interface {
	Insert(*models.Item) error
	FindById(int64) (*models.Item, error)
	FindByUserId(int64) (*models.Item, error)
	Update(*models.Item) error
	Delete(*models.Item) error
}
