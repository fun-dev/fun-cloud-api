package auth

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type (
	IAuthService interface {
		EnsureIsExist() gin.HandlerFunc
	}

	Service struct {
		AuthRepository Repository
	}
)

func NewAuthService(authRepo Repository) IAuthService {
	return &Service{authRepo}
}

func (s Service) EnsureIsExist() gin.HandlerFunc {
	return func(c *gin.Context) {
		key := "Authorization"
		accessToken := c.GetHeader(key)
		userID, err := s.AuthRepository.GetUserIDByToken(accessToken)
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"message": "your access token is not found"})
			c.Abort()
			return
		}
		key = "USER_ID"
		c.Set(key, userID)
		c.Next()
	}
}
