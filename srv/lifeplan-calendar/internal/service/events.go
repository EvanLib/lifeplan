package service

import (
	"context"
	"log"

	events "github.com/evanlib/lifeplan/srv/lifeplan-calendar/proto"
	"github.com/globalsign/mgo/bson"
	rrule "github.com/teambition/rrule-go"
)

func (ev *CalendarService) CreateEvent(ctx context.Context, req *events.Event, rsp *events.EventResponse) error {
	// TODO: grabs the id from context..

	if req.Recurring {
		// create rule from string
		r, err := rrule.StrToRRule(req.Rrule)
		if err != nil {
			return err
		}
		dates := r.All()
		req.End = dates[len(r.All())-1]
	}

	event := &events.Event{
		Id:        bson.NewObjectId().Hex(),
		Title:     req.Title,
		Userid:    "1",
		Start:     req.Start,
		End:       req.End,
		Duration:  req.Duration,
		Recurring: req.Recurring,
		Allday:    req.Allday,
		Rrule:     req.Rrule,
	}
	err := ev.db.Collection(CollectionEvents).Insert(event)
	if err != nil {
		return err
	}

	// rrule https://tools.ietf.org/html/rfc5545#section-3.8.5.3

	rsp.Event = event
	return nil
}

func (ev *CalendarService) UpdateEvent(ctx context.Context, req *events.Event, rsp *events.EventResponse) error {
	err := ev.db.Collection(CollectionEvents).UpdateId(req.Id, req)
	if err != nil {
		return err
	}

	var event *events.Event
	err = ev.db.Collection(CollectionEvents).FindId(req.Id).One(&event)
	if err != nil {
		log.Printf("Error on geting Event %v", err)
		return err
	}
	rsp.Event = event
	return nil
}

func (ev *CalendarService) GetEvent(ctx context.Context, req *events.FincByIdRequest, rsp *events.EventResponse) error {
	var event *events.Event
	err := ev.db.Collection(CollectionEvents).FindId(req.Id).One(&event)
	if err != nil {
		log.Printf("Error on geting Event %v", err)
		return err
	}
	rsp.Event = event
	return nil
}

func (ev *CalendarService) RemoveEvent(ctx context.Context, req *events.FincByIdRequest, rsp *events.EmptyResponse) error {
	//TODO: check if the userid owns the calendar first
	err := ev.db.Collection(CollectionEvents).RemoveId(req.Id)
	if err != nil {
		return err
	}
	return nil
}

func (ev *CalendarService) GetEventsRange(ctx context.Context, req *events.EventRangeRequest, rsp *events.EventRangeResponse) error {

	// check if the request event is in the recurring event bounds
	var events []*events.Event
	query := bson.M{
		"$or": []bson.M{
			// starts in range
			bson.M{"start": bson.M{"$gte": req.Start, "$lte": req.End}},
			// ends in range
			bson.M{"end": bson.M{"$gte": req.Start, "$lte": req.End}},
			// spans range
			bson.M{"start": bson.M{"$lte": req.Start}, "end": bson.M{"$gte": req.End}},
		},
	}
	err := ev.db.Collection(CollectionEvents).Find(query).All(&events)
	if err != nil {
		return err
	}

	rsp.Events = events
	return nil
}
