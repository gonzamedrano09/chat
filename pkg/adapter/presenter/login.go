package presenter

import (
	"github.com/gin-gonic/gin"
	"github.com/gonzamedrano09/chat/pkg/usecase/presenter"
	"net/http"
)

type LoginOutput struct {
	Token string `json:"token"`
}

type LoginOutputPort struct {
	ctx *gin.Context
}

func NewLoginOutputPort(ctx *gin.Context) presenter.LoginOutputInterface {
	return &LoginOutputPort{ctx: ctx}
}

func (lo *LoginOutputPort) RenderToken(token string) error {
	var loginOutput = LoginOutput{Token: token}
	lo.ctx.JSON(http.StatusOK, loginOutput)
	return nil
}
