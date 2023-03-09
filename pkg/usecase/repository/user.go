package repository

import "github.com/gonzamedrano09/chat/pkg/entity"

type UserRepositoryInterface interface {
	InsertUser(user *entity.User) (*entity.User, error)
	SelectOne(user *entity.User, id int) (*entity.User, error)
	SelectAll(users []*entity.User) ([]*entity.User, error)
	UpdateUser(user *entity.User, id int) (*entity.User, error)
	DeleteUser(id int) error
}
