package controllers

import (
	"github.com/fun-dev/cloud-api/application/controllers/interfaces"
	"github.com/fun-dev/cloud-api/application/viewmodels"
	"github.com/fun-dev/cloud-api/domain"
	"github.com/fun-dev/cloud-api/domain/models"
	isrv "github.com/fun-dev/cloud-api/domain/services/interfaces"
	"github.com/fun-dev/cloud-api/middleware"
	"github.com/gin-gonic/gin"
	"net/http"
)

type UserController struct {
	Srv isrv.IUserService
}

func NewUserController() interfaces.IUserController {
	return UserController{
		Srv: domain.UserSrv,
	}
}

func (ctrl UserController) Get(c *gin.Context) {
	token := c.GetHeader("Authorization")
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
	if token != "" {
		claim, err := middleware.JWTValidate(token) // [Get] User Claim from JWT
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"err": err.Error()})
			return
		}
		// [Set] User Attribute
		user := models.User{}
		user.AccessToken = token
		user.Email = claim.Email
		user.GoogleName = claim.Name
		user.IconUrl = claim.Picture
		err = ctrl.Srv.Add(&user)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"err": err.Error()})
			return
		}
		c.Status(http.StatusCreated)
		return
	}
	c.JSON(http.StatusInternalServerError, gin.H{"err": "Can't get token"})
}

func (ctrl UserController) Update(c *gin.Context) {
	token := c.GetHeader("Authorization")
	if token != "" {
		user, err := ctrl.Srv.Get(token)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"err": err.Error()})
			return
		}
		// [Set] Update Access Token
		user.AccessToken = token
		err = ctrl.Srv.Update(user)
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
