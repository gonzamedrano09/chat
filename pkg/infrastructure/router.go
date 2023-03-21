package infrastructure

import (
	"github.com/gin-gonic/gin"
	"github.com/gonzamedrano09/chat/pkg/adapter/controller"
)

func NewRouter(router *gin.Engine, controller *controller.AppController) *gin.Engine {
	router.POST("/login", controller.LoginHttpController.Login)
	router.POST("/users", controller.UserHttpController.NewUser)

	authorizedGroup := router.Group("/")
	authorizedGroup.Use(Authorization(controller.MiddlewareDatabaseConnection))
	{
		authorizedGroup.GET("/users/:id", controller.UserHttpController.GetUser)
		authorizedGroup.GET("/users", controller.UserHttpController.GetUsers)
		authorizedGroup.PUT("/users/:id", controller.UserHttpController.UpdateUser)
		authorizedGroup.DELETE("/users/:id", controller.UserHttpController.RemoveUser)
	}
	return router
}
