package registry

import (
	"github.com/gonzamedrano09/chat/pkg/adapter/controller"
)

func (r *Registry) NewMiddlewareDatabaseConnection() controller.MiddlewareDatabaseConnection {
	return controller.NewMiddlewareDatabaseConnection(r.Database)
}
