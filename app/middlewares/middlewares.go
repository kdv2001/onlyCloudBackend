package middlewares

import "github.com/gofiber/fiber/v2"

type Middlewares struct {
}

func NewMiddlewares() Middlewares {
	return Middlewares{}
}

func (m *Middlewares) AuthMiddleware(c *fiber.Ctx) error {
	return nil
}
