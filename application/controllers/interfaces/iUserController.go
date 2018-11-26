package interfaces

import (
	"github.com/gin-gonic/gin"
)

type IUserController interface {
	Get(*gin.Context)
	Create(*gin.Context)
}
