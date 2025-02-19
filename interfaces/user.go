package interfaces

import (
	"github.com/gesangwidigdo/go-shorturl/dto"
	"github.com/gesangwidigdo/go-shorturl/model"
	"github.com/gin-gonic/gin"
)

type UserRepository interface {
	Register(user *model.User) error
	Login(user *model.User) error
	GetUserByEmail(email string) (model.User, error)
}

type UserService interface {
	Register(request dto.UserRegisterRequest) error
	Login(request dto.UserLoginRequest) (string, error)
}

type UserController interface {
	Register(ctx *gin.Context)
	Login(ctx *gin.Context)
}
