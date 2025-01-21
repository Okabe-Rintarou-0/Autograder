package request

type LoginRequest struct {
	Identifier string
	Password   string
}

type RegisterRequest struct {
	Username string
	RealName string
	Email    string
	Password string
}
