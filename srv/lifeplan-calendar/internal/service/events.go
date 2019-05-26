package service

import (
	"context"
	"time"
	events "github.com/evanlib/lifeplan/srv/lifeplan-calendar/proto"
	"github.com/globalsign/mgo/bson"
	rrule "github.com/teambition/rrule-go"
	"log"
)

// CreateEvent Inserts new event into data store from given request event.
func (ev *CalendarService) CreateEvent(ctx context.Context, req *events.Event, rsp *events.EventResponse) error {
	// TODO: grabs the id from context..
	// duration
	duration := req.End.Sub(req.Start)

	
	// You must be able to distinguish between the recurrence pattern end date
	// and the end date of each event instance to enable practical querying
	if req.Recurring {
		// create rule from string
		r, err := rrule.StrToRRule(req.Rrule)
		if err != nil {
			return err
		}
		dates := r.All()
		req.End = dates[len(r.All())-1].Add(duration)
	}

	event := &events.Event{
		Id:        bson.NewObjectId().Hex(),
		Title:     req.Title,
		Userid:    "1",
		Start:     req.Start,
		End:       req.End,
		Duration:  duration,
		Recurring: req.Recurring,
		Allday:    req.Allday,
		Rrule:     req.Rrule,
		Exrule:    req.Exrule,
		Exdates:   req.Exdates,
	}
	err := ev.db.Collection(CollectionEvents).Insert(event)
	if err != nil {
		return err
	}

	rsp.Event = event
	return nil
}

// UpdateEvent updates a event in data store from events request.
// There are three possibilities when it comes to updating.
// Updating all instances of an event.
// Update only future event instances while keeping past instances.
// Update only a single selected instance.
func (ev *CalendarService) UpdateEvent(ctx context.Context, req *events.EventUpdateRequest, rsp *events.EventResponse) error {
	var event *events.Event
	err := ev.db.Collection(CollectionEvents).FindId(req.Event.Id).One(&event)
	if err != nil {
		return err
	}

	// update duration(Should this be called everytime?)
	duration := req.Event.End.Sub(req.Event.Start)
	req.Event.Duration = duration

	switch req.Updatetype {
	case events.SingleInstance:
		// add exception date to recurring event
		exception := time.Date(req.Event.Start.Year(), req.Event.Start.Month(),
			req.Event.Start.Day(), event.Start.Hour(), event.Start.Minute(), event.Start.Second(), 0, time.UTC)
		event.Exdates = append(event.Exdates, exception)
		err = ev.db.Collection(CollectionEvents).UpdateId(event.Id, event)
		if err != nil {
			return err
		}
		// create new event based on requirments.
		req.Event.Recurring = false
		req.Event.Rrule = ""
		return ev.CreateEvent(ctx, req.Event, rsp)
	case events.AllInstances:
		if event.Recurring && event.Rrule != "" {
			err = ev.db.Collection(CollectionEvents).UpdateId(req.Event.Id, req.Event)
			if err != nil {
				return err
			}
			reqid := &events.FincByIdRequest{
				Id: req.Event.Id,
			}
			return ev.GetEvent(ctx, reqid, rsp)
		}
		break
	case events.FutureInstance:
		if event.Recurring && event.Rrule != "" {
			// find last occurance up to today().
			r, err := rrule.StrToRRule(event.Rrule)
			if err != nil {
				return err
			}
			rangedEvents := r.Between(event.Start, req.Event.Start, true)
			// TODO: fix this workaround...
			if len(rangedEvents) < 2 {
				req.Updatetype = events.AllInstances
				ev.UpdateEvent(ctx, req, rsp)
			}
			lastOccurnace := rangedEvents[len(rangedEvents)-2]
			r.OrigOptions.Until = lastOccurnace.Add(event.Duration)
			event.End = lastOccurnace.Add(event.Duration)
			event.Rrule = r.String()
			// set old event end to last occurance
			err = ev.db.Collection(CollectionEvents).UpdateId(event.Id, event)
			if err != nil {
				return err
			}
			// create new event with new rules
			return ev.CreateEvent(ctx, req.Event, rsp)
		}

		break
	default:
		// non recurring events
		err = ev.db.Collection(CollectionEvents).UpdateId(req.Event.Id, req.Event)
		if err != nil {
			return err
		}
		reqid := &events.FincByIdRequest{
			Id: req.Event.Id,
		}
		return ev.GetEvent(ctx, reqid, rsp)
	}

	// handle RRULE change
	// if start/end changes the client should update the RRULE?

	return nil
}

