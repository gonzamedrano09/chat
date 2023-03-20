package validator

import (
	"github.com/gonzamedrano09/chat/pkg/entity"
	validationError "github.com/gonzamedrano09/chat/pkg/usecase/error"
	"github.com/gonzamedrano09/chat/pkg/usecase/validator"
	"gorm.io/gorm"
)

type UserValidator struct {
	Database *gorm.DB
}

func NewUserValidator(database *gorm.DB) validator.UserValidatorInterface {
	return &UserValidator{Database: database}
}

func (uv *UserValidator) ValidateFieldsToCreate(user *entity.User) error {
	errs := make(validationError.ValidationError)
	uv.ValidateUsernameLength(user, errs)
	uv.ValidateUsernameExistence(user, errs)
	uv.ValidatePasswordLength(user, errs)
	uv.ValidateFirstNameLength(user, errs)
	uv.ValidateLastNameLength(user, errs)
	uv.ValidateValidGender(user, errs)
	if len(errs) > 0 {
		return errs
	}
	return nil
}

func (uv *UserValidator) ValidateFieldsToUpdate(user *entity.User) error {
	errs := make(validationError.ValidationError)
	uv.ValidateFirstNameLength(user, errs)
	uv.ValidateLastNameLength(user, errs)
	uv.ValidateValidGender(user, errs)
	if len(errs) > 0 {
		return errs
	}
	return nil
}

func (uv *UserValidator) ValidateExistingUser(id uint) error {
	errs := make(validationError.ValidationError)
	if result := uv.Database.Find(&entity.User{}, id); result.RowsAffected < 1 {
		errs["user"] = append(errs["user"], "user not found")
	}
	if len(errs) > 0 {
		return errs
	}
	return nil
}

func (uv *UserValidator) ValidateUsernameLength(user *entity.User, errs validationError.ValidationError) {
	if len(user.Username) < 5 || len(user.Username) > 25 {
		errs["username"] = append(errs["username"], "username length must be more than 5 and less than 25")
	}
}

func (uv *UserValidator) ValidateUsernameExistence(user *entity.User, errs validationError.ValidationError) {
	if result := uv.Database.Find(&entity.User{}, "username = ?", user.Username); result.RowsAffected > 0 {
		errs["username"] = append(errs["username"], "username already use")
	}
}

func (uv *UserValidator) ValidatePasswordLength(user *entity.User, errs validationError.ValidationError) {
	if len(user.Password) < 8 {
		errs["password"] = append(errs["password"], "password length must be more than 8")
	}
}

func (uv *UserValidator) ValidateFirstNameLength(user *entity.User, errs validationError.ValidationError) {
	if len(user.FirstName) < 3 || len(user.FirstName) > 100 {
		errs["first_name"] = append(errs["first_name"], "first_name length must be more than 3 and less than 100")
	}
}

func (uv *UserValidator) ValidateLastNameLength(user *entity.User, errs validationError.ValidationError) {
	if len(user.LastName) < 3 || len(user.LastName) > 100 {
		errs["last_name"] = append(errs["last_name"], "last_name length must be more than 3 and less than 100")
	}
}

func (uv *UserValidator) ValidateValidGender(user *entity.User, errs validationError.ValidationError) {
	if len(user.Gender) > 0 {
		validGender := false
		for _, iteratedGender := range entity.Genders {
			if iteratedGender == user.Gender {
				validGender = true
				break
			}
		}
		if !validGender {
			errs["gender"] = append(errs["gender"], "gender must be Male, Female or Undefined")
		}
	} else {
		user.Gender = "Undefined"
	}
}
