package main

import (
	"context"
	"log"
	"os"

	pb "github.com/evanlib/lifeplan/srv/lifeplan-calendar/proto"
	microclient "github.com/micro/go-micro/client"
	"github.com/micro/go-micro/cmd"
)

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
	os.Exit(0)
}
