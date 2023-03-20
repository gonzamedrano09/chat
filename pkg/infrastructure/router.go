package infrastructure

import (
	"github.com/gin-gonic/gin"
	"github.com/gonzamedrano09/chat/pkg/adapter/controller"
)

func NewRouter(router *gin.Engine, controller *controller.AppController) *gin.Engine {
	router.POST("/login", controller.LoginHttpController.Login)

	router.POST("/users", controller.UserHttpController.NewUser)
	router.GET("/users/:id", controller.UserHttpController.GetUser)
	router.GET("/users", controller.UserHttpController.GetUsers)
	router.PUT("/users/:id", controller.UserHttpController.UpdateUser)
	router.DELETE("/users/:id", controller.UserHttpController.RemoveUser)
	return router
}
