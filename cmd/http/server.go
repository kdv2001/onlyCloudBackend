package main

import (
	"context"
	"errors"
	"fmt"
	"net"
	"reflect"

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

func (p postgresData) configurePostgresAddress() (string, error) {
	v := reflect.ValueOf(p)
	for i := 0; i < v.NumField(); i++ {
		if v.Field(i).Interface() == "" {
			return "", errors.New(fmt.Sprintf("postgresData field %s is empty", v.Type().Field(i).Name))
		}
	}

	return fmt.Sprintf("postgres://%s:%s@%s/%s",
		p.UserName, p.Password, net.JoinHostPort(p.Address, p.Port), p.DbName), nil
}

func NewServer(logger *zap.Logger, listenAddr string) AppServer {
	// postgres
	p := postgresData{}

	err := viper.UnmarshalKey("postgresql", &p)
	if err != nil {
		logger.Sugar().Fatal(err)
	}

	uri, err := p.configurePostgresAddress()
	if err != nil {
		logger.Sugar().Fatal(err)
	}

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
