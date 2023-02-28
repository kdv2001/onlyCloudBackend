package handler

import "github.com/gofiber/fiber/v2"

func Register(app *fiber.App, handler *UsersHandler) {
	route := app.Group("users")
	route.Post("/sign-in", handler.SingIn)
}
