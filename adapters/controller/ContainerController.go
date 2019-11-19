package controller

import (
	"context"
	"github.com/fun-dev/ccms/application/usecase"
	"github.com/fun-dev/ccms/infrastructure/apperror/ctlerr"
	"github.com/gin-gonic/gin"
	"net/http"
)

// ContainerController is
type (
	ContainerController struct {
		usecase.ContainerCreateUsecase
		usecase.ContainerDeleteUsecase
	}
	PostRequest struct {
		ImageName string `json:"image_name"`
	}
)

func NewContainerController(cCre usecase.ContainerCreateUsecase, cDel usecase.ContainerDeleteUsecase) *ContainerController {
	return &ContainerController{
		cCre,
		cDel,
	}
}

// Post is
// Header: key is Authorization
// BODY: {"image_name": "nginx:latest"}
func (cc ContainerController) Post(c *gin.Context) {
	ctx, err := setAuthorizationContext(c)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	var json PostRequest
	if err := c.ShouldBindJSON(&json); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	// TODO: fix to common value
	key := "USER_ID"
	userID, ok := c.Get(key)
	if !ok {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "sorry fix immediately"})
		return
	}
	if err := cc.ContainerCreateUsecase.Execute(ctx, userID.(string), json.ImageName); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, gin.H{"message": "success on creating container"})
	return
}

// Delete is
func (cc ContainerController) Delete(c *gin.Context) {
	containerID := c.Param("container_id")
	if containerID == "" {
		c.JSON(http.StatusBadRequest, gin.H{"message": ctlerr.ContainerIDCanNotBeFoundOnParam.Error()})
		return
	}
	// Authentication Information
	ctx, err := setAuthorizationContext(c)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if err := cc.ContainerDeleteUsecase.Execute(ctx, containerID); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, gin.H{"message": "success on deleting container"})
	return
}

func setAuthorizationContext(ginCtx *gin.Context) (context.Context, error) {
	key := "Authorization"
	ctx := context.Background()
	authorization, ok := ginCtx.Get(key)
	if !ok {
		return nil, ctlerr.AuthorizationIsNotFoundOnHeader
	}
	return context.WithValue(ctx, key, authorization), nil
}
