package interfaces

import (
	"github.com/fun-dev/cloud-api/domain/models"
)

type IContainerService interface {
	GetContainersByToken(userToken string) ([]models.Container, error)
	PostContainerByID(userID int, containerID int) []models.Container
	PostContainerByToken(userToken, imageID string) (models.Container, error)
	DeleteContainerByID(userToken string, containerID int64) error
}
