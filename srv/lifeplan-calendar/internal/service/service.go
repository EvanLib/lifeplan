package service

import (
	"context"
	"log"

	"github.com/evanlib/lifeplan/srv/lifeplan-calendar/internal/database"
	calendar "github.com/evanlib/lifeplan/srv/lifeplan-calendar/proto"
	"gopkg.in/mgo.v2/bson"
)

const (
	CollectionCalendar = "Lifeplan-Calendars"
	CollectionEvents   = "Lifeplan-Events"
	calErrorNotFound   = "calendar with specified data not found"
)

type CalendarService struct {
	db *database.Source
}

// InitCalednarService returns pointer to a new calendar Service
func NewCalendarService(db *database.Source) *CalendarService {
	return &CalendarService{
		db: db,
	}
}

// CreateCalendar stores a new Calendar into database source
// rsp calendar is the newly added calendar
// returns any storage error
func (cs *CalendarService) CreateCalendar(ctx context.Context, req *calendar.Calendar, rsp *calendar.CalendarResponse) error {
	// TODO: grabs the id from context..
	cal := &calendar.Calendar{
		Id:     bson.NewObjectId().Hex(),
		Name:   req.Name,
		Userid: "1",
	}

	err := cs.db.Collection(CollectionCalendar).Insert(cal)
	if err != nil {
		return err
	}

	rsp.Calendar = cal
	return nil
}

func (cs *CalendarService) GetCalendar(ctx context.Context, req *calendar.FincByIdRequest, rsp *calendar.CalendarResponse) error {
	var cal *calendar.Calendar
	err := cs.db.Collection(CollectionCalendar).FindId(req.Id).One(&cal)
	if err != nil {
		log.Printf("Error on geting calendar %v", err)
		return err
	}
	rsp.Calendar = cal
	return nil
}

func (cs *CalendarService) RemoveCalendar(ctx context.Context, req *calendar.FincByIdRequest, rsp *calendar.EmptyResponse) error {
	//TODO: check if the userid owns the calendar first
	err := cs.db.Collection(CollectionCalendar).RemoveId(req.Id)
	if err != nil {
		return err
	}
	return nil
}

func (cs *CalendarService) UpdateCalendar(ctx context.Context, req *calendar.Calendar, rsp *calendar.CalendarResponse) error {
	err := cs.db.Collection(CollectionCalendar).UpdateId(req.Id, req)
	if err != nil {
		return err
	}

	var cal *calendar.Calendar
	err = cs.db.Collection(CollectionCalendar).FindId(req.Id).One(&cal)
	if err != nil {
		log.Printf("Error on geting calendar %v", err)
		return err
	}
	rsp.Calendar = cal
	return nil
}
