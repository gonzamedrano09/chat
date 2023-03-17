package validator

import (
	"github.com/gonzamedrano09/chat/pkg/entity"
)

type UserValidatorInterface interface {
	ValidateCreateUser(user *entity.User) error
	ValidateUpdateUser(user *entity.User) error
	ValidateExistingUser(id uint) error
}
