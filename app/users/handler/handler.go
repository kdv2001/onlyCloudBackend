package handler

import "github.com/gofiber/fiber/v2"

type UsersHandler struct{}

func NewUsershandler() *UsersHandler {
	return &UsersHandler{}
}

// SingUp godoc
// @Summary Регистрация пользователя
// @Tags auth
// @Accept  json
// @Produce  json
// @Param data body users.User true "Данные для регистрации"
// @Success 200 {object} users.SingUp
// @Failure 400 {object} appErrors.AppError
// @Failure 500 {object} appErrors.AppError
// @Router /users/sign-up [post]
func (h *UsersHandler) SingUp(c fiber.Ctx) error {
	return nil
}

// SingIn godoc
// @Summary Регистрация пользователя
// @Tags auth
// @Accept  json
// @Produce  json
// @Param data body users.SignIn true "Данные для авторизации"
// @Success 200 {object} users.User
// @Failure 400 {object} appErrors.AppError
// @Failure 401 {object} appErrors.AppError
// @Failure 500 {object} appErrors.AppError
// @Router /users/sign-in [post]
func (h *UsersHandler) SingIn(c fiber.Ctx) error {
	return nil
}

// Authorize godoc
// @Summary Проверка авторизации по Cookie
// @Tags auth
// @Produce  json
// @Success 200 {object} users.User
// @Failure 400 {object} appErrors.AppError
// @Failure 401 {object} appErrors.AppError
// @Failure 500 {object} appErrors.AppError
// @Router /users [get]
func (h *UsersHandler) Authorize(c fiber.Ctx) error {
	return nil
}
