package presenter

import (
	"github.com/gin-gonic/gin"
	"github.com/goccy/go-json"
	"github.com/gonzamedrano09/chat/pkg/entity"
	"github.com/gonzamedrano09/chat/pkg/usecase/presenter"
	"net/http"
	"time"
)

type UserOutput struct {
	ID          uint      `json:"id,omitempty"`
	Username    string    `json:"username,omitempty"`
	FirstName   string    `json:"first_name,omitempty"`
	LastName    string    `json:"last_name,omitempty"`
	Gender      string    `json:"gender,omitempty"`
	DateOfBirth time.Time `json:"date_of_birth,omitempty"`
}

type UserOutputPort struct {
	ctx *gin.Context
}

func NewUserOutputPort(ctx *gin.Context) presenter.UserOutputInterface {
	return &UserOutputPort{ctx: ctx}
}

func (uop *UserOutputPort) ParseUser(user *entity.User, userOutput *UserOutput) error {
	userJson, err := json.Marshal(user)
	if err != nil {
		return err
	}
	if err = json.Unmarshal(userJson, userOutput); err != nil {
		return err
	}
	return nil
}

func (uop *UserOutputPort) RenderUser(user *entity.User) error {
	var userOutput UserOutput
	if err := uop.ParseUser(user, &userOutput); err != nil {
		return err
	}
	uop.ctx.JSON(http.StatusOK, userOutput)
	return nil
}

func (uop *UserOutputPort) RenderUsers(users *[]entity.User) error {
	var usersOutput []UserOutput
	for _, user := range *users {
		var userOutput UserOutput
		if err := uop.ParseUser(&user, &userOutput); err != nil {
			return err
		}
		usersOutput = append(usersOutput, userOutput)
	}
	uop.ctx.JSON(http.StatusOK, usersOutput)
	return nil
}
