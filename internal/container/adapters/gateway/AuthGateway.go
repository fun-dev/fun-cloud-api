package gateway

import (
	"github.com/fun-dev/fun-cloud-api/internal/container/domain/auth"
)

type AuthGateway struct{}

func NewAuthGateway() auth.AuthRepository {
	return &AuthGateway{}
}

func (a AuthGateway) GetUserIDByToken(token string) (userID string, err error) {
	// TODO: implement auth service token validate endpoint
	userID = token
	return
}
