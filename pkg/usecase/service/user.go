package service

import (
	"github.com/goccy/go-json"
	"github.com/gonzamedrano09/chat/pkg/entity"
	"github.com/gonzamedrano09/chat/pkg/usecase/presenter"
	"github.com/gonzamedrano09/chat/pkg/usecase/repository"
	"github.com/gonzamedrano09/chat/pkg/usecase/validator"
)

type UserService struct {
	UserRepository repository.UserRepositoryInterface
	UserValidator  validator.UserValidatorInterface
}

func NewUserService(ur repository.UserRepositoryInterface, uv validator.UserValidatorInterface) presenter.UserInputInterface {
	return &UserService{
		UserRepository: ur,
		UserValidator:  uv,
	}
}

func (us *UserService) CreateUser(userCreate *presenter.UserCreateInput, userOutput presenter.UserOutputInterface) error {
	userCreateJson, err := json.Marshal(userCreate)
	if err != nil {
		return err
	}
	var user entity.User
	if err := json.Unmarshal(userCreateJson, &user); err != nil {
		return err
	}
	if err := us.UserValidator.ValidateFieldsToCreate(&user); err != nil {
		return err
	}

	if err := us.UserRepository.InsertUser(&user); err != nil {
		return err
	}
	return userOutput.RenderUser(&user)
}

func (us *UserService) RetrieveUser(id uint, userOutput presenter.UserOutputInterface) error {
	if err := us.UserValidator.ValidateExistingUser(id); err != nil {
		return err
	}

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
	if err := us.UserValidator.ValidateExistingUser(id); err != nil {
		return err
	}
	
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
	if err := us.UserValidator.ValidateFieldsToUpdate(&user); err != nil {
		return err
	}

	if err := us.UserRepository.UpdateUser(&user, id); err != nil {
		return err
	}
	return userOutput.RenderUser(&user)
}

func (us *UserService) DestroyUser(id uint, userOutput presenter.UserOutputInterface) error {
	if err := us.UserValidator.ValidateExistingUser(id); err != nil {
		return err
	}

	if err := us.UserRepository.DeleteUser(id); err != nil {
		return err
	}
	return userOutput.RenderUser(&entity.User{})
}
