// Code generated by protoc-gen-micro. DO NOT EDIT.
// source: proto/calendar.proto

/*
Package calendar is a generated protocol buffer package.

It is generated from these files:
	proto/calendar.proto

It has these top-level messages:
	EmptyResponse
	Calendar
	CalendarResponse
	FincByIdRequest
	EventRangeRequest
	EventUpdateRequest
	EventRangeResponse
	EventResponse
	Event
	Task
	TreeNode
	TasksTree
	TasksRequest
	TaskList
	TaskResponse
*/
package calendar

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import _ "github.com/golang/protobuf/ptypes/timestamp"
import _ "github.com/golang/protobuf/ptypes/duration"
import _ "github.com/gogo/protobuf/gogoproto"

import (
	context "context"
	client "github.com/micro/go-micro/client"
	server "github.com/micro/go-micro/server"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ client.Option
var _ server.Option

// Client API for CalendarService service

type CalendarService interface {
	CreateCalendar(ctx context.Context, in *Calendar, opts ...client.CallOption) (*CalendarResponse, error)
	GetCalendar(ctx context.Context, in *FincByIdRequest, opts ...client.CallOption) (*CalendarResponse, error)
	RemoveCalendar(ctx context.Context, in *FincByIdRequest, opts ...client.CallOption) (*EmptyResponse, error)
	UpdateCalendar(ctx context.Context, in *Calendar, opts ...client.CallOption) (*CalendarResponse, error)
	CreateEvent(ctx context.Context, in *Event, opts ...client.CallOption) (*EventResponse, error)
	GetEvent(ctx context.Context, in *FincByIdRequest, opts ...client.CallOption) (*EventResponse, error)
	UpdateEvent(ctx context.Context, in *EventUpdateRequest, opts ...client.CallOption) (*EventResponse, error)
	RemoveEvent(ctx context.Context, in *EventUpdateRequest, opts ...client.CallOption) (*EmptyResponse, error)
	GetEventsRange(ctx context.Context, in *EventRangeRequest, opts ...client.CallOption) (*EventRangeResponse, error)
	CreateTask(ctx context.Context, in *Task, opts ...client.CallOption) (*TaskResponse, error)
	GetTask(ctx context.Context, in *FincByIdRequest, opts ...client.CallOption) (*TaskResponse, error)
	GetTasks(ctx context.Context, in *TasksRequest, opts ...client.CallOption) (*TasksTree, error)
	UpdateTasksState(ctx context.Context, in *TaskList, opts ...client.CallOption) (*TasksTree, error)
	UpdateTask(ctx context.Context, in *Task, opts ...client.CallOption) (*TaskResponse, error)
	DeleteTask(ctx context.Context, in *FincByIdRequest, opts ...client.CallOption) (*EmptyResponse, error)
}

type calendarService struct {
	c    client.Client
	name string
}

func NewCalendarService(name string, c client.Client) CalendarService {
	if c == nil {
		c = client.NewClient()
	}
	if len(name) == 0 {
		name = "calendar"
	}
	return &calendarService{
		c:    c,
		name: name,
	}
}

func (c *calendarService) CreateCalendar(ctx context.Context, in *Calendar, opts ...client.CallOption) (*CalendarResponse, error) {
	req := c.c.NewRequest(c.name, "CalendarService.CreateCalendar", in)
	out := new(CalendarResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *calendarService) GetCalendar(ctx context.Context, in *FincByIdRequest, opts ...client.CallOption) (*CalendarResponse, error) {
	req := c.c.NewRequest(c.name, "CalendarService.GetCalendar", in)
	out := new(CalendarResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *calendarService) RemoveCalendar(ctx context.Context, in *FincByIdRequest, opts ...client.CallOption) (*EmptyResponse, error) {
	req := c.c.NewRequest(c.name, "CalendarService.RemoveCalendar", in)
	out := new(EmptyResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *calendarService) UpdateCalendar(ctx context.Context, in *Calendar, opts ...client.CallOption) (*CalendarResponse, error) {
	req := c.c.NewRequest(c.name, "CalendarService.UpdateCalendar", in)
	out := new(CalendarResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *calendarService) CreateEvent(ctx context.Context, in *Event, opts ...client.CallOption) (*EventResponse, error) {
	req := c.c.NewRequest(c.name, "CalendarService.CreateEvent", in)
	out := new(EventResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *calendarService) GetEvent(ctx context.Context, in *FincByIdRequest, opts ...client.CallOption) (*EventResponse, error) {
	req := c.c.NewRequest(c.name, "CalendarService.GetEvent", in)
	out := new(EventResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *calendarService) UpdateEvent(ctx context.Context, in *EventUpdateRequest, opts ...client.CallOption) (*EventResponse, error) {
	req := c.c.NewRequest(c.name, "CalendarService.UpdateEvent", in)
	out := new(EventResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *calendarService) RemoveEvent(ctx context.Context, in *EventUpdateRequest, opts ...client.CallOption) (*EmptyResponse, error) {
	req := c.c.NewRequest(c.name, "CalendarService.RemoveEvent", in)
	out := new(EmptyResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *calendarService) GetEventsRange(ctx context.Context, in *EventRangeRequest, opts ...client.CallOption) (*EventRangeResponse, error) {
	req := c.c.NewRequest(c.name, "CalendarService.GetEventsRange", in)
	out := new(EventRangeResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *calendarService) CreateTask(ctx context.Context, in *Task, opts ...client.CallOption) (*TaskResponse, error) {
	req := c.c.NewRequest(c.name, "CalendarService.CreateTask", in)
	out := new(TaskResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *calendarService) GetTask(ctx context.Context, in *FincByIdRequest, opts ...client.CallOption) (*TaskResponse, error) {
	req := c.c.NewRequest(c.name, "CalendarService.GetTask", in)
	out := new(TaskResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *calendarService) GetTasks(ctx context.Context, in *TasksRequest, opts ...client.CallOption) (*TasksTree, error) {
	req := c.c.NewRequest(c.name, "CalendarService.GetTasks", in)
	out := new(TasksTree)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *calendarService) UpdateTasksState(ctx context.Context, in *TaskList, opts ...client.CallOption) (*TasksTree, error) {
	req := c.c.NewRequest(c.name, "CalendarService.UpdateTasksState", in)
	out := new(TasksTree)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *calendarService) UpdateTask(ctx context.Context, in *Task, opts ...client.CallOption) (*TaskResponse, error) {
	req := c.c.NewRequest(c.name, "CalendarService.UpdateTask", in)
	out := new(TaskResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *calendarService) DeleteTask(ctx context.Context, in *FincByIdRequest, opts ...client.CallOption) (*EmptyResponse, error) {
	req := c.c.NewRequest(c.name, "CalendarService.DeleteTask", in)
	out := new(EmptyResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for CalendarService service

type CalendarServiceHandler interface {
	CreateCalendar(context.Context, *Calendar, *CalendarResponse) error
	GetCalendar(context.Context, *FincByIdRequest, *CalendarResponse) error
	RemoveCalendar(context.Context, *FincByIdRequest, *EmptyResponse) error
	UpdateCalendar(context.Context, *Calendar, *CalendarResponse) error
	CreateEvent(context.Context, *Event, *EventResponse) error
	GetEvent(context.Context, *FincByIdRequest, *EventResponse) error
	UpdateEvent(context.Context, *EventUpdateRequest, *EventResponse) error
	RemoveEvent(context.Context, *EventUpdateRequest, *EmptyResponse) error
	GetEventsRange(context.Context, *EventRangeRequest, *EventRangeResponse) error
	CreateTask(context.Context, *Task, *TaskResponse) error
	GetTask(context.Context, *FincByIdRequest, *TaskResponse) error
	GetTasks(context.Context, *TasksRequest, *TasksTree) error
	UpdateTasksState(context.Context, *TaskList, *TasksTree) error
	UpdateTask(context.Context, *Task, *TaskResponse) error
	DeleteTask(context.Context, *FincByIdRequest, *EmptyResponse) error
}

func RegisterCalendarServiceHandler(s server.Server, hdlr CalendarServiceHandler, opts ...server.HandlerOption) error {
	type calendarService interface {
		CreateCalendar(ctx context.Context, in *Calendar, out *CalendarResponse) error
		GetCalendar(ctx context.Context, in *FincByIdRequest, out *CalendarResponse) error
		RemoveCalendar(ctx context.Context, in *FincByIdRequest, out *EmptyResponse) error
		UpdateCalendar(ctx context.Context, in *Calendar, out *CalendarResponse) error
		CreateEvent(ctx context.Context, in *Event, out *EventResponse) error
		GetEvent(ctx context.Context, in *FincByIdRequest, out *EventResponse) error
		UpdateEvent(ctx context.Context, in *EventUpdateRequest, out *EventResponse) error
		RemoveEvent(ctx context.Context, in *EventUpdateRequest, out *EmptyResponse) error
		GetEventsRange(ctx context.Context, in *EventRangeRequest, out *EventRangeResponse) error
		CreateTask(ctx context.Context, in *Task, out *TaskResponse) error
		GetTask(ctx context.Context, in *FincByIdRequest, out *TaskResponse) error
		GetTasks(ctx context.Context, in *TasksRequest, out *TasksTree) error
		UpdateTasksState(ctx context.Context, in *TaskList, out *TasksTree) error
		UpdateTask(ctx context.Context, in *Task, out *TaskResponse) error
		DeleteTask(ctx context.Context, in *FincByIdRequest, out *EmptyResponse) error
	}
	type CalendarService struct {
		calendarService
	}
	h := &calendarServiceHandler{hdlr}
	return s.Handle(s.NewHandler(&CalendarService{h}, opts...))
}

type calendarServiceHandler struct {
	CalendarServiceHandler
}

func (h *calendarServiceHandler) CreateCalendar(ctx context.Context, in *Calendar, out *CalendarResponse) error {
	return h.CalendarServiceHandler.CreateCalendar(ctx, in, out)
}

func (h *calendarServiceHandler) GetCalendar(ctx context.Context, in *FincByIdRequest, out *CalendarResponse) error {
	return h.CalendarServiceHandler.GetCalendar(ctx, in, out)
}

func (h *calendarServiceHandler) RemoveCalendar(ctx context.Context, in *FincByIdRequest, out *EmptyResponse) error {
	return h.CalendarServiceHandler.RemoveCalendar(ctx, in, out)
}

func (h *calendarServiceHandler) UpdateCalendar(ctx context.Context, in *Calendar, out *CalendarResponse) error {
	return h.CalendarServiceHandler.UpdateCalendar(ctx, in, out)
}

func (h *calendarServiceHandler) CreateEvent(ctx context.Context, in *Event, out *EventResponse) error {
	return h.CalendarServiceHandler.CreateEvent(ctx, in, out)
}

func (h *calendarServiceHandler) GetEvent(ctx context.Context, in *FincByIdRequest, out *EventResponse) error {
	return h.CalendarServiceHandler.GetEvent(ctx, in, out)
}

func (h *calendarServiceHandler) UpdateEvent(ctx context.Context, in *EventUpdateRequest, out *EventResponse) error {
	return h.CalendarServiceHandler.UpdateEvent(ctx, in, out)
}

func (h *calendarServiceHandler) RemoveEvent(ctx context.Context, in *EventUpdateRequest, out *EmptyResponse) error {
	return h.CalendarServiceHandler.RemoveEvent(ctx, in, out)
}

func (h *calendarServiceHandler) GetEventsRange(ctx context.Context, in *EventRangeRequest, out *EventRangeResponse) error {
	return h.CalendarServiceHandler.GetEventsRange(ctx, in, out)
}

func (h *calendarServiceHandler) CreateTask(ctx context.Context, in *Task, out *TaskResponse) error {
	return h.CalendarServiceHandler.CreateTask(ctx, in, out)
}

func (h *calendarServiceHandler) GetTask(ctx context.Context, in *FincByIdRequest, out *TaskResponse) error {
	return h.CalendarServiceHandler.GetTask(ctx, in, out)
}

func (h *calendarServiceHandler) GetTasks(ctx context.Context, in *TasksRequest, out *TasksTree) error {
	return h.CalendarServiceHandler.GetTasks(ctx, in, out)
}

func (h *calendarServiceHandler) UpdateTasksState(ctx context.Context, in *TaskList, out *TasksTree) error {
	return h.CalendarServiceHandler.UpdateTasksState(ctx, in, out)
}

func (h *calendarServiceHandler) UpdateTask(ctx context.Context, in *Task, out *TaskResponse) error {
	return h.CalendarServiceHandler.UpdateTask(ctx, in, out)
}

func (h *calendarServiceHandler) DeleteTask(ctx context.Context, in *FincByIdRequest, out *EmptyResponse) error {
	return h.CalendarServiceHandler.DeleteTask(ctx, in, out)
}
