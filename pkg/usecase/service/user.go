package service

import (
	"fmt"
	"github.com/goccy/go-json"
	"github.com/gonzamedrano09/chat/pkg/entity"
	"github.com/gonzamedrano09/chat/pkg/usecase/presenter"
	"github.com/gonzamedrano09/chat/pkg/usecase/repository"
)

type UserService struct {
	UserRepository repository.UserRepositoryInterface
}

func NewUserService(ur repository.UserRepositoryInterface) presenter.UserInputInterface {
	return &UserService{
		UserRepository: ur,
	}
}

func (us *UserService) CreateUser(userCreate *presenter.UserCreateInput, userOutput presenter.UserOutputInterface) error {
	userCreateJson, err := json.Marshal(userCreate)
	if err != nil {
		return err
	}
	var user entity.User
	if err = json.Unmarshal(userCreateJson, &user); err != nil {
		return err
	}

	if !us.CheckGender(user.Gender) {
		return fmt.Errorf("invalid gender")
	}
	ps := PasswordService{}
	if hashedPassword, err := ps.HashPassword(user.Password); err != nil {
		return err
	} else {
		user.Password = hashedPassword
	}

	if err := us.UserRepository.InsertUser(&user); err != nil {
		return err
	}
	return userOutput.RenderUser(&user)
}

func (us *UserService) RetrieveUser(id uint, userOutput presenter.UserOutputInterface) error {
	var user entity.User
	if err := us.UserRepository.SelectOne(&user, id); err != nil {
		return err
	}
	return userOutput.RenderUser(&user)
}

func (us *UserService) ListUsers(userOutput presenter.UserOutputInterface) error {
	var users []entity.User
	if err := us.UserRepository.SelectAll(&users); err != nil {
		return err
	}
	return userOutput.RenderUsers(&users)
}

func (us *UserService) UpdateUser(id uint, userUpdate *presenter.UserUpdateInput, userOutput presenter.UserOutputInterface) error {
	userUpdateJson, err := json.Marshal(userUpdate)
	if err != nil {
		return err
	}
	var user entity.User
	if err := us.UserRepository.SelectOne(&user, id); err != nil {
		return err
	}
	if err = json.Unmarshal(userUpdateJson, &user); err != nil {
		return err
	}

	if !us.CheckGender(user.Gender) {
		return fmt.Errorf("invalid gender")
	}

	if err := us.UserRepository.UpdateUser(&user, id); err != nil {
		return err
	}
	return userOutput.RenderUser(&user)
}

func (us *UserService) DestroyUser(id uint, userOutput presenter.UserOutputInterface) error {
	if err := us.UserRepository.DeleteUser(id); err != nil {
		return err
	}
	return userOutput.RenderUser(&entity.User{})
}

// Validations

func (us *UserService) CheckGender(gender string) bool {
	for _, iteratedGender := range entity.Genders {
		if gender == iteratedGender {
			return true
		}
	}
	return false
}
