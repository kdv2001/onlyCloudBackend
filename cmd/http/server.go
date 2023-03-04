package main

import (
	"context"
	"fmt"
	"net"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/swagger"
	"github.com/jackc/pgx/v5"
	"github.com/spf13/viper"
	"go.uber.org/zap"

	filesHandler "onlyCloudBackend/app/files/handler"
	"onlyCloudBackend/app/middlewares"
	usersHandler "onlyCloudBackend/app/users/handler"
	"onlyCloudBackend/app/users/repository"
	"onlyCloudBackend/app/users/usecase"
	_ "onlyCloudBackend/swagger"
)

type AppServer struct {
	app        *fiber.App
	logger     *zap.Logger
	listenAddr string
}

func NewServer(logger *zap.Logger, listenAddr string) AppServer {
	// postgres
	postgresUser := viper.GetString("postgresql.Username")
	postgresPassword := viper.GetString("postgresql.Password")
	postgresPort := viper.GetString("postgresql.Port")
	postgresDbAddress := viper.GetString("postgresql.Address")
	postgresDbName := viper.GetString("postgresql.dbName")
	postgresAddress := fmt.Sprintf("postgres://%s:%s@%s/%s",
		postgresUser, postgresPassword, net.JoinHostPort(postgresDbAddress, postgresPort), postgresDbName)

	conn, err := pgx.Connect(context.Background(), postgresAddress)
	if err != nil {
		logger.Sugar().Fatal(err)
	}

	// users
	usersRepository := repository.NewPostgresUsersRepository(conn)
	usersUseCase := usecase.NewUsersUseCase(usersRepository)

	// middlewares
	mw := middlewares.NewMiddlewares(usersUseCase)

	fiberConfig := fiber.Config{
		ErrorHandler: mw.ErrorHandler(),
	}

	app := fiber.New(fiberConfig)
	app.Use(mw.LoggingMiddlewares())

	// register HTTP
	filesHandler.Register(app, filesHandler.NewFilesHandler())
	usersHandler.Register(app, usersHandler.NewUsershandler(usersUseCase))

	app.Get("/swagger/*", swagger.New(swagger.Config{
		PersistAuthorization: true,
	}))

	return AppServer{app: app, logger: logger, listenAddr: listenAddr}
}

// @title OnlyCloud
// @version 1.0
// @description Modern cloud service
// @BasePath /

// @securityDefinitions.apikey ApiKeyAuth
// @in cookie
// @name auth
// @description authToken
func (a *AppServer) Run() {
	a.logger.Sugar().Fatal(a.app.Listen(a.listenAddr))
}
