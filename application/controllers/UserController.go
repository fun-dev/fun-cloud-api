package controllers

import (
	"github.com/fun-dev/cloud-api/application/controllers/interfaces"
	"github.com/fun-dev/cloud-api/application/viewmodels"
	"github.com/fun-dev/cloud-api/domain"
	"github.com/fun-dev/cloud-api/domain/models"
	isrv "github.com/fun-dev/cloud-api/domain/services/interfaces"
	"github.com/gin-gonic/gin"
	"net/http"
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
	token := ""
	model, err := ctrl.Srv.Get(token)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"err": err.Error()})
		return
	}
	viewModel := domainModelToViewModel(model)
	c.JSON(http.StatusOK, viewModel)
}

func (ctrl UserController) Create(c *gin.Context) {
	token := c.GetHeader("Authorization")
	user := models.User{}
	user.IconUrl = ""
	user.GoogleName = ""
	user.AccessToken = token
	if token != "" {
		err := ctrl.Srv.Add(&user)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"err": err.Error()})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"err": "Can't get token"})
		return
	}
	c.Status(http.StatusOK)
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
