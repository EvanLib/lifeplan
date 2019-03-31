package service

import (
	"context"
	"testing"
	"time"

	"github.com/evanlib/lifeplan/srv/lifeplan-calendar/internal/config"
	"github.com/evanlib/lifeplan/srv/lifeplan-calendar/internal/database"
	events "github.com/evanlib/lifeplan/srv/lifeplan-calendar/proto"
	"github.com/globalsign/mgo/bson"
	rrule "github.com/teambition/rrule-go"

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
	start := time.Date(now.Year(), now.Month(), now.Day(), 12, 0, 0, 0, time.UTC)
	end := time.Date(now.Year(), now.Month(), now.Day(), 13, 0, 0, 0, time.UTC)
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
	start := time.Date(2019, 12, 1, 12, 0, 0, 0, time.UTC)
	end := time.Date(2019, 12, 1, 13, 0, 0, 0, time.UTC)
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

func (suite *EventsTestSuite) TestExDates() {
	// start time timenow + 1 hour
	start := time.Date(2019, 12, 1, 12, 0, 0, 0, time.UTC)
	end := time.Date(2019, 12, 1, 13, 0, 0, 0, time.UTC)
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
		Byweekday: []rrule.Weekday{rrule.SA},
		Dtstart:   start,
	})

	// remove Monday, Tuesday, Wednessday DEC 2019
	set.ExDate(start.AddDate(0, 0, 1))
	set.ExDate(start.AddDate(0, 0, 2))
	set.ExDate(start.AddDate(0, 0, 3))
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
		Exdates:   set.GetExDate(),
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
	assert.Equal(suite.T(), 3, len(rangeRsp.Events))
}

func (suite *EventsTestSuite) TestEventUpdateTitle() {
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

	rsp = &events.EventResponse{}
	suite.service.CreateEvent(context.TODO(), req, rsp)

	// update title
	rsp.Event.Title = "Updated Event Title"
	updatereq := &events.EventUpdateRequest{
		Event: rsp.Event,
	}
	rsp = &events.EventResponse{}

	err := suite.service.UpdateEvent(context.TODO(), updatereq, rsp)
	assert.Nil(suite.T(), err)
	assert.NotNil(suite.T(), rsp.Event)
	assert.Equal(suite.T(), "Updated Event Title", rsp.Event.Title)

}

