package controllers

import (
	"net/http"

	"github.com/fun-dev/cloud-api/domain/models"

	"github.com/fun-dev/cloud-api/application/controllers/interfaces"
	"github.com/fun-dev/cloud-api/application/viewmodels"
	"github.com/fun-dev/cloud-api/domain"
	isrv "github.com/fun-dev/cloud-api/domain/services/interfaces"
	"github.com/gin-gonic/gin"
)

type UserController struct {
	Srv isrv.IUserService
}

func NewUserController() interfaces.IUserController {
	return UserController{
		Srv: domain.UserService,
	}
}

func (ctrl UserController) Get(c *gin.Context) {
	token := getToken(c)
	if token == "" {
		c.JSON(http.StatusBadRequest, gin.H{"err": "Can't get token"})
		return
	}

	model, err := ctrl.Srv.Get(token)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"err": err.Error()})
		return
	}
	viewModel := domainModelToViewModel(model)
	c.JSON(http.StatusOK, viewModel)
}

func (ctrl UserController) Create(c *gin.Context) {
	var json viewmodels.User
	if err := c.ShouldBindJSON(&json); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"err": err.Error()})
		return
	}

	model := viewModelToDomainModel(&json)

	err := ctrl.Srv.Add(model)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"err": err.Error()})
		return
	}

	c.Status(http.StatusOK)
}

func getToken(c *gin.Context) string {
	token := c.GetHeader("Authorization")
	return token
}

func domainModelToViewModel(user *models.User) *viewmodels.User {
	return &viewmodels.User{
		IconUrl:     user.IconUrl,
		GoogleName:  user.GoogleName,
		AccessToken: user.AccessToken,
	}
}

func viewModelToDomainModel(user *viewmodels.User) *models.User {
	return &models.User{
		IconUrl:     user.IconUrl,
		GoogleName:  user.GoogleName,
		AccessToken: user.AccessToken,
	}
}
