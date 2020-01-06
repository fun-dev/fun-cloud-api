package controller

import (
	"github.com/fun-dev/fun-cloud-api/pkg/jwt"
	"github.com/fun-dev/fun-cloud-api/pkg/term"
	"github.com/gin-gonic/gin"
	"net/http"
)

type (
	AuthController struct {
		Jwt  jwt.IJwt
	}

	IAuthController interface {
		TokenValidate(c *gin.Context)
	}
)

func NewAuthController() IAuthController {
	return &AuthController{}
}

func (ac AuthController) TokenValidate(c *gin.Context) {
	accessToken := c.GetHeader("Authorization")
	if accessToken == term.NullString {
		c.JSON(http.StatusBadRequest, gin.H{"message":"please set access token on the header"})
		return
	}
	_, err := ac.Jwt.InspectGoogleIdToken(accessToken)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message":"your access token is invalid"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message":"your access token is valid"})
	return
}