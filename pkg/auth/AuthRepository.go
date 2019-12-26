package auth

// AuthRepository is
type Repository interface {
	GetUserIDByToken(token string) (userID string, err error)
}
