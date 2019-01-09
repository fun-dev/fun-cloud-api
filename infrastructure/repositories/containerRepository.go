package repositories

import (
	"github.com/fun-dev/cloud-api/domain/models"
	"github.com/fun-dev/cloud-api/infrastructure/repositories/interfaces"
	"github.com/go-xorm/xorm"
)

type containerRepository struct {
	Engine *xorm.Engine
}

func NewContainerRepository(engine *xorm.Engine) interfaces.IContainerRepository {
	return containerRepository{Engine: engine}
}

func (repo containerRepository) GetContainersByNamespace(namespace string) ([]models.Container, error) {
	// TODO: ここにk8sからコンテナの一覧を取得する処理を記述する
	return []models.Container{}, nil
}

func (repo containerRepository) CreateContainer(userToken, imageID string) (models.Container, error) {
	// TODO youtangai コンテナ作成の処理を記述
	return models.Container{}, nil
}

func (repo containerRepository) DeleteContainerByID(userToken string, containerID int64) error {
	// TODO youtangai コンテナ削除の処理を記述する
	return nil
}
