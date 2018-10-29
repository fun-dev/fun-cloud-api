package domain

import (
	"github.com/fun-dev/cloud-api/domain/services"
	"github.com/fun-dev/cloud-api/domain/services/interfaces"
)

var (
	UserService interfaces.IUserService
)

func init() {
	UserService = services.NewUserService()
}
