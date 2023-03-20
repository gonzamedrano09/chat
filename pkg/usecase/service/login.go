package service

import (
	"github.com/gonzamedrano09/chat/pkg/usecase/presenter"
	"github.com/gonzamedrano09/chat/pkg/usecase/validator"
)

type LoginService struct {
	LoginValidator validator.LoginValidatorInterface
}

func NewLoginService(lv validator.LoginValidatorInterface) presenter.LoginInputInterface {
	return &LoginService{LoginValidator: lv}
}

func (ls *LoginService) Login(login *presenter.LoginInput, loginOutput presenter.LoginOutputInterface) error {
	if err := ls.LoginValidator.ValidateFieldsToLogin(login); err != nil {
		return err
	}
	js := NewJWTService()
	tokenString, err := js.GenerateJWT(login.Username)
	if err != nil {
		return err
	}
	return loginOutput.RenderToken(tokenString)
}
