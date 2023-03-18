package registry

import (
	"github.com/gonzamedrano09/chat/pkg/adapter/controller"
	"github.com/gonzamedrano09/chat/pkg/adapter/repository"
	"github.com/gonzamedrano09/chat/pkg/adapter/validator"
	"github.com/gonzamedrano09/chat/pkg/usecase/service"
)

func (r *Registry) NewUserHttpController() controller.UserHttpControllerInterface {
	userService := service.NewUserService(
		repository.NewUserRepository(r.Database),
		validator.NewUserValidator(r.Database))
	return controller.NewUserHttpController(userService)
}