// GetEvent retrieves Event from datastore from FincByIdRequest.
func (ev *CalendarService) GetEvent(ctx context.Context, req *events.FincByIdRequest, rsp *events.EventResponse) error {
	var event *events.Event
	err := ev.db.Collection(CollectionEvents).FindId(req.Id).One(&event)
	if err != nil {
		return err
	}
	rsp.Event = event
	return nil
}

// RemoveEvent deletes an event in data store based on event request.
// There are three possibilities when it comes to deleting.
// Removing all instances of an event
// Removing only future event instances while keeping past instances.
// Removing only a single selected instance.
func (ev *CalendarService) RemoveEvent(ctx context.Context, req *events.EventUpdateRequest, rsp *events.EmptyResponse) error {
	var event *events.Event
	err := ev.db.Collection(CollectionEvents).FindId(req.Event.Id).One(&event)
	if err != nil {
		return err
	}

	switch req.Updatetype {
	case events.SingleInstance:
		// add exception date to recurring event
		exception := time.Date(req.Event.Start.Year(), req.Event.Start.Month(), req.Event.Start.Day(), event.Start.Hour(), event.Start.Minute(), event.Start.Second(), 0, time.UTC)
		event.Exdates = append(event.Exdates, exception)
		err = ev.db.Collection(CollectionEvents).UpdateId(event.Id, event)
		if err != nil {
			return err
		}
		// create new event based on requirments.
		req.Event.Recurring = false
		req.Event.Rrule = ""
		break
	case events.AllInstances:
		if event.Recurring && event.Rrule != "" {
			err = ev.db.Collection(CollectionEvents).RemoveId(event.Id)
			if err != nil {
				return err
			}
		}
		break
	case events.FutureInstance:
		if event.Recurring && event.Rrule != "" {
			// find last occurance up to today().
			r, err := rrule.StrToRRule(event.Rrule)
			if err != nil {
				return err
			}
			rangedEvents := r.Between(event.Start, req.Event.Start, true)
			// TODO: rework this algorithm
			if len(rangedEvents) < 2 {
				req.Updatetype = events.AllInstances
				ev.RemoveEvent(ctx, req, rsp)
			}
			lastOccurnace := rangedEvents[len(rangedEvents)-2]
			r.OrigOptions.Until = lastOccurnace.Add(event.Duration)
			event.End = lastOccurnace.Add(event.Duration)
			event.Rrule = r.String()
			// set old event end to last occurance
			err = ev.db.Collection(CollectionEvents).UpdateId(event.Id, event)
			if err != nil {
				return err
			}
			break
		}
		break
	default:
		err = ev.db.Collection(CollectionEvents).RemoveId(event.Id)
		if err != nil {
			return err
		}
	}

	return nil
}

// GetEventsRange retrieves events from datastore based on given start and end
// timestamps.
func (ev *CalendarService) GetEventsRange(ctx context.Context, req *events.EventRangeRequest, rsp *events.EventRangeResponse) error {
	// check if the request event is in the recurring event bounds
	var responseevents []*events.Event
	var removeEvents []int
	query := bson.M{
		"$or": []bson.M{
			// starts in range
			bson.M{"start": bson.M{"$gte": req.Start, "$lte": req.End}},
			// ends in range
			bson.M{"end": bson.M{"$gte": req.Start, "$lte": req.End}},
			// spans range
			bson.M{"start": bson.M{"$lte": req.Start}, "end": bson.M{"$gte": req.End}},
		},
		"userid": req.Userid,
	}
	log.Println(query)
	err := ev.db.Collection(CollectionEvents).Find(query).All(&responseevents)
	if err != nil {
		return err
	}

	// loop through events
	// there might be a better way of doing this?
	for i, event := range responseevents {
		if event.Recurring && event.Rrule != "" {
			set := rrule.Set{}
			r, err := rrule.StrToRRule(event.Rrule)
			if err != nil {
				return err
			}
			set.RRule(r)
			if event.Exrule != "" {
				exr, err := rrule.StrToRRule(event.Exrule)
				if err != nil {
					return err
				}
				set.ExRule(exr)
			}

			if len(event.Exdates) > 0 {
				for _, time := range event.Exdates {
					set.ExDate(time)
				}
			}
			times := set.Between(req.Start, req.End, true)
			for _, time := range times {
				newEnd := time.Add(event.Duration)
				eventcp := &events.Event{}
				*eventcp = *event
				eventcp.Start = time
				eventcp.End = newEnd
				eventcp.Recurring = false
				responseevents = append(responseevents, eventcp)
			}
			// remove the original event
			removeEvents = append(removeEvents, i)
		}
	}
	for _, j := range removeEvents {
		responseevents = responseevents[:j+copy(responseevents[j:], responseevents[j+1:])]
	}
	rsp.Events = responseevents
	return nil
}
