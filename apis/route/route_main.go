package route

import (
	"github.com/gin-gonic/gin"
	"simple-login/apis/handlers"
)

func NewRouter() *gin.Engine {

	router := gin.Default()
	userHandler := handlers.NewUserHandler()
	userHandler.RegisterRoutes(router.Group("/user"))

	return router
}
