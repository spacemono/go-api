package command

type CreateUser struct {
	Username string `validate:"required,min=5,max=32"`
	Email    string `validate:"required,email"`
	Password string `validate:"required,min=8,max=32"`
}

type CreateUserResult struct {
	Username string
	Email    string
	Password string
}
