package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/gonzamedrano09/chat/pkg/entity"
	"github.com/gonzamedrano09/chat/pkg/usecase/service"
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
	UserService service.UserServiceInterface
}

func NewUserHttpController(userService service.UserServiceInterface) UserHttpControllerInterface {
	return &UserHttpController{UserService: userService}
}

func (uc *UserHttpController) NewUser(ctx *gin.Context) {
	var user entity.User
	if err := ctx.BindJSON(&user); err != nil {
		ctx.AbortWithStatus(http.StatusInternalServerError)
		return
	}
	if err := uc.UserService.CreateUser(&user); err != nil {
		ctx.AbortWithStatus(http.StatusBadRequest)
		return
	}
	ctx.JSON(http.StatusCreated, user)
}

func (uc *UserHttpController) GetUser(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		ctx.AbortWithStatus(http.StatusBadRequest)
		return
	}
	var user entity.User
	if err = uc.UserService.RetrieveUser(&user, uint(id)); err != nil {
		ctx.AbortWithStatus(http.StatusBadRequest)
		return
	}
	ctx.JSON(http.StatusOK, user)
}

func (uc *UserHttpController) GetUsers(ctx *gin.Context) {
	var users []entity.User
	if err := uc.UserService.ListUsers(&users); err != nil {
		ctx.AbortWithStatus(http.StatusBadRequest)
		return
	}
	ctx.JSON(http.StatusOK, users)
}

func (uc *UserHttpController) UpdateUser(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		ctx.AbortWithStatus(http.StatusBadRequest)
		return
	}
	var user entity.User
	if err = uc.UserService.RetrieveUser(&user, uint(id)); err != nil {
		ctx.AbortWithStatus(http.StatusBadRequest)
		return
	}
	if err := ctx.BindJSON(&user); err != nil {
		ctx.AbortWithStatus(http.StatusInternalServerError)
		return
	}
	if err = uc.UserService.UpdateUser(&user, uint(id)); err != nil {
		ctx.AbortWithStatus(http.StatusBadRequest)
		return
	}
	ctx.JSON(http.StatusOK, user)
}

func (uc *UserHttpController) RemoveUser(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		ctx.AbortWithStatus(http.StatusBadRequest)
		return
	}
	if err = uc.UserService.DestroyUser(uint(id)); err != nil {
		ctx.AbortWithStatus(http.StatusBadRequest)
		return
	}
	ctx.Status(http.StatusNoContent)
}
