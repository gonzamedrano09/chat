package controller

import (
	"github.com/gin-gonic/gin"
	presenterOutput "github.com/gonzamedrano09/chat/pkg/adapter/presenter"
	"github.com/gonzamedrano09/chat/pkg/usecase/presenter"
	"net/http"
)

type LoginHttpControllerInterface interface {
	Login(ctx *gin.Context)
}

type LoginHttpController struct {
	LoginService presenter.LoginInputInterface
}

func NewLoginHttpController(loginService presenter.LoginInputInterface) LoginHttpControllerInterface {
	return &LoginHttpController{LoginService: loginService}
}

func (lc *LoginHttpController) Login(ctx *gin.Context) {
	var login presenter.LoginInput
	if err := ctx.BindJSON(&login); err != nil {
		ctx.AbortWithStatus(http.StatusInternalServerError)
		return
	}
	loginOutput := presenterOutput.NewLoginOutputPort(ctx)
	if err := lc.LoginService.Login(&login, loginOutput); err != nil {
		ctx.JSON(http.StatusBadRequest, err)
		return
	}
}
