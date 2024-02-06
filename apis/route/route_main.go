package route

import (
	"github.com/gin-gonic/gin"
)

type RouteRegistrer interface {
	RegisterRoutes(router *gin.RouterGroup, basePath string)
}

func NewRouter() *gin.Engine {

	router := gin.Default()

	return router
}
