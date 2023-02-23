package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/swagger"
	"go.uber.org/zap"

	handler2 "onlyCloudBackend/app/files/handler"
	"onlyCloudBackend/app/users/handler"
	_ "onlyCloudBackend/swagger"
)

type AppServer struct {
	app        *fiber.App
	logger     *zap.Logger
	listenAddr string
}

func NewServer(logger *zap.Logger, listenAddr string) AppServer {
	app := fiber.New()

	handler.Register(app, handler.NewUsershandler())
	handler2.Register(app, handler2.NewFilesHandler())

	app.Get("/swagger/*", swagger.New(swagger.Config{
		PersistAuthorization: true,
	}))

	return AppServer{app: app, logger: logger, listenAddr: listenAddr}
}

// @title OnlyCloud
// @version 1.0
// @description Modern cloud service
// @BasePath /
func (a *AppServer) Run() {
	a.logger.Sugar().Fatal(a.app.Listen(a.listenAddr))
}
