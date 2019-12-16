package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type (
	UserController struct {
	Name           string `db:"name",json:"name"`
	GoogleUserName string `db:"google_name",json:"google_name"`
	AccessToken    string `db:"access_token",json:"access_token"`
	IconURL        string `db:"icon_url",json:"icon_url"`

	User IUser
	}

	IUserController interface {
		GET(c *gin.Context)
		POST(c *gin.Context)
	}
)

func NewUserController(user IUser) IUserController {
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

}
