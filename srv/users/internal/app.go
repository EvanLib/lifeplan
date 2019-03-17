package internal

import (
	"log"

	"github.com/evanlib/lifeplan/srv/users/internal/config"
	"github.com/evanlib/lifeplan/srv/users/internal/database"
	"github.com/evanlib/lifeplan/srv/users/internal/service"
	grpc "github.com/evanlib/lifeplan/srv/users/proto"
	micro "github.com/micro/go-micro"
)

type Application struct {
	cfg      *config.Config
	service  micro.Service
	database *database.Source
}

func NewAppliction() *Application {
	return &Application{}
}

func (app *Application) Init() {

	// config
	cfg, err := config.NewConfig()
	if err != nil {
		log.Fatal("[USERS] Config load failed", err)
	}
	app.cfg = cfg

	// database
	app.initDatabase()
	// micro options
	options := []micro.Option{
		micro.Name("go.micro.src.users"),
		micro.Version("0.1"),
	}

	// app service
	app.service = micro.NewService(options...)
	app.service.Init()

	// users service
	svc := service.NewUsersService(app.database)

	err = grpc.RegisterUsersHandler(app.service.Server(), svc)

	if err != nil {
		log.Fatal(err)
	}
}

func (app *Application) initDatabase() {
	settings := database.Connection{
		Host:     app.cfg.MongoHost,
		Database: app.cfg.MongoDatabase,
		User:     app.cfg.MongoUser,
		Password: app.cfg.MongoPassword,
	}

	db, err := database.NewDatabase(settings)

	if err != nil {
		log.Fatal("[USERS] Databased fail ", err)
	}

	app.database = db
}

func (app *Application) Run() {
	if err := app.service.Run(); err != nil {
		log.Fatal(err)
	}
}

func (app *Application) Stop() {
	log.Println("[USERS] Stopped")
}
