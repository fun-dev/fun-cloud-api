package gateway

import (
	"github.com/fun-dev/ccms/adapters/gateway/repository"
)

type AuthGateway struct {}

func NewAuthGateway() repository.AuthRepository {
	return &AuthGateway{}
}

func (a AuthGateway) GetUserIDByToken(token string) (userID string, err error) {
	// TODO: implement auth service token validate endpoint
	userID = token
	return
}
