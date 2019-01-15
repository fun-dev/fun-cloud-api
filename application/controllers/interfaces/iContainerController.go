package interfaces

import (
	"github.com/gin-gonic/gin"
)

type IContainerController interface {
	Get(*gin.Context)
	Post(*gin.Context)
	Delete(*gin.Context)
}
