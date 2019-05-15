package main

import (
	"encoding/json"
	"log"
	"strings"

	calendar "github.com/evanlib/lifeplan/srv/lifeplan-calendar/proto"
	"github.com/micro/go-micro"
	"github.com/micro/go-micro/errors"
	api "github.com/micro/micro/api/proto"

	"context"
)

type CalendarAPI struct {
	Client calendar.CalendarService
}

func (cal *CalendarAPI) CreateEvent(ctx context.Context, req *api.Request, rsp *api.Response) {
	log.Print("Received Calendar.CreateEvent API request")
	return nil
}

func main() {
	service := micro.NewService(
		micro.Name("go.micro.api.calendar"),
	)

	// // parse command line flags
	// service.Server().Handle{

	// }
}