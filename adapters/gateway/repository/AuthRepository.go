package repository

import (
	"github.com/fun-dev/ccms-poc/domain/value"
)

// AuthRepository is
type AuthRepository interface {
	GetUserIDByToken(token string) (*value.UserID, error)
}
