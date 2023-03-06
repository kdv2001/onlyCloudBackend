package usecase

import "onlyCloudBackend/app/users"

type usersUseCase struct {
	usersRepo users.AuthRepo
}

func NewUsersUseCase(usersRepo users.AuthRepo) *usersUseCase {
	return &usersUseCase{usersRepo: usersRepo}
}

func (uc *usersUseCase) SignIn() error {
	return nil
}
