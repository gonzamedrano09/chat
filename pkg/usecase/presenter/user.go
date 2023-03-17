package presenter

import (
	"github.com/gonzamedrano09/chat/pkg/entity"
)

type UserCreateInput struct {
	Username  string `json:"username"`
	Password  string `json:"password"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Gender    string `json:"gender"`
}

type UserUpdateInput struct {
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Gender    string `json:"gender"`
}

type UserInputInterface interface {
	CreateUser(userCreate *UserCreateInput, userOutput UserOutputInterface) error
	RetrieveUser(id uint, userOutput UserOutputInterface) error
	ListUsers(userOutput UserOutputInterface) error
	UpdateUser(id uint, userUpdate *UserUpdateInput, userOutput UserOutputInterface) error
	DestroyUser(id uint, userOutput UserOutputInterface) error
}

type UserOutputInterface interface {
	RenderUser(user *entity.User) error
	RenderUsers(users *[]entity.User) error
}
