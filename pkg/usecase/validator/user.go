package validator

import (
	"github.com/gonzamedrano09/chat/pkg/entity"
)

type UserValidatorInterface interface {
	ValidateFieldsToCreate(user *entity.User) error
	ValidateFieldsToUpdate(user *entity.User) error
	ValidateExistingUser(id uint) error
}
