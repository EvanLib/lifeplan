package service

import (
	"context"
	"testing"
	"time"

	"github.com/evanlib/lifeplan/srv/lifeplan-calendar/internal/config"
	"github.com/evanlib/lifeplan/srv/lifeplan-calendar/internal/database"
	events "github.com/evanlib/lifeplan/srv/lifeplan-calendar/proto"
	rrule "github.com/teambition/rrule-go"

	"github.com/globalsign/mgo/bson"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

type EventsTestSuite struct {
	suite.Suite
	service *CalendarService
}

func Test_Events(t *testing.T) {
	suite.Run(t, new(EventsTestSuite))
}

func (suite *EventsTestSuite) SetupTest() {
	cfg, err := config.NewConfig()
	if err != nil {
		suite.FailNow("Config load failed", "%v", err)
	}

	settings := database.Connection{
		Host:     cfg.MongoHost,
		Database: cfg.MongoDatabase,
		User:     cfg.MongoUser,
		Password: cfg.MongoPassword,
	}

	db, err := database.NewDatabase(settings)
	if err != nil {
		suite.FailNow("Database connection failed", "%v", err)
	}

	suite.service = NewCalendarService(db)

}

func (suite *EventsTestSuite) TearDownTest() {
	if err := suite.service.db.Drop(); err != nil {
		suite.FailNow("Database deletion failed", "%v", err)
	}

	suite.service.db.Close()
}

func (suite *EventsTestSuite) TestEventCreate() {
	// start time timenow + 1 hour
	start := time.Now().Add(time.Hour)
	end := time.Now().Add(time.Hour * 2)
	dur := end.Sub(start)

	rsp := &events.EventResponse{}

	// create testing request
	req := &events.Event{
		Id:        bson.NewObjectId().Hex(),
		Title:     "Test Event Title",
		Userid:    "1",
		Start:     start,
		End:       end,
		Duration:  dur,
		Recurring: false,
		Allday:    false,
	}
	err := suite.service.CreateEvent(context.TODO(), req, rsp)

	assert.Nil(suite.T(), err)
	assert.NotNil(suite.T(), rsp.Event)
	assert.Equal(suite.T(), req.Start, rsp.Event.Start)
}

func (suite *EventsTestSuite) TestRangeEvents() {
	// start time timenow + 1 hour
	now := time.Now()
	start := time.Date(now.Year(), now.Month(), now.Day(), 12, 0, 0, 0, now.Location())
	end := time.Date(now.Year(), now.Month(), now.Day(), 13, 0, 0, 0, now.Location())
	dur := end.Sub(start)

	// create an rrule
	r, _ := rrule.NewRRule(rrule.ROption{
		Freq:    rrule.DAILY,
		Count:   10,
		Dtstart: time.Now(),
	})

	req := &events.Event{
		Title:     "Some habit",
		Start:     start,
		End:       end,
		Duration:  dur,
		Allday:    false,
		Recurring: true,
		Rrule:     r.String(),
	}

	rsp := &events.EventResponse{}
	err := suite.service.CreateEvent(context.TODO(), req, rsp)
	assert.Nil(suite.T(), err)
	assert.NotNil(suite.T(), rsp.Event)
	assert.Equal(suite.T(), req.Start, rsp.Event.Start)

	// get range
	startRange := time.Date(now.Year(), now.Month(), now.Day(), 0, 0, 0, 0, now.Location())
	endRange := startRange.AddDate(0, 0, 10)

	rangeReq := &events.EventRangeRequest{
		Start: startRange,
		End:   endRange,
	}
	rangeRsp := &events.EventRangeResponse{}

	err = suite.service.GetEventsRange(context.TODO(), rangeReq, rangeRsp)
	assert.Nil(suite.T(), err)
	assert.Equal(suite.T(), 10, len(rangeRsp.Events))
}

func (suite *EventsTestSuite) TestExRrule() {
	// start time timenow + 1 hour
	now := time.Now()
	start := time.Date(2019, 12, 1, 12, 0, 0, 0, now.Location())
	end := time.Date(2019, 12, 1, 13, 0, 0, 0, now.Location())
	dur := end.Sub(start)

	// create an rrule
	set := rrule.Set{}
	r, _ := rrule.NewRRule(rrule.ROption{
		Freq:    rrule.DAILY,
		Count:   7,
		Dtstart: start,
	})
	set.RRule(r)

	exr, _ := rrule.NewRRule(rrule.ROption{
		Freq:      rrule.YEARLY,
		Byweekday: []rrule.Weekday{rrule.SA, rrule.SU},
		Dtstart:   start,
	})
	set.ExRule(exr)

	req := &events.Event{
		Title:     "Some habit",
		Start:     start,
		End:       end,
		Duration:  dur,
		Allday:    false,
		Recurring: true,
		Rrule:     r.String(),
		Exrule:    exr.String(),
	}

	rsp := &events.EventResponse{}
	err := suite.service.CreateEvent(context.TODO(), req, rsp)
	assert.Nil(suite.T(), err)
	assert.NotNil(suite.T(), rsp.Event)
	assert.Equal(suite.T(), req.Start, rsp.Event.Start)

	// get range
	endRange := start.AddDate(0, 0, 7)
	rangeReq := &events.EventRangeRequest{
		Start: start,
		End:   endRange,
	}
	rangeRsp := &events.EventRangeResponse{}

	err = suite.service.GetEventsRange(context.TODO(), rangeReq, rangeRsp)
	assert.Nil(suite.T(), err)
	assert.Equal(suite.T(), 5, len(rangeRsp.Events))
}
