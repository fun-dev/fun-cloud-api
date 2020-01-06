package controller

import (
	"net/http"

	"github.com/fun-dev/fun-cloud-api/internal/auth/model"
	"github.com/fun-dev/fun-cloud-api/pkg/jwt"
	"github.com/fun-dev/fun-cloud-api/pkg/term"
	"github.com/gin-gonic/gin"
)

type (
	UserController struct {
		// for connecting user store
		User model.IUser
		// validate user access token
		Jwt jwt.IJwt
	}

	IUserController interface {
		GET(c *gin.Context)
		POST(c *gin.Context)
	}
)

func NewUserController(user model.IUser) IUserController {
	return &UserController{User: user}
}

func (u UserController) GET(c *gin.Context) {
	accessToken := c.GetHeader("Authorization")
	if accessToken == "" {
		c.JSON(http.StatusBadRequest, gin.H{})
		return
	}
	user, err := u.User.GetByAccessToken(accessToken)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{})
		return
	}
	if user != nil {
		c.JSON(http.StatusNoContent, gin.H{})
		return
	}
	c.JSON(http.StatusOK, user)
	return
}

func (u UserController) POST(c *gin.Context) {
	accessToken := c.GetHeader("Authorization")
	if accessToken == term.NullString {
		//TODO: implement error handling
	}
	claim, err := u.Jwt.InspectGoogleIdToken(accessToken)
	if err != nil {
		//TODO: implement error handling
		return
	}
	input := model.NewUser(claim.Picture, claim.Name, accessToken)
	if err := u.User.Create(*input); err != nil {
		//TODO: implement error handling
		return
	}
	c.String(http.StatusCreated, "")
}
