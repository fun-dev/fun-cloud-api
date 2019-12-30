package controller

import (
	"github.com/fun-dev/fun-cloud-api/internal/container/application/usecase"
	"github.com/fun-dev/fun-cloud-api/internal/container/infrastructure/apperror/ctlerr"
	"github.com/fun-dev/fun-cloud-api/internal/container/infrastructure/server"
	"github.com/fun-dev/fun-cloud-api/pkg/logging"
	"github.com/fun-dev/fun-cloud-api/pkg/term"
	"github.com/gin-gonic/gin"
	"net/http"
)

// ContainerController is
type (
	ContainerController struct {
		usecase.ContainerReadUsecase
		usecase.ContainerCreateUsecase
		usecase.ContainerDeleteUsecase
	}

	PostRequest struct {
		ImageName string `json:"image_name"`
	}
)

func NewContainerController(cRed usecase.ContainerReadUsecase,cCre usecase.ContainerCreateUsecase, cDel usecase.ContainerDeleteUsecase) server.IContainerController {
	return &ContainerController{
		cRed,
		cCre,
		cDel,
	}
}

func (cc ContainerController) Get(c *gin.Context) {
	//var json PostRequest
	////if err := c.ShouldBindJSON(&json); err != nil {
	////	logging.Logf("error: ", err.Error())
	////}
	//TODO: error handling
	//TODO: add option userID and more
	userID := c.GetHeader("Authorization")
	resp, err := cc.ContainerReadUsecase.Execute(userID, term.NullString)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, resp.Entry.Containers)
	return
}
// Post is
// Header: key is Authorization
// BODY: {"image_name": "nginx:latest"}
func (cc ContainerController) Post(c *gin.Context) {
	var json PostRequest
	if err := c.ShouldBindJSON(&json); err != nil {
		logging.Logf("error: ", err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	//TODO: error handling
	userID := c.GetHeader("Authorization")
	if err := cc.ContainerCreateUsecase.Execute(userID, json.ImageName); err != nil {
		logging.Logf("error: ", err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, gin.H{"message": "success on creating container"})
	return
}

// Delete is
func (cc ContainerController) Delete(c *gin.Context) {
	containerID := c.Param("container_id")
	if containerID == term.NullString {
		c.JSON(http.StatusBadRequest, gin.H{"message": ctlerr.ContainerIDCanNotBeFoundOnParam.Error()})
		return
	}
	//TODO: error handling
	userID := c.GetHeader("Authorization")
	if err := cc.ContainerDeleteUsecase.Execute(userID, containerID); err != nil {
		logging.Logf("error: ", err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, gin.H{"message": "success on deleting container"})
	return
}

//func setAuthorizationContext(ginCtx *gin.Context) (context.Context, error) {
//	key := "Authorization"
//	ctx := context.Background()
//	authorization, ok := ginCtx.Get(key)
//	if !ok {
//		return nil, ctlerr.AuthorizationIsNotFoundOnHeader
//	}
//	return context.WithValue(ctx, key, authorization), nil
//}
