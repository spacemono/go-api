package command

type CreateUser struct {
	Username string `validate:"required,min=5,max=32"`
	Name     string `validate:"required,min=3,max=48"`
	Password string `validate:"required,min=8,max=32"`
}

type CreateUserResult struct {
	Username string
	Name     string
	Password string
}
