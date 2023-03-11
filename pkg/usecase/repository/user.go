package repository

import "github.com/gonzamedrano09/chat/pkg/entity"

type UserRepositoryInterface interface {
	InsertUser(user *entity.User) error
	SelectOne(user *entity.User, id uint) error
	SelectAll(users *[]entity.User) error
	UpdateUser(user *entity.User, id uint) error
	DeleteUser(id uint) error
}
