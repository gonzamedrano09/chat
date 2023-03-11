package service

import (
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
