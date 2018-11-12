package services

import (
	"github.com/fun-dev/cloud-api/domain/models"
	isrv "github.com/fun-dev/cloud-api/domain/services/interfaces"
	"github.com/fun-dev/cloud-api/infrastructure"
	irepo "github.com/fun-dev/cloud-api/infrastructure/repositories/interfaces"
)

type UserService struct {
	Repo irepo.IUserRepository
}

func NewUserService() isrv.IUserService {
	return UserService{
		Repo: infrastructure.UserRepo,
	}
}

// Get : return User model from token
func (srv UserService) Get(token string) (*models.User, error) {
	user, err := srv.Repo.FingByToken(token)
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (srv UserService) Add(user *models.User) error {
	err := srv.Repo.Insert(user)
	if err != nil {
		return err
	}
	return nil
}

func (srv UserService) Update(user *models.User) error {
	err := srv.Repo.Update(user)
	if err != nil {
		return err
	}
	return nil
}

func (srv UserService) Delete(id int64) error {
	err := srv.Repo.Delete(id)
	if err != nil {
		return err
	}
	return nil
}
