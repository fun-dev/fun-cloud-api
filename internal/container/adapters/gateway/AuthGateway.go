package gateway

import (
	"github.com/fun-dev/fun-cloud-api/pkg/auth"
)

type AuthGateway struct{}

func NewAuthGateway() auth.Repository {
	return &AuthGateway{}
}

func (a AuthGateway) GetUserIDByToken(token string) (userID string, err error) {
	// TODO: implement auth service token validate endpoint
	userID = token
	return
}
