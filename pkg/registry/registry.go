package registry

import (
	"github.com/gonzamedrano09/chat/pkg/adapter/controller"
	"gorm.io/gorm"
)

type RegistryInterface interface {
	NewAppController() controller.AppController
}

type Registry struct {
	Database *gorm.DB
}

func NewRegistry(database *gorm.DB) RegistryInterface {
	return &Registry{Database: database}
}

func (r *Registry) NewAppController() controller.AppController {
	return controller.AppController{
		MiddlewareDatabaseConnection: r.NewMiddlewareDatabaseConnection(),
		LoginHttpController:          r.NewLoginHttpController(),
		UserHttpController:           r.NewUserHttpController(),
	}
}
