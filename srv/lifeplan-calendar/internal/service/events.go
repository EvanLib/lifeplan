package service

import (
	"context"
	"log"

	events "github.com/evanlib/lifeplan/srv/lifeplan-calendar/proto"

	"gopkg.in/mgo.v2/bson"
)

func (ev *CalendarService) CreateEvent(ctx context.Context, req *events.Event, rsp *events.EventResponse) error {
	// TODO: grabs the id from context..
	event := &events.Event{
		Id:     bson.NewObjectId().Hex(),
		Name:   req.Name,
		Userid: "1",
		Start:  req.Start,
		End:    req.End,
	}
	err := ev.db.Collection(CollectionEvents).Insert(event)
	if err != nil {
		return err
	}

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
	// var events []*events.Event
	// start, err := timestamp.Timestamp(req.Start)
	// if err != nil {
	// 	return err
	// }

	// end, err := timestamp.Timestamp(req.End)
	// if err != nil {
	// 	return err
	// }

	return nil
}
