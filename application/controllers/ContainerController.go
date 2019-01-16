package controllers

import (
	"fmt"
	"log"
	"net/http"

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
		fmt.Println("unauthorization err:", err)
		c.JSON(http.StatusUnauthorized, gin.H{"err": err.Error()})
		return
	}
	log.Println("api:info:uniqueUserID:", uniqueUserID)

	containers, err := ctrl.Srv.GetContainersByUniqueUserID(uniqueUserID)
	if err != nil {
		fmt.Println("bad request err:", err)
		c.JSON(http.StatusBadRequest, gin.H{"err": err.Error()})
		return
	}
	c.JSON(http.StatusOK, containers)
}

func (ctrl ContainerController) Post(c *gin.Context) {
	uniqueUserID, err := getUniqueUserIDFromJWTInHeader(c)
	if err != nil {
		fmt.Println("unahtorized err:", err)
		c.JSON(http.StatusUnauthorized, gin.H{"err": err.Error()})
		return
	}

	var containerImage viewmodels.ContainerImage
	err = c.BindJSON(&containerImage)
	if err != nil {
		fmt.Println("bad request err:cant bind json:", err)
		c.JSON(http.StatusBadRequest, gin.H{"err": err.Error()})
		return
	}

	err = ctrl.Srv.CreateContainer(uniqueUserID, containerImage.ImageName)
	if err != nil {
		fmt.Println("internal server err: cant create container:", err)
		c.JSON(http.StatusInternalServerError, gin.H{"err": err.Error()})
		return
	}

	c.Status(http.StatusCreated)
}

func (ctrl ContainerController) Delete(c *gin.Context) {
	uniqueUserID, err := getUniqueUserIDFromJWTInHeader(c)
	if err != nil {
		fmt.Println("unauthorized err", err)
		c.JSON(http.StatusUnauthorized, gin.H{"err": err.Error()})
		return
	}

	containerID := c.Param("id")
	if containerID == "" {
		fmt.Println("bad request err: container id is not specified:", err)
		c.JSON(http.StatusBadRequest, gin.H{"err": "コンテナのIDが指定されていません"})
		return
	}

	err = ctrl.Srv.DeleteContainer(uniqueUserID, containerID)
	if err != nil {
		fmt.Println("internal server err: cant delete container:", err)
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
	sub := claim.Sub
	if sub == "" {
		return "", fmt.Errorf("subの取得に失敗しました．jwtの期限が切れている可能性があります")
	}
	return claim.Sub, nil
}
