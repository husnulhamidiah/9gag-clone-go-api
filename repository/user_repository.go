package repository

import (
	"9gag-api/model"
)

type UserRepository struct {

}

func NewUserRepository() *UserRepository {
	return &UserRepository{}
}

func (r *UserRepository) Create(username string, email string, password string) error{
	user := model.User{
		Username: username,
		Email:    email,
		Password: password,
	}
	if result := model.DB.Model(&model.User{}).Save(&user); result.Error != nil {
		return result.Error
	}
	return nil
}

func (r *UserRepository) GetByUsername(username string) (*model.User, error) {
	var user model.User
	result := model.DB.Where("username = ?", username).First(&user)
	if result.Error != nil {
		return nil, result.Error
	}
	return &user, nil
}

func (r *UserRepository) GetByUsernameOrEmail(username string, email string) (*model.User, error) {
	var user model.User
	result := model.DB.Where("username = ? or email = ?", username, email).First(&user)
	if result.Error != nil {
		return nil, result.Error
	}
	return &user, nil
}
