package controller

import (
	"github.com/fun-dev/fun-cloud-api/pkg/cloudauth"
	"github.com/fun-dev/fun-cloud-api/pkg/cloudutil"
	"net/http"

	"github.com/fun-dev/fun-cloud-api/internal/auth/model"
	"github.com/gin-gonic/gin"
)

type (
	UserController struct {
		User model.IUser
		Jwt  cloudauth.IJwt
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
		c.JSON(http.StatusBadRequest, gin.H{"message": "accessToken is nil"})
		return
	}
	user, err := u.User.GetByAccessToken(accessToken)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Not match correct accessToken"})
		return
	}
	if user != nil {
		c.JSON(http.StatusNoContent, gin.H{"message": "There are no users"})
		return
	}
	c.JSON(http.StatusOK, user)
	return
}

func (u UserController) POST(c *gin.Context) {
	accessToken := c.GetHeader("Authorization")
	if accessToken == cloudutil.NullString {
		// ex error message: please set value on authorization header
		c.JSON(http.StatusBadRequest, gin.H{"message": "Please set value on authorization header"})
		return
	}
	claim, err := u.Jwt.InspectGoogleIdToken(accessToken)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Your acccess token is invalid"})
		return
	}
	input := model.NewUser(claim.Picture, claim.Name, accessToken)
	if err := u.User.Create(*input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Failed to create new user"})
		return
	}
	c.String(http.StatusCreated, "")
	return
}
