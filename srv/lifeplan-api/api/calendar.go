package api

import (
	"bytes"
	"context"
	"io/ioutil"
	"net/http"

	calendar "github.com/evanlib/lifeplan/srv/lifeplan-calendar/proto"
	"github.com/labstack/echo/v4"
)

type EventJsonBinder struct{}
type EventRangeJsonBinder struct{}
type EventUpdateJsonBinder struct{}

type calendarRoute struct {
	*Api
}

func (api *Api) InitCalendarRoutes() *Api {
	route := calendarRoute{Api: api}
	api.Http.POST("/api/v1/event", route.createEvent)
	api.Http.POST("/api/v1/event/:id", route.updateEvent)
	api.Http.GET("/api/v1/events/getByRange", route.getEventByRange)
	api.Http.GET("api/v1/events", route.getEvents)
	api.Http.GET("/api/v1/event/:id", route.getEvent)
	api.Http.DELETE("api/v1/event/:id", route.deleteEvent)
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

func (cb *EventRangeJsonBinder) Bind(i interface{}, ctx echo.Context) (err error) {
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

func (cb *EventUpdateJsonBinder) Bind(i interface{}, ctx echo.Context) (err error) {
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
		return echo.NewHTTPError(http.StatusBadRequest, errorQueryParamsIncorrect)
	}

	rsp, err := r.Api.calendarservice.CreateEvent(context.TODO(), req)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	return ctx.JSON(http.StatusOK, rsp)
}

func (r *calendarRoute) getEvent(ctx echo.Context) error {
	id := ctx.Param(requestFindByID)
	if id == "" {
		return echo.NewHTTPError(http.StatusBadRequest, errorQueryParamsIncorrect)
	}

	rsp, err := r.calendarservice.GetEvent(context.TODO(), &calendar.FincByIdRequest{Id: id})
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	return ctx.JSON(http.StatusOK, rsp)
}
func (r *calendarRoute) getEventByRange(ctx echo.Context) error {
	req := &calendar.EventRangeRequest{}
	err := (&EventRangeJsonBinder{}).Bind(req, ctx)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, errorQueryParamsIncorrect)
	}
	//user id from cookie
	req.Userid = "1"

	rsp, err := r.calendarservice.GetEventsRange(context.TODO(), req)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	return ctx.JSON(http.StatusOK, rsp)
}

func (r *calendarRoute) updateEvent(ctx echo.Context) error {
	// double id checks....
	id := ctx.Param(requestFindByID)
	if id == "" {
		return echo.NewHTTPError(http.StatusBadRequest, errorQueryParamsIncorrect)
	}

	req := &calendar.EventUpdateRequest{}
	err := (&EventUpdateJsonBinder{}).Bind(req, ctx)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, errorQueryParamsIncorrect)
	}

	rsp, err := r.calendarservice.UpdateEvent(context.TODO(), req)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	return ctx.JSON(http.StatusOK, rsp)
}

func (r *calendarRoute) deleteEvent(ctx echo.Context) error {
	// double id checks....
	id := ctx.Param(requestFindByID)
	if id == "" {
		return echo.NewHTTPError(http.StatusBadRequest, errorQueryParamsIncorrect)
	}

	req := &calendar.EventUpdateRequest{}
	err := (&EventUpdateJsonBinder{}).Bind(req, ctx)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, errorQueryParamsIncorrect)
	}

	rsp, err := r.calendarservice.RemoveEvent(context.TODO(), req)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	return ctx.JSON(http.StatusOK, rsp)
}

func (r *calendarRoute) getEvents(ctx echo.Context) error {

	return ctx.JSON(http.StatusOK, rsp)
}
