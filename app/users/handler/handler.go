package handler

import (
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
)

type UsersHandler struct {
	usersValidate *validator.Validate
}

func NewUsershandler() *UsersHandler {
	userValidator := validator.New()
	// регистрация кастомных типов валидации
	RegisterCustomTypes(userValidator)

	return &UsersHandler{usersValidate: userValidator}
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
// @Summary Авторизация пользователя
// @Tags auth
// @Accept  json
// @Produce  json
// @Param data body users.SignIn true "Данные для авторизации"
// @Success 200 {object} users.User
// @Failure 400 {object} appErrors.AppError
// @Failure 401 {object} appErrors.AppError
// @Failure 500 {object} appErrors.AppError
// @Router /users/sign-in [post]
func (h *UsersHandler) SingIn(c *fiber.Ctx) error {
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

// Sessions godoc
// @Summary Получение сессий пользователя
// @Tags auth
// @Produce  json
// @Success 200 {object}
// @Failure 400 {object} appErrors.AppError
// @Failure 401 {object} appErrors.AppError
// @Failure 500 {object} appErrors.AppError
// @Router /users/sessions [get]
func (h *UsersHandler) Sessions(c fiber.Ctx) error {
	return nil
}
