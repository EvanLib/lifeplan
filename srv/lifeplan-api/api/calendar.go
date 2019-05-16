package api

import (
	"bytes"
	"context"
	"fmt"
	"io/ioutil"

	calendar "github.com/evanlib/lifeplan/srv/lifeplan-calendar/proto"
	"github.com/labstack/echo/v4"
)

type EventJsonBinder struct{}

type calendarRoute struct {
	*Api
}

func (api *Api) InitCalendarRoutes() *Api {
	route := calendarRoute{Api: api}
	api.Http.POST("/api/v1/event", route.createEvent)
	return api
}

func (cb *EventJsonBinder) Bind(i interface{}, ctx echo.Context) (err error) {
	var buf []byte

	if ctx.Request().Body != nil {
		buf, err = ioutil.ReadAll(ctx.Request().Body)
		rdr := ioutil.NopCloser(bytes.NewBuffer(buf))

		if err != nil {
			return err
		}

		ctx.Request().Body = rdr
	}

	db := new(echo.DefaultBinder)
	if err = db.Bind(i, ctx); err != nil {
		return err
	}

	return
}

func (r *calendarRoute) createEvent(ctx echo.Context) error {
	req := &calendar.Event{}
	err := (&EventJsonBinder{}).Bind(req, ctx)
	if err != nil {
		return err
	}

	fmt.Println(req.Title)

	rsp, err := r.Api.calendarservice.CreateEvent(context.TODO(), req)
	if err != nil {
		return err
	}

	fmt.Println(rsp.Event.Title)
	return nil
}
