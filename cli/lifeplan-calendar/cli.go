package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"time"

	pb "github.com/evanlib/lifeplan/srv/lifeplan-calendar/proto"
	microclient "github.com/micro/go-micro/client"
	"github.com/micro/go-micro/cmd"
	rrule "github.com/teambition/rrule-go"
)

func createWeekOfEvents() []*pb.Event {
	var events []*pb.Event
	timenow := time.Now().AddDate(0, 0, 5)

	for i := 0; i <= 7; i++ {
		name := fmt.Sprintf("Event %v", i)
		start := timenow.Add(time.Duration(i) * (time.Hour * 24))
		event := &pb.Event{
			Title: name,
			Start: start,
			End:   start.Add(time.Hour),
		}
		events = append(events, event)
	}

	return events
}

func main() {

	cmd.Init()
	// create new user client
	client := pb.NewCalendarService("go.micro.src.lifeplan-calendar", microclient.DefaultClient)

	calendarname := "Evan's Calendar"

	rsp, err := client.CreateCalendar(context.TODO(), &pb.Calendar{
		Name: calendarname,
	})
	if err != nil {
		log.Printf("Could not create calendar: %v", err)
	}
	log.Printf("Created Calendar: %s", rsp.Calendar.Name)

	// get calendar by id
	getCal := &pb.FincByIdRequest{
		Id: rsp.Calendar.Id,
	}
	rsp, err = client.GetCalendar(context.TODO(), getCal)
	if err != nil {
		log.Printf("Could not get calendar %v", err)
	}
	log.Printf("Calendar GET: %s:", rsp.Calendar.Name)

	// update calendar
	updateCal := &pb.Calendar{
		Id:   rsp.Calendar.Id,
		Name: "Updated Name",
	}
	rsp, err = client.UpdateCalendar(context.TODO(), updateCal)
	if err != nil {
		log.Printf("Could not update calendar: %v", err)
	}
	log.Printf("Updated calendar %s: %s", rsp.Calendar.Id, rsp.Calendar.Name)

	// remove calendar
	_, err = client.RemoveCalendar(context.TODO(), getCal)
	if err != nil {
		log.Printf("Could not remove calendar %v", err)
	}
	log.Printf("Removed calendar: %s", getCal.Id)

	// events creation
	start := time.Now().AddDate(0, 0, 2)
	end := start.Add(time.Hour)
	dur := end.Sub(start)
	fmt.Println(dur)
	event, err := client.CreateEvent(context.TODO(), &pb.Event{
		Title:    "Clean the car. This one",
		Start:    start,
		End:      end,
		Duration: dur,
	})

	if err != nil {
		log.Printf("Could not create Event. %v", err)
	}
	log.Printf("Created event: %s", event.Event.Title)

	eventrsp, err := client.GetEvent(context.TODO(), &pb.FincByIdRequest{Id: event.Event.Id})
	if err != nil {
		log.Printf("Could not get Event. %v", err)
	}
	log.Printf("Event GET %s", event.Event.Title)
	log.Printf("Event start %s, Event end %s", eventrsp.Event.Start, eventrsp.Event.End)

	// recurrance events
	r, _ := rrule.NewRRule(rrule.ROption{
		Freq:    rrule.DAILY,
		Count:   10,
		Dtstart: time.Now(),
	})
	event, err = client.CreateEvent(context.TODO(), &pb.Event{
		Title:     "Some habit",
		Start:     start,
		End:       end,
		Duration:  dur,
		Recurring: true,
		Rrule:     r.String(),
	})

	log.Print(r.String())
	eventrsp, err = client.GetEvent(context.TODO(), &pb.FincByIdRequest{Id: event.Event.Id})
	if err != nil {
		log.Printf("Could not get Event. %v", err)
	}
	log.Printf("Retrieved recurrence string %s", eventrsp.Event.Rrule)

	// tomorrow events.
	b := time.Now().AddDate(0, 0, 2)
	e := time.Now().AddDate(0, 0, 4)
	d := 24 * time.Hour
	log.Printf("Retrieving date range start: %s and end %s", b, e)
	eventsrsp, err := client.GetEventsRange(context.TODO(), &pb.EventRangeRequest{
		Start: b.Truncate(d),
		End:   e.Truncate(d),
	})
	if err != nil {
		log.Printf("Error getting tomorrow events: %v", err)
	}
	for _, e := range eventsrsp.Events {
		log.Printf("Title: %s, RR: %v  Stime: %v Etime: %v\n", e.Title, e.Recurring, e.Start, e.End)
	}

	r, err = rrule.StrToRRule(eventrsp.Event.Rrule)
	if err != nil {
		log.Printf("Error in updating event %v", err)
	}

	// Test title
	eventrsp.Event.Title = "UPDATE TITLE :D"
	eventUpdate, err := client.UpdateEvent(context.TODO(), eventrsp.Event)
	if err != nil {
		log.Printf("Error on updating event %v", err)
	}
	log.Printf("Updated recurrence event %s", eventUpdate.Event.Title)

	// Test rrule stuff
	
	os.Exit(0)
}
