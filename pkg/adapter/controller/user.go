package controller

import (
	"github.com/gin-gonic/gin"
	presenterOutput "github.com/gonzamedrano09/chat/pkg/adapter/presenter"
	"github.com/gonzamedrano09/chat/pkg/usecase/presenter"
	"net/http"
	"strconv"
)

type UserHttpControllerInterface interface {
	NewUser(ctx *gin.Context)
	GetUser(ctx *gin.Context)
	GetUsers(ctx *gin.Context)
	UpdateUser(ctx *gin.Context)
	RemoveUser(ctx *gin.Context)
}

type UserHttpController struct {
	UserService presenter.UserInputInterface
}

func NewUserHttpController(userService presenter.UserInputInterface) UserHttpControllerInterface {
	return &UserHttpController{UserService: userService}
}

func (uc *UserHttpController) NewUser(ctx *gin.Context) {
	var user presenter.UserCreateInput
	if err := ctx.BindJSON(&user); err != nil {
		ctx.AbortWithStatus(http.StatusInternalServerError)
		return
	}
	userOutput := presenterOutput.NewUserOutputPort(ctx)
	if err := uc.UserService.CreateUser(&user, userOutput); err != nil {
		ctx.AbortWithStatus(http.StatusBadRequest)
		return
	}
}

func (uc *UserHttpController) GetUser(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		ctx.AbortWithStatus(http.StatusBadRequest)
		return
	}
	userOutput := presenterOutput.NewUserOutputPort(ctx)
	if err = uc.UserService.RetrieveUser(uint(id), userOutput); err != nil {
		ctx.AbortWithStatus(http.StatusBadRequest)
		return
	}
}

func (uc *UserHttpController) GetUsers(ctx *gin.Context) {
	userOutput := presenterOutput.NewUserOutputPort(ctx)
	if err := uc.UserService.ListUsers(userOutput); err != nil {
		ctx.AbortWithStatus(http.StatusBadRequest)
		return
	}
}

func (uc *UserHttpController) UpdateUser(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		ctx.AbortWithStatus(http.StatusBadRequest)
		return
	}
	var user presenter.UserUpdateInput
	if err := ctx.BindJSON(&user); err != nil {
		ctx.AbortWithStatus(http.StatusInternalServerError)
		return
	}
	userOutput := presenterOutput.NewUserOutputPort(ctx)
	if err = uc.UserService.UpdateUser(uint(id), &user, userOutput); err != nil {
		ctx.AbortWithStatus(http.StatusBadRequest)
		return
	}
}

func (uc *UserHttpController) RemoveUser(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		ctx.AbortWithStatus(http.StatusBadRequest)
		return
	}
	userOutput := presenterOutput.NewUserOutputPort(ctx)
	if err = uc.UserService.DestroyUser(uint(id), userOutput); err != nil {
		ctx.AbortWithStatus(http.StatusBadRequest)
		return
	}
}
