package service

import (
	"github.com/gonzamedrano09/chat/pkg/entity"
	"github.com/gonzamedrano09/chat/pkg/usecase/repository"
)

type UserServiceInterface interface {
	CreateUser(user *entity.User) (*entity.User, error)
	RetrieveUser(user *entity.User, id int) (*entity.User, error)
	ListUsers(users []*entity.User) ([]*entity.User, error)
	UpdateUser(user *entity.User, id int) (*entity.User, error)
	DestroyUser(id int) error
}

type UserService struct {
	UserRepository repository.UserRepositoryInterface
}

func NewUserService(ur repository.UserRepositoryInterface) UserServiceInterface {
	return &UserService{UserRepository: ur}
}

func (us *UserService) CreateUser(user *entity.User) (*entity.User, error) {
	_, err := us.UserRepository.InsertUser(user)
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (us *UserService) RetrieveUser(user *entity.User, id int) (*entity.User, error) {
	_, err := us.UserRepository.SelectOne(user, id)
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (us *UserService) ListUsers(users []*entity.User) ([]*entity.User, error) {
	_, err := us.UserRepository.SelectAll(users)
	if err != nil {
		return nil, err
	}
	return users, nil
}

func (us *UserService) UpdateUser(user *entity.User, id int) (*entity.User, error) {
	_, err := us.UserRepository.UpdateUser(user, id)
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (us *UserService) DestroyUser(id int) error {
	err := us.UserRepository.DeleteUser(id)
	if err != nil {
		return err
	}
	return nil
}
