package repository

import (
	"github.com/gesangwidigdo/go-shorturl/interfaces"
	"github.com/gesangwidigdo/go-shorturl/model"
	"gorm.io/gorm"
)

type userRepository struct {
	db *gorm.DB
}

// GetUserByEmail implements interfaces.UserRepository.
func (u *userRepository) GetUserByEmail(email string) (model.User, error) {
	var user model.User
	if err := u.db.Where("email = ?", email).Select("id, email, password").Take(&user).Error; err != nil {
		return model.User{}, err
	}
	return user, nil
}

// Login implements interfaces.UserRepository.
func (u *userRepository) Login(user *model.User) error {
	if err := u.db.Where("email = ? AND password = ?", user.Email, user.Password).Take(&user).Error; err != nil {
		return err
	}
	return nil
}

// Register implements interfaces.UserRepository.
func (u *userRepository) Register(user *model.User) error {
	if err := u.db.Create(&user).Error; err != nil {
		return err
	}
	return nil
}

func NewUserRepository(db *gorm.DB) interfaces.UserRepository {
	return &userRepository{db}
}
