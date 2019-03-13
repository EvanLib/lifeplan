package internal

import (
	"log"

	"github.com/evanlib/lifeplan/srv/users/internal/config"
	"github.com/evanlib/lifeplan/srv/users/internal/database"
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

	// micro options
	options := []micro.Option{
		micro.Name("Users"),
		micro.Version("0.1"),
	}

	// app service
	app.service = micro.NewService(options...)

	app.service.Init()

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
		log.Fatal("[USERS] Databased fail", err)
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
