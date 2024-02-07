package handlers

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"simple-login/dtos"
	"simple-login/services"
)

type UserHandler struct {
	userUseCase services.UserUseCase
}

func NewUserHandler() UserHandler {
	return UserHandler{
		userUseCase: services.NewUserUseCase(),
	}
}

func (h UserHandler) UserSignup(c *gin.Context) {
	var createUser dtos.CreateUserDto
	if err := c.BindJSON(&createUser); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err := h.userUseCase.CreateUser(createUser)

	if err != nil {
		fmt.Println("Error:", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{})
}

func (h UserHandler) RegisterRoutes(routerGroup *gin.RouterGroup) {
	routerGroup.POST("/signup", h.UserSignup)
}
