package controller

import (
	"github.com/gesangwidigdo/go-shorturl/dto"
	"github.com/gesangwidigdo/go-shorturl/interfaces"
	"github.com/gin-gonic/gin"
)

type userController struct {
	userService interfaces.UserService
}

func NewUserController(userService interfaces.UserService) *userController {
	return &userController{userService: userService}
}

func (c *userController) Register(ctx *gin.Context) {
	var userRequest dto.UserRegisterRequest
	if err := ctx.ShouldBindJSON(&userRequest); err != nil {
		ctx.JSON(400, gin.H{"error": err.Error()})
		return
	}

	if err := c.userService.Register(userRequest); err != nil {
		ctx.JSON(400, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(200, gin.H{"message": "success"})
}

func (c *userController) Login(ctx *gin.Context) {
	var userRequest dto.UserLoginRequest
	if err := ctx.ShouldBindJSON(&userRequest); err != nil {
		ctx.JSON(400, gin.H{"error": err.Error()})
		return
	}

	token, err := c.userService.Login(userRequest)
	if err != nil {
		ctx.JSON(400, gin.H{"error": err.Error()})
		return
	}
	ctx.SetCookie("token", token, 3600, "/", "localhost", false, true)
	ctx.JSON(200, gin.H{"message": "success"})
}
