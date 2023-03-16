package presenter

import (
	"github.com/gonzamedrano09/chat/pkg/entity"
	"time"
)

type UserCreateInput struct {
	Username    string    `json:"username"`
	Password    string    `json:"password"`
	FirstName   string    `json:"first_name,omitempty"`
	LastName    string    `json:"last_name,omitempty"`
	Gender      string    `json:"gender,omitempty"`
	DateOfBirth time.Time `json:"date_of_birth,omitempty"`
}

type UserUpdateInput struct {
	FirstName   string    `json:"first_name,omitempty"`
	LastName    string    `json:"last_name,omitempty"`
	Gender      string    `json:"gender,omitempty"`
	DateOfBirth time.Time `json:"date_of_birth,omitempty"`
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
