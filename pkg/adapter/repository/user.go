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

func (ur *UserRepository) InsertUser(user *entity.User) (*entity.User, error) {
	r := ur.Database.Create(user)
	if err := r.Error; err != nil {
		return nil, err
	}
	return user, nil
}

func (ur *UserRepository) SelectOne(user *entity.User, id int) (*entity.User, error) {
	r := ur.Database.Find(user, id)
	if err := r.Error; err != nil {
		return nil, err
	}
	return user, nil
}

func (ur *UserRepository) SelectAll(users []*entity.User) ([]*entity.User, error) {
	r := ur.Database.Find(users)
	if err := r.Error; err != nil {
		return nil, err
	}
	return users, nil
}

func (ur *UserRepository) UpdateUser(user *entity.User, id int) (*entity.User, error) {
	r := ur.Database.Save(user)
	if err := r.Error; err != nil {
		return nil, err
	}
	return user, nil
}

func (ur *UserRepository) DeleteUser(id int) error {
	r := ur.Database.Delete(&entity.User{}, id)
	if err := r.Error; err != nil {
		return err
	}
	return nil
}
