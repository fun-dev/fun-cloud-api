package services

import (
	"github.com/fun-dev/cloud-api/domain/models"
	"github.com/stretchr/testify/assert"
	"testing"
)

type userRepositoryTest struct {
}

func (repo userRepositoryTest) FindByToken(token string) (*models.User, error) {
	model := &models.User{
		IconUrl:     "TEST1",
		GoogleName:  "TEST1",
		AccessToken: "TEST1",
	}
	return model, nil
}

func (repo userRepositoryTest) Insert(user *models.User) error {
	return nil
}

func (repo userRepositoryTest) Update(user *models.User) error {
	return nil
}

func (repo userRepositoryTest) Delete(id int64) error {
	return nil
}

// Get : return User model from token
func TestGet(t *testing.T) {
	assert := assert.New(t)
	repo := userRepositoryTest{}
	srv := UserService{
		Repo: repo,
	}
	token := "TEST1"
	actual, _ := UserService.Get(srv, token)
	expected := &models.User{
		IconUrl:     "TEST1",
		GoogleName:  "TEST1",
		AccessToken: "TEST1",
	}
	// assert equality
	assert.Equal(expected, actual, "they should be equal")
}
