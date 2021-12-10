package repository

import "9gag-api/model"

type UserRepositoryInterface interface {
	Create(username string, email string, password string) error
	GetByUsername(username string) (*model.User, error)
	GetByUsernameOrEmail(username string, email string) (*model.User, error)
}
