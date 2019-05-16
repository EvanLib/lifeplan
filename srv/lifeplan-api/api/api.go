package api

import (
	"context"
	"log"
	"time"

	"github.com/labstack/echo/v4"
	"github.com/micro/go-micro"

	calendarservice "github.com/evanlib/lifeplan/srv/lifeplan-calendar/proto"
)

type Api struct {
	Http            *echo.Echo
	calendarservice calendarservice.CalendarService
	service         micro.Service
}

func NewServer() (*Api, error) {
	//
	api := &Api{
		Http: echo.New(),
	}

	api.InitServices()
	return api, nil
}

func (api *Api) InitServices() {
	// create api service
	options := []micro.Option{
		micro.Name("p1payapi"),
		micro.Version("0.1"),
	}
	api.service = micro.NewService(options...)
	api.service.Init()

	// create client services
	api.calendarservice = calendarservice.NewCalendarService("go.micro.srv.calendar", api.service.Client())

	// http routing
	api.InitCalendarRoutes()
}

func (api *Api) Start() error {
	go func() {
		if err := api.service.Run(); err != nil {
			log.Fatal(err)
		}
	}()

	go func() {
		if err := api.Http.Start(":3001"); err != nil {
			log.Fatal(err)
		}
	}()
	return nil
}

func (api *Api) Stop() {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := api.Http.Shutdown(ctx); err != nil {
		log.Fatal("Server Shutdown:", err)
	}

	log.Println("Http server exiting")

	//api.serviceCancel()
	log.Println("Micro server exiting")
}
