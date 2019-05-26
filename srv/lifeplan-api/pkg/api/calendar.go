package api

import (
	"context"
	"net/http"
	"time"

	calendar "github.com/evanlib/lifeplan/srv/lifeplan-calendar/proto"
	"github.com/labstack/echo/v4"
)

const (
	CalendarType   = "calendars"
	CalendarDomain = "calendar"

	EventType = "event"
)

type calendarRoute struct {
	*Api
}

func (api *Api) InitCalendarRoutes() *Api {
	// route := calendarRoute{Api: api}
	// api.Http.POST("/api/v1/event", route.createEvent)
	// api.Http.POST("/api/v1/event/:id", route.updateEvent)
	// api.Http.GET("/api/v1/events/getByRange", route.getEventByRange)
	// api.Http.GET("api/v1/events", route.getEvents)
	// api.Http.GET("/api/v1/event/:id", route.getEvent)
	// api.Http.DELETE("api/v1/event/:id", route.deleteEvent)
	return api
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

	req := &calendar.EventRangeRequest{
		Userid: "1",
		Start:  time.Time{},
		End:    time.Date(9999, 1, 1, 0, 0, 0, 0, time.UTC),
	}

	rsp, err := r.calendarservice.GetEventsRange(context.TODO(), req)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	return ctx.JSON(http.StatusOK, rsp)
}
