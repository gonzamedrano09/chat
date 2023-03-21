package controller

type AppController struct {
	MiddlewareDatabaseConnection MiddlewareDatabaseConnection
	LoginHttpController          LoginHttpControllerInterface
	UserHttpController           UserHttpControllerInterface
}
