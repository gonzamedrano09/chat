package registry

import (
	"github.com/gonzamedrano09/chat/pkg/adapter/controller"
	"github.com/gonzamedrano09/chat/pkg/adapter/repository"
	"github.com/gonzamedrano09/chat/pkg/usecase/service"
)

func (r *Registry) NewUserController() controller.UserControllerInterface {
	userService := service.NewUserService(repository.NewUserRepository(r.Database))
	return controller.NewUserController(userService)
}
