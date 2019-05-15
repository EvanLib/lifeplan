package main

import (
	"log"
	"encoding/json"

	calendar "github.com/evanlib/lifeplan/srv/lifeplan-calendar/proto"
	"github.com/micro/go-micro"
	api "github.com/micro/micro/api/proto"

	"context"
)

type CalendarAPI struct {
	Client calendar.CalendarService
}

func (cal *CalendarAPI) CreateEvent(ctx context.Context, req *api.Request, rsp *api.Response) error {
	log.Print("Received Calendar.CreateEvent API request")
	rsp.StatusCode = 200
	b, _ := json.Marshal(map[string]string{
		"message": "AYYYEEE",
	})
	rsp.Body = string(b)
	return nil
}

func main() {
	service := micro.NewService(
		micro.Name("go.micro.api.calendar"),
	)

	service.Init()
	service.Server().Handle(
		service.Server().NewHandler(
			&CalendarAPI{
				Client: calendar.NewCalendarService("go.micro.src.lifeplan-calendar", service.Client())},
		),
	)
}
