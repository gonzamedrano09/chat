package repository

import (
	"github.com/gonzamedrano09/chat/pkg/entity"
	"github.com/gonzamedrano09/chat/pkg/usecase/repository"
	"gorm.io/gorm"
)

type UserRepository struct {
	Database *gorm.DB
}

func NewUserRepository(database *gorm.DB) repository.UserRepositoryInterface {
	return &UserRepository{Database: database}
}

func (ur *UserRepository) InsertUser(user *entity.User) error {
	if r := ur.Database.Create(user); r.Error != nil {
		return r.Error
	}
	return nil
}

func (ur *UserRepository) SelectOne(user *entity.User, id uint) error {
	if r := ur.Database.Find(user, id); r.Error != nil {
		return r.Error
	}
	return nil
}

func (ur *UserRepository) SelectAll(users *[]entity.User) error {
	if r := ur.Database.Find(users); r.Error != nil {
		return r.Error
	}
	return nil
}

func (ur *UserRepository) UpdateUser(user *entity.User, id uint) error {
	if r := ur.Database.Save(user); r.Error != nil {
		return r.Error
	}
	return nil
}

func (ur *UserRepository) DeleteUser(id uint) error {
	if r := ur.Database.Delete(&entity.User{}, id); r.Error != nil {
		return r.Error
	}
	return nil
}
