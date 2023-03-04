package users

type AuthUseCase interface {
	SignIn() error
}

type AuthRepo interface {
	SignIn() error
}
