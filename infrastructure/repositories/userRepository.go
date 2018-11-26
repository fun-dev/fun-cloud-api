package repositories

import (
	"fmt"

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

func (repo userRepository) Insert(user *models.User) error {
	session := repo.Engine.NewSession()
	defer session.Close()
	if err := session.Begin(); err != nil {
		return err
	}

	dbmodel := domainModelToDBmodel(user)
	_, err := session.Insert(dbmodel)
	if err != nil {
		session.Rollback()
		return err
	}

	if err := session.Commit(); err != nil {
		return err
	}
	return nil
}

func (repo userRepository) FindById(id int64) (*models.User, error) {
	var user dbmodels.User
	isExist, err := repo.Engine.Id(id).Get(&user)
	if err != nil {
		return nil, err
	}
	if !isExist {
		return nil, fmt.Errorf("no such user in databse")
	}
	model := dbmodelToDomainModel(&user)
	return model, nil
}

func (repo userRepository) FindByToken(token string) (*models.User, error) {
	var user dbmodels.User
	isExist, err := repo.Engine.Where("token = ?", token).Get(&user)
	if err != nil {
		return nil, err
	}
	if !isExist {
		return nil, fmt.Errorf("no such user in databse")
	}
	model := dbmodelToDomainModel(&user)
	return model, nil
}

func (repo userRepository) Update(user *models.User) error {
	session := repo.Engine.NewSession()
	defer session.Close()
	if err := session.Begin(); err != nil {
		return err
	}

	dbmodel := domainModelToDBmodel(user)
	if _, err := session.Id(dbmodel.Id).Update(dbmodel); err != nil {
		session.Rollback()
		return err
	}

	if err := session.Commit(); err != nil {
		return err
	}
	return nil
}

func (repo userRepository) Delete(id int64) error {
	session := repo.Engine.NewSession()
	defer session.Close()
	if err := session.Begin(); err != nil {
		return err
	}

	if _, err := session.Id(id).Delete(new(dbmodels.User)); err != nil {
		session.Rollback()
		return err
	}

	if err := session.Commit(); err != nil {
		return err
	}
	return nil
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
