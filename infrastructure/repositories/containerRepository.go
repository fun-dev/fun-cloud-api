package repositories

import (
	"github.com/fun-dev/cloud-api/infrastructure/dbmodels"
	"github.com/fun-dev/cloud-api/infrastructure/repositories/interfaces"
	"github.com/go-xorm/xorm"
)

type containerRepository struct {
	Engine *xorm.Engine
}

func NewContainerRepository(engine *xorm.Engine) interfaces.IContainerRepository {
	return containerRepository{Engine: engine}
}

func (repo containerRepository) GetContainersByNamespace(namespace string) ([]dbmodels.Container, error) {
	// TODO: ここにk8sからコンテナの一覧を取得する処理を記述する
	return []dbmodels.Container{}, nil
}
