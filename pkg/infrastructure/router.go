package infrastructure

import (
	"github.com/gin-gonic/gin"
	"github.com/gonzamedrano09/chat/pkg/adapter/controller"
)

func NewRouter(router *gin.Engine, controller *controller.AppController) *gin.Engine {
	router.POST("/users", controller.UserController.NewUser)
	router.GET("/users/:id", controller.UserController.GetUser)
	router.GET("/users", controller.UserController.GetUsers)
	router.PUT("/users/:id", controller.UserController.UpdateUser)
	router.DELETE("/users/:id", controller.UserController.RemoveUser)
	return router
}
