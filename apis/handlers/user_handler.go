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
	authUseCase services.AuthUseCase
}

func NewUserHandler() UserHandler {
	return UserHandler{
		userUseCase: services.NewUserUseCase(),
		authUseCase: services.NewAuthUseCase(),
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

func (h UserHandler) UserLogin(c *gin.Context) {
	var loginUser dtos.LoginUserDto
	if err := c.BindJSON(&loginUser); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err := h.authUseCase.LoginUser(loginUser)

	if err != nil {
		fmt.Println("Error:", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{})
}

func (h UserHandler) RegisterRoutes(routerGroup *gin.RouterGroup) {
	routerGroup.POST("/signup", h.UserSignup)
	routerGroup.POST("/login", h.UserLogin)
}
