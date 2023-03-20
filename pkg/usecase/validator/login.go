package validator

import "github.com/gonzamedrano09/chat/pkg/usecase/presenter"

type LoginValidatorInterface interface {
	ValidateFieldsToLogin(login *presenter.LoginInput) error
}
