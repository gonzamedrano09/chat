package presenter

type LoginInput struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type LoginInputInterface interface {
	Login(login *LoginInput, loginOutput LoginOutputInterface) error
}

type LoginOutputInterface interface {
	RenderToken(token string) error
}
