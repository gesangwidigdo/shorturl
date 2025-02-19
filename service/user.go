package service

import (
	"errors"
	"strings"

	"github.com/gesangwidigdo/go-shorturl/dto"
	"github.com/gesangwidigdo/go-shorturl/interfaces"
	"github.com/gesangwidigdo/go-shorturl/model"
	"github.com/gesangwidigdo/go-shorturl/utils"
)

type userService struct {
	userRepo interfaces.UserRepository
}

// Login implements interfaces.UserService.
func (u *userService) Login(request dto.UserLoginRequest) (string, error) {
	if request.Email == "" || request.Password == "" {
		return "", errors.New("email or password is empty")
	}

	userFound, err := u.userRepo.GetUserByEmail(request.Email)
	if err != nil {
		return  "", errors.New("username does not exist, please register first")
	}

	if !utils.CheckPasswordHash(userFound.Password, request.Password) {
		return "", errors.New("password is incorrect")
	}

	tokenString, err := utils.CreateToken(userFound)
	if err != nil {
		return "", err
	}

	return tokenString, nil
}

// Register implements interfaces.UserService.
func (u *userService) Register(request dto.UserRegisterRequest) error {
	hashedPass, err := utils.HashPassword(request.Password)
	if err != nil {
		return err
	}
	newUser := model.User{
		Username: request.Username,
		Email:    request.Email,
		Password: hashedPass,
	}

	if err := u.userRepo.Register(&newUser); err != nil {
		if strings.Contains(err.Error(), "Duplicate Entry") {
			if strings.Contains(err.Error(), "username") {
				return errors.New("username already exists")
			} else if strings.Contains(err.Error(), "email") {
				return errors.New("email already exists")
			}
		}
		return err
	}

	return nil
}

func NewUserService(userRepo interfaces.UserRepository) interfaces.UserService {
	return &userService{
		userRepo,
	}
}