// There are three possibilities when it comes to updating recurring events.
// Updating all instances of an event
// Update only future event instances while keeping past instances
// Update only a single selected instance
func (suite *EventsTestSuite) TestUpdateEventAllInstance() {
	start := time.Date(2019, 12, 1, 12, 0, 0, 0, time.UTC)
	end := time.Date(2019, 12, 1, 13, 0, 0, 0, time.UTC)
	dur := end.Sub(start)

	// create an rrule
	r, _ := rrule.NewRRule(rrule.ROption{
		Freq:    rrule.DAILY,
		Count:   7,
		Dtstart: start,
	})

	req := &events.Event{
		Title:     "Some Event",
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

	// change the start time
	// rrule updates are handled on client side
	start = time.Date(2019, 12, 1, 13, 0, 0, 0, time.UTC)
	end = time.Date(2019, 12, 1, 14, 0, 0, 0, time.UTC)
	rsp.Event.Start = start
	rsp.Event.End = end
	r.OrigOptions.Dtstart = start
	rsp.Event.Rrule = r.String()

	// update all instance of recurring
	updatereq := &events.EventUpdateRequest{}
	updatereq.Updatetype = events.AllInstances
	updatereq.Event = rsp.Event

	err = suite.service.UpdateEvent(context.TODO(), updatereq, rsp)
	assert.Nil(suite.T(), err)
	assert.Equal(suite.T(), start, rsp.Event.Start)
	assert.Equal(suite.T(), end, rsp.Event.End)
	assert.Equal(suite.T(), r.String(), rsp.Event.Rrule)
}

func (suite *EventsTestSuite) TestUpdateEventSingleInstance() {
	start := time.Date(2019, 12, 1, 12, 0, 0, 0, time.UTC)
	end := time.Date(2019, 12, 1, 13, 0, 0, 0, time.UTC)
	dur := end.Sub(start)

	// create an rrule
	r, _ := rrule.NewRRule(rrule.ROption{
		Freq:    rrule.DAILY,
		Count:   7,
		Dtstart: start,
	})

	req := &events.Event{
		Title:     "Some Event",
		Start:     start,
		End:       end,
		Duration:  dur,
		Allday:    false,
		Recurring: true,
		Rrule:     r.String(),
	}

	rsp := &events.EventResponse{}
	err := suite.service.CreateEvent(context.TODO(), req, rsp)
	oldID := rsp.Event.Id
	assert.Nil(suite.T(), err)

	// change the start time
	// rrule updates are handled on client side
	start = time.Date(2019, 12, 1, 13, 0, 0, 0, time.UTC)
	end = time.Date(2019, 12, 1, 14, 0, 0, 0, time.UTC)
	rsp.Event.Start = start
	rsp.Event.End = end

	// update single instance of recurring
	updatereq := &events.EventUpdateRequest{}
	updatereq.Updatetype = events.SingleInstance
	updatereq.Event = rsp.Event

	err = suite.service.UpdateEvent(context.TODO(), updatereq, rsp)
	assert.Nil(suite.T(), err)
	assert.Equal(suite.T(), start, rsp.Event.Start)
	assert.Equal(suite.T(), end, rsp.Event.End)
	assert.NotEqual(suite.T(), oldID, rsp.Event.Id)

	// change the start time
	// rrule updates are handled on client side
	start = time.Date(2019, 12, 7, 13, 0, 0, 0, time.UTC)
	end = time.Date(2019, 12, 7, 14, 0, 0, 0, time.UTC)
	rsp.Event.Start = start
	rsp.Event.End = end
	rsp.Event.Id = oldID

	// update single instance of recurring
	updatereq = &events.EventUpdateRequest{}
	updatereq.Updatetype = events.SingleInstance
	updatereq.Event = rsp.Event

	err = suite.service.UpdateEvent(context.TODO(), updatereq, rsp)
	assert.Nil(suite.T(), err)
	assert.Equal(suite.T(), start, rsp.Event.Start)
	assert.Equal(suite.T(), end, rsp.Event.End)
	assert.NotEqual(suite.T(), oldID, rsp.Event.Id)

	// get range
	startRange := time.Date(2019, 12, 1, 12, 0, 0, 0, time.UTC)
	endRange := startRange.AddDate(0, 0, 7)
	rangeReq := &events.EventRangeRequest{
		Start: startRange,
		End:   endRange,
	}
	rangeRsp := &events.EventRangeResponse{}

	err = suite.service.GetEventsRange(context.TODO(), rangeReq, rangeRsp)
	// for _, time := range rangeRsp.Events {
	// 	fmt.Println(time.Start)
	// }
	assert.Nil(suite.T(), err)
	assert.Equal(suite.T(), 7, len(rangeRsp.Events))

}

func (suite *EventsTestSuite) TestUpdateEventFutureInstance() {
	start := time.Date(2019, 12, 1, 12, 0, 0, 0, time.UTC)
	end := time.Date(2019, 12, 1, 13, 0, 0, 0, time.UTC)
	dur := end.Sub(start)

	// create an rrule
	r, _ := rrule.NewRRule(rrule.ROption{
		Freq:    rrule.DAILY,
		Until:   start.AddDate(0, 0, 7),
		Dtstart: start,
	})

	req := &events.Event{
		Title:     "Some Event",
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

	// change the start time
	// rrule updates are handled on client side
	start = time.Date(2019, 12, 2, 13, 0, 0, 0, time.UTC)
	end = time.Date(2019, 12, 2, 14, 0, 0, 0, time.UTC)
	rsp.Event.Start = start
	rsp.Event.End = end
	r.OrigOptions.Dtstart = start
	rsp.Event.Rrule = r.String()

	// update all instance of recurring
	updatereq := &events.EventUpdateRequest{}
	updatereq.Updatetype = events.FutureInstance
	updatereq.Event = rsp.Event

	err = suite.service.UpdateEvent(context.TODO(), updatereq, rsp)
	assert.Nil(suite.T(), err)
	assert.Equal(suite.T(), start, rsp.Event.Start)
	assert.Equal(suite.T(), r.String(), rsp.Event.Rrule)

	// get range
	startRange := time.Date(2019, 12, 1, 0, 0, 0, 0, time.UTC)
	endRange := startRange.AddDate(0, 0, 7)
	rangeReq := &events.EventRangeRequest{
		Start: startRange,
		End:   endRange,
	}
	rangeRsp := &events.EventRangeResponse{}

	err = suite.service.GetEventsRange(context.TODO(), rangeReq, rangeRsp)
	// for i, time := range rangeRsp.Events {
	// 	fmt.Println(i, time.Start)
	// }
	assert.Nil(suite.T(), err)
	assert.Equal(suite.T(), 7, len(rangeRsp.Events))

}
