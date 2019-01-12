package controllers

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/fun-dev/cloud-api/middleware"

	"github.com/fun-dev/cloud-api/application/controllers/interfaces"
	"github.com/fun-dev/cloud-api/application/viewmodels"
	"github.com/fun-dev/cloud-api/domain"
	isrv "github.com/fun-dev/cloud-api/domain/services/interfaces"
	"github.com/gin-gonic/gin"
)

type ContainerController struct {
	Srv isrv.IContainerService
}

func NewContainerController() interfaces.IContainerController {
	return ContainerController{
		Srv: domain.ContainerSrv,
	}
}

func (ctrl ContainerController) Get(c *gin.Context) {
	uniqueUserID, err := getUniqueUserIDFromJWTInHeader(c)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"err": err.Error()})
	}

	containers, err := ctrl.Srv.GetContainersByUniqueUserID(uniqueUserID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"err": err.Error()})
		return
	}
	c.JSON(http.StatusOK, containers)
}

func (ctrl ContainerController) Post(c *gin.Context) {
	uniqueUserID, err := getUniqueUserIDFromJWTInHeader(c)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"err": err.Error()})
	}

	var containerImage viewmodels.ContainerImage
	err = c.BindJSON(&containerImage)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"err": err.Error()})
	}

	err = ctrl.Srv.CreateContainer(uniqueUserID, containerImage.ImageName)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"err": err.Error()})
	}

	c.Status(http.StatusCreated)
}

func (ctrl ContainerController) Delete(c *gin.Context) {
	uniqueUserID, err := getUniqueUserIDFromJWTInHeader(c)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"err": err.Error()})
		return
	}

	containerIDString := c.Param("id")
	if containerIDString == "" {
		c.JSON(http.StatusBadRequest, gin.H{"err": "コンテナのIDが指定されていません"})
		return
	}

	containerIDInt, err := strconv.Atoi(containerIDString)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"err": err.Error()})
	}

	err = ctrl.Srv.DeleteContainer(uniqueUserID, int64(containerIDInt))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"err": err.Error()})
		return
	}
	c.Status(http.StatusNoContent)
}

// ヘッダーから JWT を取り出す関数
// ヘッダーが空っぽならエラーを返す
func getUniqueUserIDFromJWTInHeader(c *gin.Context) (string, error) {
	userToken := c.GetHeader("Authorization")
	if userToken == "" {
		return "", fmt.Errorf("authorization headerにtokenがありません")
	}
	claim, err := middleware.JWTValidate(userToken)
	if err != nil {
		return "", err
	}
	return claim.Sub, nil
}
