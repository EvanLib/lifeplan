package api

import (
	"context"
	"log"
	"time"

	"github.com/ProtocolONE/rbac"
	apirbac "github.com/evanlib/lifeplan/srv/lifeplan-api/pkg/api/apirbac"
	calendarservice "github.com/evanlib/lifeplan/srv/lifeplan-calendar/proto"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/micro/go-micro"
	"gopkg.in/go-playground/validator.v9"
)

const (
	errorQueryParamsIncorrect = "incorrect query parameters"

	requestFindByID = "id"
	requestTimeMin  = "timeMin"
	requestTimeMax  = "timeMax"
)

type ApiOptions struct {
	Enforcer *rbac.Enforcer
}

type Api struct {
	Http            *echo.Echo
	Router          *echo.Group
	enforcer        *rbac.Enforcer
	calendarservice calendarservice.CalendarService
	service         micro.Service
}

type APIValidator struct {
	validator *validator.Validate
}

func (apiVal *APIValidator) Validate(i interface{}) error {
	return apiVal.validator.Struct(i)
}

func NewServer(options *ApiOptions) (*Api, error) {
	api := &Api{
		Http:     echo.New(),
		enforcer: options.Enforcer,
	}

	// Echo Middlewares
	appContext := apirbac.NewAppContextMiddleware(api.enforcer)
	api.Router = api.Http.Group("/api/v1")
	api.Http.Use(appContext)

	api.Http.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"},
	}))
	//create new validator
	validator := validator.New()
	//register the custom validation for ISISO8601 date
	validator.RegisterValidation("ISO8601date", IsISO8601Date)
	api.Http.Validator = &APIValidator{validator: validator}

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
	eventsRoute := &EventsRouter{
		Calendarservice: api.calendarservice,
	}

	eventsRoute.InitRoutes(api.Router)
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
