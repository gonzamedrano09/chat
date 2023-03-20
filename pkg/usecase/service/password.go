package service

import (
	"golang.org/x/crypto/bcrypt"
)

type UserPasswordServiceInterface interface {
	HashPassword(password string) (string, error)
	CheckPassword(password string, hashedPassword string) bool
}

type PasswordService struct {
}

func NewPasswordService() UserPasswordServiceInterface {
	return &PasswordService{}
}

func (ps *PasswordService) HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}

func (ps *PasswordService) CheckPassword(password, hashedPassword string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
	return err == nil
}
