package auth

import (
	"github.com/fun-dev/ccms/internal/container/adapters/gateway/repository"
	"github.com/gin-gonic/gin"
	"net/http"
)

type (
	IAuthService interface {
		EnsureIsExist() gin.HandlerFunc
	}

	AuthService struct {
		repository.AuthRepository
	}
)

func NewAuthService(aRepo repository.AuthRepository) IAuthService {
	return &AuthService{aRepo}
}

func (a AuthService) EnsureIsExist() gin.HandlerFunc {
	return func(c *gin.Context) {
		key := "Authorization"
		accessToken := c.GetHeader(key)
		userID, err := a.AuthRepository.GetUserIDByToken(accessToken)
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
