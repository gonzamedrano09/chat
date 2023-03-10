package main

import (
	"github.com/gin-gonic/gin"
	"github.com/gonzamedrano09/chat/pkg/config"
	"github.com/gonzamedrano09/chat/pkg/infrastructure"
	"github.com/gonzamedrano09/chat/pkg/registry"
)

func main() {
	config.LoadConfig()
	database := infrastructure.NewDatabase()

	r := registry.NewRegistry(database)
	app := r.NewAppController()

	router := gin.Default()
	infrastructure.NewRouter(router, &app)
	router.Run("localhost:8080")
}
