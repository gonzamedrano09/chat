package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/gonzamedrano09/chat/pkg/entity"
	"github.com/gonzamedrano09/chat/pkg/usecase/service"
	"net/http"
	"strconv"
)

type UserControllerInterface interface {
	NewUser(ctx *gin.Context)
	GetUser(ctx *gin.Context)
	GetUsers(ctx *gin.Context)
	UpdateUser(ctx *gin.Context)
	RemoveUser(ctx *gin.Context)
}

type UserController struct {
	UserService service.UserServiceInterface
}

func NewUserController(userService service.UserServiceInterface) UserControllerInterface {
	return &UserController{UserService: userService}
}

func (uc *UserController) NewUser(ctx *gin.Context) {
	var user entity.User
	err := ctx.BindJSON(&user)
	if err != nil {
		return
	}
	_, err = uc.UserService.CreateUser(&user)
	if err != nil {
		return
	}
	ctx.IndentedJSON(http.StatusCreated, user)
}

func (uc *UserController) GetUser(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		return
	}
	var user entity.User
	_, err = uc.UserService.RetrieveUser(&user, id)
	if err != nil {
		return
	}
	ctx.IndentedJSON(http.StatusOK, user)
}

func (uc *UserController) GetUsers(ctx *gin.Context) {
	var users []*entity.User
	_, err := uc.UserService.ListUsers(users)
	if err != nil {
		return
	}
	ctx.IndentedJSON(http.StatusOK, users)
}

func (uc *UserController) UpdateUser(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		return
	}
	var user entity.User
	_, err = uc.UserService.UpdateUser(&user, id)
	if err != nil {
		return
	}
	ctx.IndentedJSON(http.StatusOK, user)
}

func (uc *UserController) RemoveUser(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		return
	}
	err = uc.UserService.DestroyUser(id)
	if err != nil {
		return
	}
	ctx.IndentedJSON(http.StatusNoContent, nil)
}
