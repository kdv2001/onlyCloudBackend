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

type postgresData struct {
	UserName string
	Password string
	Address  string
	Port     string
	DbName   string
}

func (p postgresData) configurePostgresAddress() string {
	return fmt.Sprintf("postgres://%s:%s@%s/%s",
		p.UserName, p.Password, net.JoinHostPort(p.Address, p.Port), p.DbName)
}

func NewServer(logger *zap.Logger, listenAddr string) AppServer {
	// postgres
	p := postgresData{}

	err := viper.UnmarshalKey("postgresql", &p)
	if err != nil {
		logger.Sugar().Fatal(err)
	}

	uri := p.configurePostgresAddress()

	conn, err := pgx.Connect(context.Background(), uri)
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

// Run
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
