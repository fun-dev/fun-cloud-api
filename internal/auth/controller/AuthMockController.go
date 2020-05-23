package controller

import (
	"github.com/fun-dev/fun-cloud-api/pkg/cloudutil"
	"github.com/gin-gonic/gin"
	"net/http"
)

type AuthMockController struct{}

func NewAuthMockController() IAuthController {
	return &AuthMockController{}
}

func (am AuthMockController) TokenValidate(c *gin.Context) {
	accessToken := c.GetHeader("Authorization")
	if accessToken == cloudutil.NullString {
		c.JSON(http.StatusBadRequest, gin.H{"message": "please set access token on the header"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "your access token is valid"})
	return
}
