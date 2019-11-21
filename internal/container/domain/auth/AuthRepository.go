package auth

// AuthRepository is
type AuthRepository interface {
	GetUserIDByToken(token string) (userID string, err error)
}
