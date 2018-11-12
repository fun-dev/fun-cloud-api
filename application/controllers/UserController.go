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

// func (ctrl UserController) Update(c *gin.Context) {
// 	id, err := getId(c)
// 	if err != nil {
// 		c.JSON(http.StatusBadRequest, gin.H{"get id err": err.Error()})
// 		return
// 	}

// 	var json viewmodels.User
// 	if err := c.ShouldBindJSON(&json); err != nil {
// 		c.JSON(http.StatusBadRequest, gin.H{"bind json err": err.Error()})
// 		return
// 	}

// 	model := viewModelToDomainModel(&json)
// 	model.Id = id

// 	if err := ctrl.Srv.Update(model); err != nil {
// 		c.JSON(http.StatusInternalServerError, gin.H{"db update err": err.Error()})
// 		return
// 	}

// 	c.Status(http.StatusOK)
// }

// func (ctrl UserController) Delete(c *gin.Context) {
// 	id, err := getId(c)
// 	if err != nil {
// 		c.JSON(http.StatusBadRequest, gin.H{"err": err.Error()})
// 		return
// 	}
// 	if err := ctrl.Srv.Delete(id); err != nil {
// 		c.JSON(http.StatusInternalServerError, gin.H{"err": err.Error()})
// 		return
// 	}
// 	c.Status(http.StatusOK)
// }

// func getId(c *gin.Context) (int64, error) {
// 	id, err := strconv.Atoi(c.Param("id"))
// 	if err != nil {
// 		return -1, err
// 	}
// 	return int64(id), nil
// }

func getToken(c *gin.Context) string {
	token := c.GetHeader("Authorization")
	return token
}

func domainModelToViewModel(user *models.User) *viewmodels.User {
	return &viewmodels.User{
		IconURL:     user.IconURL,
		GoogleName:  user.GoogleName,
		AccessToken: user.AccessToken,
	}
}

func viewModelToDomainModel(user *viewmodels.User) *models.User {
	return &models.User{
		IconURL:     user.IconURL,
		GoogleName:  user.GoogleName,
		AccessToken: user.AccessToken,
	}
}
