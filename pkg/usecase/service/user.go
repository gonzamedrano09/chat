package service

import (
	"fmt"
	"github.com/gonzamedrano09/chat/pkg/entity"
	"github.com/gonzamedrano09/chat/pkg/usecase/repository"
)

type UserServiceInterface interface {
	CreateUser(user *entity.User) error
	RetrieveUser(user *entity.User, id uint) error
	ListUsers(users *[]entity.User) error
	UpdateUser(user *entity.User, id uint) error
	DestroyUser(id uint) error
}

type UserService struct {
	UserRepository repository.UserRepositoryInterface
}

func NewUserService(ur repository.UserRepositoryInterface) UserServiceInterface {
	return &UserService{UserRepository: ur}
}

func (us *UserService) CreateUser(user *entity.User) error {
	gender := user.Gender
	if !us.CheckGender(gender) {
		return fmt.Errorf("invalid gender")
	}
	ps := PasswordService{}
	if hashedPassword, err := ps.HashPassword(user.Password); err != nil {
		return fmt.Errorf("password error")
	} else {
		user.Password = hashedPassword
	}
	if err := us.UserRepository.InsertUser(user); err != nil {
		return err
	}
	return nil
}

func (us *UserService) RetrieveUser(user *entity.User, id uint) error {
	if err := us.UserRepository.SelectOne(user, id); err != nil {
		return err
	}
	return nil
}

func (us *UserService) ListUsers(users *[]entity.User) error {
	if err := us.UserRepository.SelectAll(users); err != nil {
		return err
	}
	return nil
}

func (us *UserService) UpdateUser(user *entity.User, id uint) error {
	gender := user.Gender
	if !us.CheckGender(gender) {
		return fmt.Errorf("invalid gender")
	}
	if err := us.UserRepository.UpdateUser(user, id); err != nil {
		return err
	}
	return nil
}

func (us *UserService) DestroyUser(id uint) error {
	if err := us.UserRepository.DeleteUser(id); err != nil {
		return err
	}
	return nil
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
