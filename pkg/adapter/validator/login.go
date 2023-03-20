package validator

import (
	"github.com/gonzamedrano09/chat/pkg/entity"
	validationError "github.com/gonzamedrano09/chat/pkg/usecase/error"
	"github.com/gonzamedrano09/chat/pkg/usecase/presenter"
	"github.com/gonzamedrano09/chat/pkg/usecase/service"
	"github.com/gonzamedrano09/chat/pkg/usecase/validator"
	"gorm.io/gorm"
)

type LoginValidator struct {
	Database *gorm.DB
}

func NewLoginValidator(database *gorm.DB) validator.LoginValidatorInterface {
	return &LoginValidator{Database: database}
}

func (lv *LoginValidator) ValidateFieldsToLogin(login *presenter.LoginInput) error {
	errs := make(validationError.ValidationError)
	lv.ValidateExistingUser(login, errs)
	lv.ValidateCorrectPassword(login, errs)
	if len(errs) > 0 {
		return errs
	}
	return nil
}

func (lv *LoginValidator) ValidateExistingUser(login *presenter.LoginInput, errs validationError.ValidationError) {
	if result := lv.Database.Find(&entity.User{}, "username = ?", login.Username); result.RowsAffected < 1 {
		errs["username"] = append(errs["username"], "non-existent user")
	}
}

func (lv *LoginValidator) ValidateCorrectPassword(login *presenter.LoginInput, errs validationError.ValidationError) {
	var user entity.User
	if result := lv.Database.Find(&user, "username = ?", login.Username); result.RowsAffected < 1 {
		errs["password"] = append(errs["password"], "incorrect password")
		return
	}
	ps := service.NewPasswordService()
	if ok := ps.CheckPassword(login.Password, user.Password); !ok {
		errs["password"] = append(errs["password"], "incorrect password")
	}
}
