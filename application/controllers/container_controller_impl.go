package controllers

import (
	"fmt"
	ictrl "github.com/fun-dev/ccms-poc/application/controllers/interfaces"
	isrv "github.com/fun-dev/ccms-poc/domain/container/interfaces"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

type ContainerController struct {
	Srv isrv.IContainerService
}

func NewContainerController(containerSrv isrv.IContainerService) ictrl.IContainerController {
	return ContainerController{
		Srv: containerSrv,
	}
}

func (ctrl ContainerController) Get(c *gin.Context) {
	uniqueUserID, err := getUniqueUserIDFromJWTInHeader(c)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"err": err.Error()})
		return
	}
	log.Println("api:info:uniqueUserID:", uniqueUserID)

	c.JSON(http.StatusOK, "")
}

func (ctrl ContainerController) Post(c *gin.Context) {
	uniqueUserID, err := getUniqueUserIDFromJWTInHeader(c)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"err": err.Error()})
		return
	}
	c.Status(http.StatusCreated)
}

func (ctrl ContainerController) Delete(c *gin.Context) {
	uniqueUserID, err := getUniqueUserIDFromJWTInHeader(c)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"err": err.Error()})
		return
	}

	containerID := c.Param("id")
	if containerID == "" {
		c.JSON(http.StatusBadRequest, gin.H{"err": "コンテナのIDが指定されていません"})
		return
	}
	c.Status(http.StatusNoContent)
}

// Get Json Web Token from Request Header
func getUniqueUserIDFromJWTInHeader(c *gin.Context) (string, error) {
	userToken := c.GetHeader("Authorization")
	if userToken == "" {
		return "", fmt.Errorf("authorization headerにtokenがありません")
	}
	return "", nil
}
