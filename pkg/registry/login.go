package registry

import (
	"github.com/gonzamedrano09/chat/pkg/adapter/controller"
	"github.com/gonzamedrano09/chat/pkg/adapter/validator"
	"github.com/gonzamedrano09/chat/pkg/usecase/service"
)

func (r *Registry) NewLoginHttpController() controller.LoginHttpControllerInterface {
	loginService := service.NewLoginService(validator.NewLoginValidator(r.Database))
	return controller.NewLoginHttpController(loginService)
}
