package middlewares

import (
	"fmt"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"

	"onlyCloudBackend/app/users"
	"onlyCloudBackend/internal/appErrors"
)

type Middlewares struct {
	auth users.AuthUseCase
}

func NewMiddlewares(auth users.AuthUseCase) Middlewares {
	return Middlewares{auth: auth}
}

func (m *Middlewares) AuthMiddleware(c *fiber.Ctx) error {
	return nil
}

func (m *Middlewares) ErrorHandler() fiber.ErrorHandler {
	return func(ctx *fiber.Ctx, err error) error {
		if err == nil {
			return nil
		}

		appErr := appErrors.AppErrorFromError(err)

		return ctx.Status(appErr.Code).Format(appErr)
	}
}

func (m *Middlewares) LoggingMiddlewares() fiber.Handler {
	logConfig := logger.Config{
		Format: fmt.Sprintf("%s ${time} ${method} ${path} - ${latency} ${status} ${error}\n",
			time.Now().Format("2006/01/02")),
	}

	return logger.New(logConfig)
}
