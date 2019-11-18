package gateway

import (
	"github.com/fun-dev/ccms-poc/adapters/gateway/repository"
	"github.com/fun-dev/ccms-poc/domain/value"
)

type AuthGateway struct {}

func NewAuthGateway() repository.AuthRepository {
	return &AuthGateway{}
}

func (a AuthGateway) GetUserIDByToken(token string) (*value.UserID, error) {
	// TODO: implement get user id from google access token
	panic("implement me")
}

