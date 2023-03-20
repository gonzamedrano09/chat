package controller

type AppController struct {
	LoginHttpController LoginHttpControllerInterface
	UserHttpController  UserHttpControllerInterface
}
