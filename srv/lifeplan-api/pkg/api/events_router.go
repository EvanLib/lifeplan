package api

import (
	"bytes"
	"context"
	"fmt"
	"log"
	"net/http"

	"io/ioutil"

	"github.com/evanlib/lifeplan/srv/lifeplan-api/pkg/api/apirbac"
	calendar "github.com/evanlib/lifeplan/srv/lifeplan-calendar/proto"
	calendarservice "github.com/evanlib/lifeplan/srv/lifeplan-calendar/proto"
	"github.com/labstack/echo/v4"
)

type EventJsonBinder struct{}
type EventRangeJsonBinder struct{}
type EventUpdateJsonBinder struct{}

type EventsRouter struct {
	Calendarservice calendarservice.CalendarService
}

// TODO: FIX THIS BAD PRACTICE
func (r *EventsRouter) GetOwner(ctx apirbac.AppContext) (string, error) {
	owner := ctx.QueryParam("ownerTesting")
	return owner, nil
}

// THIS BAD PRACTICE !!!!!
// JWT AUTH IS NOT IMPLEMENTEED YET
func GetUserID() string {
	return "1"
}

func (cb *EventJsonBinder) Bind(i interface{}, ctx echo.Context) (err error) {
	var buf []byte

	if ctx.Request().Body != nil {
		buf, err = ioutil.ReadAll(ctx.Request().Body)
		rdr := ioutil.NopCloser(bytes.NewBuffer(buf))

		if err != nil {
			fmt.Println(err)
			return err
		}

		ctx.Request().Body = rdr
	}

	db := new(echo.DefaultBinder)
	if err = db.Bind(i, ctx); err != nil {
		fmt.Println(err)

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

func (r *EventsRouter) InitRoutes(route *echo.Group) {
	//Calendar Routing
	er := apirbac.Group(route, "/calendar", r, []string{"*", CalendarType, CalendarDomain})
	// er.GET("/:calendarId", r.Get, nil)
	er.GET("/events", r.GetEvents, nil)
	er.POST("/events", r.CreateEvent, nil)

	eventsGroup := apirbac.Group(route, "/events", r, []string{"eventID", EventType, CalendarDomain})
	eventsGroup.GET("/:eventID", r.Get, nil)
	eventsGroup.PUT("/:eventID", r.UpdateEvent, nil)
	eventsGroup.POST("/:eventID", r.DeleteEvent, nil)
}

func (r *EventsRouter) Get(ctx echo.Context) error {
	id := ctx.Param("eventID")
	fmt.Println(id)
	if id == "" {
		return echo.NewHTTPError(http.StatusBadRequest, errorQueryParamsIncorrect)
	}

	req := &calendar.FincByIdRequest{
		Id: id,
	}

	if err := ctx.Validate(req); err != nil {
		return echo.NewHTTPError(http.StatusUnprocessableEntity, err.Error())
	}

	rsp, err := r.Calendarservice.GetEvent(context.TODO(), &calendar.FincByIdRequest{Id: id})
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	return ctx.JSON(http.StatusOK, rsp.Event)
}

func (r *EventsRouter) GetEvents(ctx echo.Context) error {

	rsp, err := r.Calendarservice.GetEventsByUserID(context.TODO(), &calendar.FincByIdRequest{
		Id: GetUserID(),
	})
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	return ctx.JSON(http.StatusOK, rsp.Events)
}

func (r *EventsRouter) CreateEvent(ctx echo.Context) error {
	req := &calendar.Event{}

	err := (&EventJsonBinder{}).Bind(req, ctx)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, errorQueryParamsIncorrect)
	}
	req.Userid = GetUserID()
	if err := ctx.Validate(req); err != nil {
		return echo.NewHTTPError(http.StatusUnprocessableEntity, err.Error())
	}

	rsp, err := r.Calendarservice.CreateEvent(context.TODO(), req)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	return ctx.JSON(http.StatusOK, rsp.Event)
}

func (r *EventsRouter) UpdateEvent(ctx echo.Context) error {
	// double id checks....
	id := ctx.Param("eventID")
	if id == "" {
		return echo.NewHTTPError(http.StatusBadRequest, errorQueryParamsIncorrect)
	}

	req := &calendar.EventUpdateRequest{}
	err := (&EventUpdateJsonBinder{}).Bind(req, ctx)
	if err != nil {
		log.Println(err)
		return echo.NewHTTPError(http.StatusBadRequest, errorQueryParamsIncorrect)
	}

	if err := ctx.Validate(req); err != nil {
		return echo.NewHTTPError(http.StatusUnprocessableEntity, err.Error())
	}

	rsp, err := r.Calendarservice.UpdateEvent(context.TODO(), req)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	return ctx.JSON(http.StatusOK, rsp.Event)
}

func (r *EventsRouter) DeleteEvent(ctx echo.Context) error {
	// double id checks....
	id := ctx.Param("eventID")
	if id == "" {
		return echo.NewHTTPError(http.StatusBadRequest, errorQueryParamsIncorrect)
	}

	req := &calendar.EventUpdateRequest{}
	err := (&EventUpdateJsonBinder{}).Bind(req, ctx)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, errorQueryParamsIncorrect)
	}

	if err := ctx.Validate(req); err != nil {
		return echo.NewHTTPError(http.StatusUnprocessableEntity, err.Error())
	}

	rsp, err := r.Calendarservice.RemoveEvent(context.TODO(), req)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	return ctx.JSON(http.StatusOK, rsp)
}
