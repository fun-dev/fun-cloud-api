package repositories

import (
	"github.com/fun-dev/cloud-api/domain/models"
	"github.com/fun-dev/cloud-api/infrastructure/dbmodels"
	"github.com/fun-dev/cloud-api/infrastructure/repositories/interfaces"
	"github.com/go-xorm/xorm"
)

type userRepository struct {
	Engine *xorm.Engine
}

// NewUserRepository is reuturn UserRepository
func NewUserRepository(engine *xorm.Engine) interfaces.IUserRepository {
	return userRepository{Engine: engine}
}

func (repo userRepository) Insert(user *models.User) (err error) {
	session := repo.Engine.NewSession()
	defer session.Close()
	if err = session.Begin(); err != nil {
		return
	}
	dbmodel := domainModelToDBmodel(user)
	_, err = session.Insert(dbmodel)
	if err != nil {
		session.Rollback()
		return
	}
	if err = session.Commit(); err != nil {
		return
	}
	return
}

func (repo userRepository) FindByToken(token string) (*models.User, error) {
	var user dbmodels.User
	_, err := repo.Engine.Where("access_token = ?", token).Get(&user)
	if err != nil {
		return nil, err
	}
	model := dbmodelToDomainModel(&user)
	return model, nil
}

func domainModelToDBmodel(user *models.User) *dbmodels.User {
	return &dbmodels.User{
		IconUrl:     user.IconUrl,
		GoogleName:  user.GoogleName,
		AccessToken: user.AccessToken,
	}
}

func dbmodelToDomainModel(user *dbmodels.User) *models.User {
	return &models.User{
		IconUrl:     user.IconUrl,
		GoogleName:  user.GoogleName,
		AccessToken: user.AccessToken,
	}
}
