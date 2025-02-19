package router

import (
	"github.com/gesangwidigdo/go-shorturl/controller"
	"github.com/gesangwidigdo/go-shorturl/interfaces"
	"github.com/gesangwidigdo/go-shorturl/repository"
	"github.com/gesangwidigdo/go-shorturl/service"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func UserRouter(r *gin.RouterGroup, db *gorm.DB) {
	var userRepo interfaces.UserRepository = repository.NewUserRepository(db)
	var userService interfaces.UserService = service.NewUserService(userRepo)
	var userController interfaces.UserController = controller.NewUserController(userService)

	r.POST("/register", userController.Register)
	r.POST("/login", userController.Login)
}
