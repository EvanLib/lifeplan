// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: proto/calendar.proto

package calendar

import (
	fmt "fmt"
	_ "github.com/gogo/protobuf/gogoproto"
	proto "github.com/gogo/protobuf/proto"
	_ "github.com/golang/protobuf/ptypes/duration"
	_ "github.com/golang/protobuf/ptypes/timestamp"
	math "math"
	time "time"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf
var _ = time.Kitchen

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion2 // please upgrade the proto package

type EmptyResponse struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-" bson:"-"`
	XXX_unrecognized     []byte   `json:"-" bson:"-"`
	XXX_sizecache        int32    `json:"-" bson:"-"`
}

func (m *EmptyResponse) Reset()         { *m = EmptyResponse{} }
func (m *EmptyResponse) String() string { return proto.CompactTextString(m) }
func (*EmptyResponse) ProtoMessage()    {}
func (*EmptyResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_319e4b398a4282cb, []int{0}
}
func (m *EmptyResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_EmptyResponse.Unmarshal(m, b)
}
func (m *EmptyResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_EmptyResponse.Marshal(b, m, deterministic)
}
func (m *EmptyResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_EmptyResponse.Merge(m, src)
}
func (m *EmptyResponse) XXX_Size() int {
	return xxx_messageInfo_EmptyResponse.Size(m)
}
func (m *EmptyResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_EmptyResponse.DiscardUnknown(m)
}

var xxx_messageInfo_EmptyResponse proto.InternalMessageInfo

type Calendar struct {
	Id                   string   `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty" bson:"_id"`
	Name                 string   `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty" bson:"name"`
	Userid               string   `protobuf:"bytes,3,opt,name=userid,proto3" json:"userid,omitempty" bson:"userid"`
	XXX_NoUnkeyedLiteral struct{} `json:"-" bson:"-"`
	XXX_unrecognized     []byte   `json:"-" bson:"-"`
	XXX_sizecache        int32    `json:"-" bson:"-"`
}

func (m *Calendar) Reset()         { *m = Calendar{} }
func (m *Calendar) String() string { return proto.CompactTextString(m) }
func (*Calendar) ProtoMessage()    {}
func (*Calendar) Descriptor() ([]byte, []int) {
	return fileDescriptor_319e4b398a4282cb, []int{1}
}
func (m *Calendar) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Calendar.Unmarshal(m, b)
}
func (m *Calendar) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Calendar.Marshal(b, m, deterministic)
}
func (m *Calendar) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Calendar.Merge(m, src)
}
func (m *Calendar) XXX_Size() int {
	return xxx_messageInfo_Calendar.Size(m)
}
func (m *Calendar) XXX_DiscardUnknown() {
	xxx_messageInfo_Calendar.DiscardUnknown(m)
}

var xxx_messageInfo_Calendar proto.InternalMessageInfo

func (m *Calendar) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *Calendar) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Calendar) GetUserid() string {
	if m != nil {
		return m.Userid
	}
	return ""
}

type CalendarResponse struct {
	Status               string    `protobuf:"bytes,1,opt,name=status,proto3" json:"status,omitempty"`
	Calendar             *Calendar `protobuf:"bytes,2,opt,name=calendar,proto3" json:"calendar,omitempty"`
	Error                string    `protobuf:"bytes,3,opt,name=error,proto3" json:"error,omitempty"`
	XXX_NoUnkeyedLiteral struct{}  `json:"-" bson:"-"`
	XXX_unrecognized     []byte    `json:"-" bson:"-"`
	XXX_sizecache        int32     `json:"-" bson:"-"`
}

func (m *CalendarResponse) Reset()         { *m = CalendarResponse{} }
func (m *CalendarResponse) String() string { return proto.CompactTextString(m) }
func (*CalendarResponse) ProtoMessage()    {}
func (*CalendarResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_319e4b398a4282cb, []int{2}
}
func (m *CalendarResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CalendarResponse.Unmarshal(m, b)
}
func (m *CalendarResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CalendarResponse.Marshal(b, m, deterministic)
}
func (m *CalendarResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CalendarResponse.Merge(m, src)
}
func (m *CalendarResponse) XXX_Size() int {
	return xxx_messageInfo_CalendarResponse.Size(m)
}
func (m *CalendarResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_CalendarResponse.DiscardUnknown(m)
}

var xxx_messageInfo_CalendarResponse proto.InternalMessageInfo

func (m *CalendarResponse) GetStatus() string {
	if m != nil {
		return m.Status
	}
	return ""
}

func (m *CalendarResponse) GetCalendar() *Calendar {
	if m != nil {
		return m.Calendar
	}
	return nil
}

func (m *CalendarResponse) GetError() string {
	if m != nil {
		return m.Error
	}
	return ""
}

type FincByIdRequest struct {
	// @inject_tag: validate:"required,hexadecimal,len=24"
	Id                   string   `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty" validate:"required,hexadecimal,len=24"`
	XXX_NoUnkeyedLiteral struct{} `json:"-" bson:"-"`
	XXX_unrecognized     []byte   `json:"-" bson:"-"`
	XXX_sizecache        int32    `json:"-" bson:"-"`
}

func (m *FincByIdRequest) Reset()         { *m = FincByIdRequest{} }
func (m *FincByIdRequest) String() string { return proto.CompactTextString(m) }
func (*FincByIdRequest) ProtoMessage()    {}
func (*FincByIdRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_319e4b398a4282cb, []int{3}
}
func (m *FincByIdRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_FincByIdRequest.Unmarshal(m, b)
}
func (m *FincByIdRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_FincByIdRequest.Marshal(b, m, deterministic)
}
func (m *FincByIdRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_FincByIdRequest.Merge(m, src)
}
func (m *FincByIdRequest) XXX_Size() int {
	return xxx_messageInfo_FincByIdRequest.Size(m)
}
func (m *FincByIdRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_FincByIdRequest.DiscardUnknown(m)
}

var xxx_messageInfo_FincByIdRequest proto.InternalMessageInfo

func (m *FincByIdRequest) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

type EventRangeRequest struct {
	Userid               string    `protobuf:"bytes,1,opt,name=userid,proto3" json:"userid,omitempty"`
	Start                time.Time `protobuf:"bytes,4,opt,name=start,proto3,stdtime" json:"start" bson:"start"`
	End                  time.Time `protobuf:"bytes,5,opt,name=end,proto3,stdtime" json:"end" bson:"end"`
	XXX_NoUnkeyedLiteral struct{}  `json:"-" bson:"-"`
	XXX_unrecognized     []byte    `json:"-" bson:"-"`
	XXX_sizecache        int32     `json:"-" bson:"-"`
}

func (m *EventRangeRequest) Reset()         { *m = EventRangeRequest{} }
func (m *EventRangeRequest) String() string { return proto.CompactTextString(m) }
func (*EventRangeRequest) ProtoMessage()    {}
func (*EventRangeRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_319e4b398a4282cb, []int{4}
}
func (m *EventRangeRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_EventRangeRequest.Unmarshal(m, b)
}
func (m *EventRangeRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_EventRangeRequest.Marshal(b, m, deterministic)
}
func (m *EventRangeRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_EventRangeRequest.Merge(m, src)
}
func (m *EventRangeRequest) XXX_Size() int {
	return xxx_messageInfo_EventRangeRequest.Size(m)
}
func (m *EventRangeRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_EventRangeRequest.DiscardUnknown(m)
}

var xxx_messageInfo_EventRangeRequest proto.InternalMessageInfo

func (m *EventRangeRequest) GetUserid() string {
	if m != nil {
		return m.Userid
	}
	return ""
}

func (m *EventRangeRequest) GetStart() time.Time {
	if m != nil {
		return m.Start
	}
	return time.Time{}
}

func (m *EventRangeRequest) GetEnd() time.Time {
	if m != nil {
		return m.End
	}
	return time.Time{}
}

type EventUpdateRequest struct {
	Updatetype           int32    `protobuf:"varint,1,opt,name=updatetype,proto3" json:"updatetype,omitempty"`
	Event                *Event   `protobuf:"bytes,2,opt,name=event,proto3" json:"event,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-" bson:"-"`
	XXX_unrecognized     []byte   `json:"-" bson:"-"`
	XXX_sizecache        int32    `json:"-" bson:"-"`
}

func (m *EventUpdateRequest) Reset()         { *m = EventUpdateRequest{} }
func (m *EventUpdateRequest) String() string { return proto.CompactTextString(m) }
func (*EventUpdateRequest) ProtoMessage()    {}
func (*EventUpdateRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_319e4b398a4282cb, []int{5}
}
func (m *EventUpdateRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_EventUpdateRequest.Unmarshal(m, b)
}
func (m *EventUpdateRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_EventUpdateRequest.Marshal(b, m, deterministic)
}
func (m *EventUpdateRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_EventUpdateRequest.Merge(m, src)
}
func (m *EventUpdateRequest) XXX_Size() int {
	return xxx_messageInfo_EventUpdateRequest.Size(m)
}
func (m *EventUpdateRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_EventUpdateRequest.DiscardUnknown(m)
}

var xxx_messageInfo_EventUpdateRequest proto.InternalMessageInfo

func (m *EventUpdateRequest) GetUpdatetype() int32 {
	if m != nil {
		return m.Updatetype
	}
	return 0
}

func (m *EventUpdateRequest) GetEvent() *Event {
	if m != nil {
		return m.Event
	}
	return nil
}

type EventRangeResponse struct {
	Events               []*Event `protobuf:"bytes,1,rep,name=events,proto3" json:"events,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-" bson:"-"`
	XXX_unrecognized     []byte   `json:"-" bson:"-"`
	XXX_sizecache        int32    `json:"-" bson:"-"`
}

func (m *EventRangeResponse) Reset()         { *m = EventRangeResponse{} }
func (m *EventRangeResponse) String() string { return proto.CompactTextString(m) }
func (*EventRangeResponse) ProtoMessage()    {}
func (*EventRangeResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_319e4b398a4282cb, []int{6}
}
func (m *EventRangeResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_EventRangeResponse.Unmarshal(m, b)
}
func (m *EventRangeResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_EventRangeResponse.Marshal(b, m, deterministic)
}
func (m *EventRangeResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_EventRangeResponse.Merge(m, src)
}
func (m *EventRangeResponse) XXX_Size() int {
	return xxx_messageInfo_EventRangeResponse.Size(m)
}
func (m *EventRangeResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_EventRangeResponse.DiscardUnknown(m)
}

var xxx_messageInfo_EventRangeResponse proto.InternalMessageInfo

func (m *EventRangeResponse) GetEvents() []*Event {
	if m != nil {
		return m.Events
	}
	return nil
}

type EventResponse struct {
	Status               int32    `protobuf:"varint,1,opt,name=status,proto3" json:"status,omitempty"`
	Err                  string   `protobuf:"bytes,2,opt,name=err,proto3" json:"err,omitempty"`
	Event                *Event   `protobuf:"bytes,3,opt,name=event,proto3" json:"event,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-" bson:"-"`
	XXX_unrecognized     []byte   `json:"-" bson:"-"`
	XXX_sizecache        int32    `json:"-" bson:"-"`
}

func (m *EventResponse) Reset()         { *m = EventResponse{} }
func (m *EventResponse) String() string { return proto.CompactTextString(m) }
func (*EventResponse) ProtoMessage()    {}
func (*EventResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_319e4b398a4282cb, []int{7}
}
func (m *EventResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_EventResponse.Unmarshal(m, b)
}
func (m *EventResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_EventResponse.Marshal(b, m, deterministic)
}
func (m *EventResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_EventResponse.Merge(m, src)
}
func (m *EventResponse) XXX_Size() int {
	return xxx_messageInfo_EventResponse.Size(m)
}
func (m *EventResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_EventResponse.DiscardUnknown(m)
}

var xxx_messageInfo_EventResponse proto.InternalMessageInfo

func (m *EventResponse) GetStatus() int32 {
	if m != nil {
		return m.Status
	}
	return 0
}

func (m *EventResponse) GetErr() string {
	if m != nil {
		return m.Err
	}
	return ""
}

func (m *EventResponse) GetEvent() *Event {
	if m != nil {
		return m.Event
	}
	return nil
}

type Event struct {
	Id                   string        `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty" bson:"_id"`
	Title                string        `protobuf:"bytes,2,opt,name=title,proto3" json:"title,omitempty" bson:"title"`
	Userid               string        `protobuf:"bytes,3,opt,name=userid,proto3" json:"userid,omitempty" bson:"userid"`
	Start                time.Time     `protobuf:"bytes,4,opt,name=start,proto3,stdtime" json:"start" bson:"start"`
	End                  time.Time     `protobuf:"bytes,5,opt,name=end,proto3,stdtime" json:"end" bson:"end"`
	Duration             time.Duration `protobuf:"bytes,6,opt,name=duration,proto3,stdduration" json:"duration" bson:"duration"`
	Recurring            bool          `protobuf:"varint,7,opt,name=recurring,proto3" json:"recurring,omitempty" bson:"recurring"`
	Allday               bool          `protobuf:"varint,8,opt,name=allday,proto3" json:"allday,omitempty" bson:"allday"`
	Rrule                string        `protobuf:"bytes,9,opt,name=rrule,proto3" json:"rrule,omitempty" bson:"rrule"`
	Exrule               string        `protobuf:"bytes,10,opt,name=exrule,proto3" json:"exrule,omitempty" bson:"exrule"`
	Exdates              []time.Time   `protobuf:"bytes,11,rep,name=exdates,proto3,stdtime" json:"exdates" bson:"exdates"`
	Tasknodeid           string        `protobuf:"bytes,12,opt,name=tasknodeid,proto3" json:"tasknodeid,omitempty" bson:"tasknodeid"`
	XXX_NoUnkeyedLiteral struct{}      `json:"-" bson:"-"`
	XXX_unrecognized     []byte        `json:"-" bson:"-"`
	XXX_sizecache        int32         `json:"-" bson:"-"`
}

func (m *Event) Reset()         { *m = Event{} }
func (m *Event) String() string { return proto.CompactTextString(m) }
func (*Event) ProtoMessage()    {}
func (*Event) Descriptor() ([]byte, []int) {
	return fileDescriptor_319e4b398a4282cb, []int{8}
}
func (m *Event) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Event.Unmarshal(m, b)
}
func (m *Event) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Event.Marshal(b, m, deterministic)
}
func (m *Event) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Event.Merge(m, src)
}
func (m *Event) XXX_Size() int {
	return xxx_messageInfo_Event.Size(m)
}
func (m *Event) XXX_DiscardUnknown() {
	xxx_messageInfo_Event.DiscardUnknown(m)
}

var xxx_messageInfo_Event proto.InternalMessageInfo

func (m *Event) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *Event) GetTitle() string {
	if m != nil {
		return m.Title
	}
	return ""
}

func (m *Event) GetUserid() string {
	if m != nil {
		return m.Userid
	}
	return ""
}

func (m *Event) GetStart() time.Time {
	if m != nil {
		return m.Start
	}
	return time.Time{}
}

func (m *Event) GetEnd() time.Time {
	if m != nil {
		return m.End
	}
	return time.Time{}
}

func (m *Event) GetDuration() time.Duration {
	if m != nil {
		return m.Duration
	}
	return 0
}

func (m *Event) GetRecurring() bool {
	if m != nil {
		return m.Recurring
	}
	return false
}

func (m *Event) GetAllday() bool {
	if m != nil {
		return m.Allday
	}
	return false
}

func (m *Event) GetRrule() string {
	if m != nil {
		return m.Rrule
	}
	return ""
}

func (m *Event) GetExrule() string {
	if m != nil {
		return m.Exrule
	}
	return ""
}

func (m *Event) GetExdates() []time.Time {
	if m != nil {
		return m.Exdates
	}
	return nil
}

func (m *Event) GetTasknodeid() string {
	if m != nil {
		return m.Tasknodeid
	}
	return ""
}

type Task struct {
	Id                   string   `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty" bson:"_id"`
	Categoryid           string   `protobuf:"bytes,2,opt,name=categoryid,proto3" json:"categoryid,omitempty" bson:"categoryid"`
	State                int32    `protobuf:"varint,3,opt,name=state,proto3" json:"state,omitempty" bson:"state"`
	Task                 string   `protobuf:"bytes,4,opt,name=task,proto3" json:"task,omitempty" bson:"task"`
	Category             string   `protobuf:"bytes,5,opt,name=category,proto3" json:"category,omitempty" bson:"category"`
	Parent               string   `protobuf:"bytes,6,opt,name=parent,proto3" json:"parent,omitempty" bson:"parent"`
	Children             []*Task  `protobuf:"bytes,7,rep,name=children,proto3" json:"children,omitempty" bson:"-"`
	XXX_NoUnkeyedLiteral struct{} `json:"-" bson:"-"`
	XXX_unrecognized     []byte   `json:"-" bson:"-"`
	XXX_sizecache        int32    `json:"-" bson:"-"`
}

func (m *Task) Reset()         { *m = Task{} }
func (m *Task) String() string { return proto.CompactTextString(m) }
func (*Task) ProtoMessage()    {}
func (*Task) Descriptor() ([]byte, []int) {
	return fileDescriptor_319e4b398a4282cb, []int{9}
}
func (m *Task) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Task.Unmarshal(m, b)
}
func (m *Task) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Task.Marshal(b, m, deterministic)
}
func (m *Task) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Task.Merge(m, src)
}
func (m *Task) XXX_Size() int {
	return xxx_messageInfo_Task.Size(m)
}
func (m *Task) XXX_DiscardUnknown() {
	xxx_messageInfo_Task.DiscardUnknown(m)
}

var xxx_messageInfo_Task proto.InternalMessageInfo

func (m *Task) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *Task) GetCategoryid() string {
	if m != nil {
		return m.Categoryid
	}
	return ""
}

func (m *Task) GetState() int32 {
	if m != nil {
		return m.State
	}
	return 0
}

func (m *Task) GetTask() string {
	if m != nil {
		return m.Task
	}
	return ""
}

func (m *Task) GetCategory() string {
	if m != nil {
		return m.Category
	}
	return ""
}

func (m *Task) GetParent() string {
	if m != nil {
		return m.Parent
	}
	return ""
}

func (m *Task) GetChildren() []*Task {
	if m != nil {
		return m.Children
	}
	return nil
}

type TreeNode struct {
	Task                 *Task       `protobuf:"bytes,1,opt,name=task,proto3" json:"task,omitempty"`
	Subtasks             []*TreeNode `protobuf:"bytes,2,rep,name=subtasks,proto3" json:"subtasks,omitempty"`
	XXX_NoUnkeyedLiteral struct{}    `json:"-" bson:"-"`
	XXX_unrecognized     []byte      `json:"-" bson:"-"`
	XXX_sizecache        int32       `json:"-" bson:"-"`
}

func (m *TreeNode) Reset()         { *m = TreeNode{} }
func (m *TreeNode) String() string { return proto.CompactTextString(m) }
func (*TreeNode) ProtoMessage()    {}
func (*TreeNode) Descriptor() ([]byte, []int) {
	return fileDescriptor_319e4b398a4282cb, []int{10}
}
func (m *TreeNode) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TreeNode.Unmarshal(m, b)
}
func (m *TreeNode) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TreeNode.Marshal(b, m, deterministic)
}
func (m *TreeNode) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TreeNode.Merge(m, src)
}
func (m *TreeNode) XXX_Size() int {
	return xxx_messageInfo_TreeNode.Size(m)
}
func (m *TreeNode) XXX_DiscardUnknown() {
	xxx_messageInfo_TreeNode.DiscardUnknown(m)
}

var xxx_messageInfo_TreeNode proto.InternalMessageInfo

func (m *TreeNode) GetTask() *Task {
	if m != nil {
		return m.Task
	}
	return nil
}

func (m *TreeNode) GetSubtasks() []*TreeNode {
	if m != nil {
		return m.Subtasks
	}
	return nil
}

type TasksTree struct {
	Nodes                []*TreeNode `protobuf:"bytes,1,rep,name=nodes,proto3" json:"nodes,omitempty"`
	Categoryid           string      `protobuf:"bytes,2,opt,name=categoryid,proto3" json:"categoryid,omitempty"`
	XXX_NoUnkeyedLiteral struct{}    `json:"-" bson:"-"`
	XXX_unrecognized     []byte      `json:"-" bson:"-"`
	XXX_sizecache        int32       `json:"-" bson:"-"`
}

func (m *TasksTree) Reset()         { *m = TasksTree{} }
func (m *TasksTree) String() string { return proto.CompactTextString(m) }
func (*TasksTree) ProtoMessage()    {}
func (*TasksTree) Descriptor() ([]byte, []int) {
	return fileDescriptor_319e4b398a4282cb, []int{11}
}
func (m *TasksTree) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TasksTree.Unmarshal(m, b)
}
func (m *TasksTree) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TasksTree.Marshal(b, m, deterministic)
}
func (m *TasksTree) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TasksTree.Merge(m, src)
}
func (m *TasksTree) XXX_Size() int {
	return xxx_messageInfo_TasksTree.Size(m)
}
func (m *TasksTree) XXX_DiscardUnknown() {
	xxx_messageInfo_TasksTree.DiscardUnknown(m)
}

var xxx_messageInfo_TasksTree proto.InternalMessageInfo

func (m *TasksTree) GetNodes() []*TreeNode {
	if m != nil {
		return m.Nodes
	}
	return nil
}

func (m *TasksTree) GetCategoryid() string {
	if m != nil {
		return m.Categoryid
	}
	return ""
}

type TasksRequest struct {
	Categoryid           string    `protobuf:"bytes,1,opt,name=categoryid,proto3" json:"categoryid,omitempty" bson:"categoryid"`
	Filter               string    `protobuf:"bytes,2,opt,name=filter,proto3" json:"filter,omitempty" bson:"filter"`
	Categorytimestamp    time.Time `protobuf:"bytes,3,opt,name=categorytimestamp,proto3,stdtime" json:"categorytimestamp" bson:"categorytimestamp"`
	XXX_NoUnkeyedLiteral struct{}  `json:"-" bson:"-"`
	XXX_unrecognized     []byte    `json:"-" bson:"-"`
	XXX_sizecache        int32     `json:"-" bson:"-"`
}

func (m *TasksRequest) Reset()         { *m = TasksRequest{} }
func (m *TasksRequest) String() string { return proto.CompactTextString(m) }
func (*TasksRequest) ProtoMessage()    {}
func (*TasksRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_319e4b398a4282cb, []int{12}
}
func (m *TasksRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TasksRequest.Unmarshal(m, b)
}
func (m *TasksRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TasksRequest.Marshal(b, m, deterministic)
}
func (m *TasksRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TasksRequest.Merge(m, src)
}
func (m *TasksRequest) XXX_Size() int {
	return xxx_messageInfo_TasksRequest.Size(m)
}
func (m *TasksRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_TasksRequest.DiscardUnknown(m)
}

var xxx_messageInfo_TasksRequest proto.InternalMessageInfo

func (m *TasksRequest) GetCategoryid() string {
	if m != nil {
		return m.Categoryid
	}
	return ""
}

func (m *TasksRequest) GetFilter() string {
	if m != nil {
		return m.Filter
	}
	return ""
}

func (m *TasksRequest) GetCategorytimestamp() time.Time {
	if m != nil {
		return m.Categorytimestamp
	}
	return time.Time{}
}

type TaskList struct {
	Categoryid           string           `protobuf:"bytes,2,opt,name=categoryid,proto3" json:"categoryid,omitempty" bson:"categoryid"`
	Categorytimestamp    time.Time        `protobuf:"bytes,3,opt,name=categorytimestamp,proto3,stdtime" json:"categorytimestamp" bson:"categorytimestamp"`
	TaskStateMap         map[string]int32 `protobuf:"bytes,4,rep,name=TaskStateMap,proto3" json:"TaskStateMap,omitempty" bson:"taskstatemap" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"varint,2,opt,name=value,proto3"`
	XXX_NoUnkeyedLiteral struct{}         `json:"-" bson:"-"`
	XXX_unrecognized     []byte           `json:"-" bson:"-"`
	XXX_sizecache        int32            `json:"-" bson:"-"`
}

func (m *TaskList) Reset()         { *m = TaskList{} }
func (m *TaskList) String() string { return proto.CompactTextString(m) }
func (*TaskList) ProtoMessage()    {}
func (*TaskList) Descriptor() ([]byte, []int) {
	return fileDescriptor_319e4b398a4282cb, []int{13}
}
func (m *TaskList) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TaskList.Unmarshal(m, b)
}
func (m *TaskList) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TaskList.Marshal(b, m, deterministic)
}
func (m *TaskList) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TaskList.Merge(m, src)
}
func (m *TaskList) XXX_Size() int {
	return xxx_messageInfo_TaskList.Size(m)
}
func (m *TaskList) XXX_DiscardUnknown() {
	xxx_messageInfo_TaskList.DiscardUnknown(m)
}

var xxx_messageInfo_TaskList proto.InternalMessageInfo

func (m *TaskList) GetCategoryid() string {
	if m != nil {
		return m.Categoryid
	}
	return ""
}

func (m *TaskList) GetCategorytimestamp() time.Time {
	if m != nil {
		return m.Categorytimestamp
	}
	return time.Time{}
}

func (m *TaskList) GetTaskStateMap() map[string]int32 {
	if m != nil {
		return m.TaskStateMap
	}
	return nil
}

type TaskResponse struct {
	Status               int32    `protobuf:"varint,1,opt,name=status,proto3" json:"status,omitempty"`
	Err                  string   `protobuf:"bytes,2,opt,name=err,proto3" json:"err,omitempty"`
	Task                 *Task    `protobuf:"bytes,3,opt,name=task,proto3" json:"task,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-" bson:"-"`
	XXX_unrecognized     []byte   `json:"-" bson:"-"`
	XXX_sizecache        int32    `json:"-" bson:"-"`
}

func (m *TaskResponse) Reset()         { *m = TaskResponse{} }
func (m *TaskResponse) String() string { return proto.CompactTextString(m) }
func (*TaskResponse) ProtoMessage()    {}
func (*TaskResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_319e4b398a4282cb, []int{14}
}
func (m *TaskResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TaskResponse.Unmarshal(m, b)
}
func (m *TaskResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TaskResponse.Marshal(b, m, deterministic)
}
func (m *TaskResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TaskResponse.Merge(m, src)
}
func (m *TaskResponse) XXX_Size() int {
	return xxx_messageInfo_TaskResponse.Size(m)
}
func (m *TaskResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_TaskResponse.DiscardUnknown(m)
}

var xxx_messageInfo_TaskResponse proto.InternalMessageInfo

func (m *TaskResponse) GetStatus() int32 {
	if m != nil {
		return m.Status
	}
	return 0
}

func (m *TaskResponse) GetErr() string {
	if m != nil {
		return m.Err
	}
	return ""
}

func (m *TaskResponse) GetTask() *Task {
	if m != nil {
		return m.Task
	}
	return nil
}

func init() {
	proto.RegisterType((*EmptyResponse)(nil), "calendar.EmptyResponse")
	proto.RegisterType((*Calendar)(nil), "calendar.Calendar")
	proto.RegisterType((*CalendarResponse)(nil), "calendar.CalendarResponse")
	proto.RegisterType((*FincByIdRequest)(nil), "calendar.FincByIdRequest")
	proto.RegisterType((*EventRangeRequest)(nil), "calendar.EventRangeRequest")
	proto.RegisterType((*EventUpdateRequest)(nil), "calendar.EventUpdateRequest")
	proto.RegisterType((*EventRangeResponse)(nil), "calendar.EventRangeResponse")
	proto.RegisterType((*EventResponse)(nil), "calendar.EventResponse")
	proto.RegisterType((*Event)(nil), "calendar.Event")
	proto.RegisterType((*Task)(nil), "calendar.Task")
	proto.RegisterType((*TreeNode)(nil), "calendar.TreeNode")
	proto.RegisterType((*TasksTree)(nil), "calendar.TasksTree")
	proto.RegisterType((*TasksRequest)(nil), "calendar.TasksRequest")
	proto.RegisterType((*TaskList)(nil), "calendar.TaskList")
	proto.RegisterMapType((map[string]int32)(nil), "calendar.TaskList.TaskStateMapEntry")
	proto.RegisterType((*TaskResponse)(nil), "calendar.TaskResponse")
}

func init() { proto.RegisterFile("proto/calendar.proto", fileDescriptor_319e4b398a4282cb) }

var fileDescriptor_319e4b398a4282cb = []byte{
	// 1207 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xcc, 0x56, 0xcd, 0x6e, 0xdb, 0xc6,
	0x13, 0xff, 0x4b, 0x32, 0x65, 0x6a, 0x64, 0x4b, 0xf6, 0xda, 0x7f, 0x87, 0x51, 0x02, 0x33, 0xdd,
	0xa6, 0xad, 0x73, 0x88, 0x0c, 0xb8, 0x09, 0x12, 0xa4, 0x5f, 0x81, 0x62, 0xd7, 0xe8, 0x27, 0x8a,
	0x8d, 0x73, 0x6a, 0xd1, 0x96, 0x16, 0x37, 0x0a, 0x61, 0x89, 0x54, 0xc9, 0xa5, 0x11, 0xf5, 0x29,
	0x7a, 0x2c, 0xd0, 0x67, 0x29, 0xd0, 0x63, 0x5f, 0x20, 0x57, 0xb5, 0xcf, 0xa0, 0x27, 0x28, 0x76,
	0x67, 0x97, 0xa4, 0x68, 0xc5, 0x76, 0x7c, 0x28, 0x7a, 0xe3, 0xce, 0xfc, 0xe6, 0xb7, 0xf3, 0xb5,
	0xc3, 0x81, 0xcd, 0x71, 0x1c, 0x89, 0x68, 0xb7, 0xef, 0x0d, 0x79, 0xe8, 0x7b, 0x71, 0x57, 0x1d,
	0x89, 0x6d, 0xce, 0x1d, 0x77, 0x10, 0x45, 0x83, 0x21, 0xdf, 0x55, 0xf2, 0xe3, 0xf4, 0xf9, 0xae,
	0x08, 0x46, 0x3c, 0x11, 0xde, 0x68, 0x8c, 0xd0, 0xce, 0x76, 0x19, 0xe0, 0xa7, 0xb1, 0x27, 0x82,
	0x28, 0xd4, 0xfa, 0xbb, 0x83, 0x40, 0xbc, 0x48, 0x8f, 0xbb, 0xfd, 0x68, 0xb4, 0x3b, 0x88, 0x06,
	0x51, 0x0e, 0x94, 0x27, 0xbc, 0x5d, 0x7e, 0x21, 0x9c, 0xb6, 0x61, 0xf5, 0x60, 0x34, 0x16, 0x13,
	0xc6, 0x93, 0x71, 0x14, 0x26, 0x9c, 0xfe, 0x0c, 0xf6, 0x13, 0xed, 0x0c, 0xd9, 0x86, 0x6a, 0xe0,
	0x3b, 0x95, 0x5b, 0x95, 0x9d, 0x46, 0xaf, 0x35, 0x9b, 0xba, 0x70, 0x9c, 0x44, 0xe1, 0x23, 0xfa,
	0x43, 0xe0, 0x53, 0x56, 0x0d, 0x7c, 0xf2, 0x36, 0x2c, 0x85, 0xde, 0x88, 0x3b, 0x55, 0x85, 0x68,
	0xcf, 0xa6, 0x6e, 0x13, 0x11, 0x52, 0x4a, 0x99, 0x52, 0x92, 0x3b, 0x50, 0x4f, 0x13, 0x1e, 0x07,
	0xbe, 0x53, 0x53, 0xb0, 0xf5, 0xd9, 0xd4, 0x5d, 0x45, 0x18, 0xca, 0x29, 0xd3, 0x00, 0x3a, 0x86,
	0x35, 0x73, 0xb7, 0xf1, 0x87, 0x6c, 0x41, 0x3d, 0x11, 0x9e, 0x48, 0x13, 0xf4, 0x83, 0xe9, 0x13,
	0xe9, 0x42, 0x96, 0x34, 0x75, 0x7f, 0x73, 0x8f, 0x74, 0xb3, 0xac, 0x66, 0x2c, 0x19, 0x86, 0x6c,
	0x82, 0xc5, 0xe3, 0x38, 0x8a, 0xd1, 0x0b, 0x86, 0x07, 0xfa, 0x16, 0xb4, 0x3f, 0x0d, 0xc2, 0x7e,
	0x6f, 0xf2, 0x99, 0xcf, 0xf8, 0x4f, 0x29, 0x4f, 0x04, 0x69, 0xe5, 0x41, 0xcb, 0x20, 0xe9, 0xef,
	0x15, 0x58, 0x3f, 0x38, 0xe5, 0xa1, 0x60, 0x5e, 0x38, 0xe0, 0x06, 0xb5, 0x95, 0x45, 0xa5, 0xdd,
	0xc2, 0x13, 0xf9, 0x1c, 0xac, 0x44, 0x78, 0xb1, 0x70, 0x96, 0x94, 0x4f, 0x9d, 0x2e, 0x96, 0xab,
	0x6b, 0xaa, 0xd0, 0x3d, 0x32, 0xf5, 0xec, 0x39, 0x7f, 0x4e, 0xdd, 0xff, 0xcd, 0xa6, 0xee, 0x0a,
	0x26, 0x43, 0x99, 0xd1, 0x5f, 0xfe, 0x72, 0x2b, 0x0c, 0x29, 0xc8, 0x3e, 0xd4, 0x78, 0xe8, 0x3b,
	0xd6, 0x85, 0x4c, 0x5b, 0x9a, 0x49, 0xd7, 0x87, 0x87, 0x3e, 0xf2, 0x48, 0x73, 0xfa, 0x2d, 0x10,
	0xe5, 0xfe, 0xb3, 0xb1, 0xef, 0x89, 0xcc, 0xff, 0x6d, 0x80, 0x54, 0x09, 0xc4, 0x64, 0xcc, 0x55,
	0x0c, 0x16, 0x2b, 0x48, 0xc8, 0x3b, 0x60, 0x71, 0x69, 0xa5, 0x73, 0xdb, 0xce, 0x73, 0x8b, 0xb9,
	0x40, 0x2d, 0xfd, 0x48, 0x93, 0xeb, 0xdc, 0xe8, 0x9a, 0xbd, 0x07, 0x75, 0xa5, 0x96, 0x35, 0xab,
	0x2d, 0xb2, 0xd6, 0x6a, 0xfa, 0x23, 0xac, 0xa2, 0x60, 0x71, 0xb5, 0xad, 0xac, 0xda, 0x6b, 0x50,
	0xe3, 0x31, 0x16, 0xba, 0xc1, 0xe4, 0x67, 0xee, 0x60, 0xed, 0x5c, 0x07, 0x7f, 0xb3, 0xc0, 0x52,
	0x82, 0x0b, 0x9b, 0xf9, 0x5d, 0xb0, 0x44, 0x20, 0x86, 0xa6, 0x9b, 0xd7, 0xf2, 0xca, 0x28, 0x31,
	0x65, 0xa8, 0x7e, 0x83, 0x7e, 0xfe, 0xef, 0x35, 0x03, 0x61, 0x60, 0x9b, 0x79, 0xe1, 0xd4, 0x15,
	0xd5, 0xf5, 0x33, 0x54, 0xfb, 0x1a, 0xd0, 0xbb, 0xa1, 0x99, 0xda, 0xc8, 0x64, 0x0c, 0xe9, 0xaf,
	0x92, 0x2e, 0xe3, 0x21, 0x7b, 0xd0, 0x88, 0x79, 0x3f, 0x8d, 0xe3, 0x20, 0x1c, 0x38, 0xcb, 0xb7,
	0x2a, 0x3b, 0x76, 0x6f, 0x73, 0x36, 0x75, 0xd7, 0xd0, 0x2a, 0x53, 0x51, 0x96, 0xc3, 0x64, 0x12,
	0xbd, 0xe1, 0xd0, 0xf7, 0x26, 0x8e, 0xad, 0x0c, 0x0a, 0x49, 0x44, 0x39, 0x65, 0x1a, 0x20, 0xeb,
	0x12, 0xc7, 0xe9, 0x90, 0x3b, 0x8d, 0x72, 0x5d, 0x94, 0x98, 0x32, 0x54, 0x4b, 0x4a, 0xfe, 0x52,
	0x01, 0xa1, 0x5c, 0x17, 0x94, 0x53, 0xa6, 0x01, 0xe4, 0x1b, 0x58, 0xe6, 0x2f, 0x65, 0xab, 0x27,
	0x4e, 0x53, 0x35, 0xe8, 0x79, 0xf9, 0xec, 0xe8, 0x2c, 0xb4, 0x0c, 0x97, 0x32, 0xc4, 0x9c, 0x1a,
	0x1a, 0x72, 0x1f, 0x40, 0x78, 0xc9, 0x49, 0x18, 0xf9, 0x3c, 0xf0, 0x9d, 0x15, 0xe5, 0xc0, 0xff,
	0x67, 0x53, 0x77, 0x5d, 0x77, 0x50, 0xa6, 0xa3, 0xac, 0x00, 0xa4, 0x7f, 0x54, 0x61, 0xe9, 0xc8,
	0x4b, 0x4e, 0x2e, 0x6c, 0xce, 0xfb, 0x00, 0x7d, 0x4f, 0xf0, 0x41, 0x14, 0x4f, 0x02, 0x5f, 0x77,
	0x68, 0x81, 0x3f, 0xd7, 0x51, 0x56, 0x00, 0xca, 0xdc, 0xc9, 0x07, 0xc4, 0x55, 0xab, 0x5a, 0xc5,
	0xdc, 0x29, 0x31, 0x65, 0xa8, 0x96, 0x83, 0x5c, 0x7a, 0xa5, 0xfa, 0x74, 0x6e, 0x90, 0x4b, 0x29,
	0x65, 0x4a, 0x49, 0x76, 0xe5, 0xc4, 0x45, 0x6a, 0xd5, 0x86, 0x8d, 0xde, 0x46, 0xde, 0x1c, 0x46,
	0x43, 0x59, 0x06, 0x92, 0x15, 0x19, 0x7b, 0xb1, 0x7c, 0xa3, 0xf5, 0x72, 0x45, 0x50, 0x4e, 0x99,
	0x06, 0x90, 0x87, 0x60, 0xf7, 0x5f, 0x04, 0x43, 0x3f, 0xe6, 0xa1, 0xb3, 0xac, 0x4a, 0xd2, 0xca,
	0x1f, 0xb4, 0xcc, 0x50, 0x6f, 0x65, 0x36, 0x75, 0x6d, 0x34, 0xbe, 0x2b, 0x2f, 0xd1, 0x68, 0xfa,
	0x3d, 0xd8, 0x47, 0x31, 0xe7, 0x5f, 0x47, 0x3e, 0x27, 0x54, 0x87, 0x51, 0x51, 0x9d, 0x5d, 0x62,
	0xd0, 0x51, 0x74, 0xc1, 0x4e, 0xd2, 0x63, 0xf9, 0x99, 0x38, 0x55, 0x75, 0x53, 0xe1, 0xbf, 0x61,
	0x98, 0x58, 0x86, 0xa1, 0xcf, 0xa0, 0x21, 0xad, 0x13, 0xa9, 0x22, 0x3b, 0x60, 0xc9, 0xca, 0x99,
	0xb9, 0xb6, 0xc8, 0x12, 0x01, 0x72, 0xbe, 0x96, 0x0b, 0x56, 0xac, 0x0c, 0xfd, 0xbb, 0x02, 0x2b,
	0x8a, 0xd7, 0x0c, 0xe4, 0xf9, 0x0a, 0x57, 0x2e, 0x5b, 0xe1, 0x3b, 0x50, 0x7f, 0x1e, 0x0c, 0x05,
	0xd7, 0xb3, 0xb1, 0x98, 0x63, 0x94, 0x53, 0xa6, 0x01, 0x24, 0x84, 0x75, 0x63, 0x98, 0x2d, 0x15,
	0x7a, 0x7a, 0x9e, 0xd7, 0xff, 0xb7, 0x75, 0xff, 0x3b, 0xf3, 0x8e, 0x64, 0x14, 0xf8, 0x12, 0xce,
	0x52, 0xd3, 0x57, 0x55, 0xb0, 0x65, 0x88, 0x5f, 0x06, 0x67, 0xc2, 0xbb, 0x74, 0x03, 0xff, 0xcb,
	0x3e, 0x13, 0x0f, 0xab, 0xf2, 0x54, 0xbe, 0x8a, 0xaf, 0xbc, 0xb1, 0xb3, 0xa4, 0xea, 0x7c, 0x7b,
	0xbe, 0x93, 0x64, 0x40, 0xdd, 0x22, 0xec, 0x20, 0x14, 0xf1, 0xa4, 0x77, 0x6d, 0x36, 0x75, 0x37,
	0xf2, 0x67, 0xa3, 0x5e, 0xd6, 0xc8, 0x1b, 0x53, 0x36, 0x47, 0xd9, 0xf9, 0x04, 0xd6, 0xcf, 0xd8,
	0xca, 0xff, 0xdb, 0x09, 0x9f, 0xe8, 0x5d, 0x42, 0x7e, 0xca, 0x7d, 0xe5, 0xd4, 0x1b, 0xa6, 0xf8,
	0x3b, 0xb2, 0x18, 0x1e, 0x1e, 0x55, 0x1f, 0x56, 0xe8, 0x77, 0xe8, 0xe3, 0x15, 0xfe, 0x99, 0xe6,
	0x7d, 0xd4, 0x5e, 0xff, 0x3e, 0xf6, 0x5e, 0x2d, 0x43, 0xdb, 0xac, 0x4f, 0x4f, 0x79, 0x7c, 0x1a,
	0xf4, 0x39, 0x79, 0x0c, 0xad, 0x27, 0x31, 0xf7, 0x04, 0xcf, 0x36, 0xc3, 0x05, 0xbb, 0x56, 0xa7,
	0xb3, 0x60, 0xff, 0x32, 0x3e, 0xee, 0x43, 0xf3, 0x90, 0x8b, 0xcc, 0xfc, 0x7a, 0x0e, 0x2d, 0xad,
	0x5f, 0x17, 0xb0, 0xb4, 0x18, 0x1f, 0x45, 0xa7, 0xfc, 0x32, 0x44, 0xd7, 0x0a, 0x1b, 0x41, 0x71,
	0xc3, 0x95, 0xd1, 0xe0, 0x2e, 0x74, 0xe5, 0x68, 0x1e, 0x40, 0x13, 0xf3, 0x81, 0x9b, 0x45, 0x79,
	0xf7, 0x98, 0xbb, 0x7a, 0x6e, 0xbd, 0xf9, 0x18, 0xec, 0x43, 0x2e, 0xd0, 0xea, 0x92, 0xae, 0xcf,
	0xd9, 0xef, 0x43, 0x13, 0x5d, 0x47, 0x8a, 0x9b, 0x25, 0xdc, 0xdc, 0x8a, 0x77, 0x2e, 0x0b, 0xa6,
	0xf1, 0x4d, 0x59, 0xe6, 0xd2, 0xf8, 0x05, 0xb4, 0x4c, 0x2c, 0x89, 0x5a, 0xff, 0xc8, 0x8d, 0xf2,
	0x85, 0x85, 0x85, 0xb9, 0x73, 0x73, 0xb1, 0x52, 0x93, 0xdd, 0x03, 0xc0, 0x8c, 0xaa, 0xbf, 0x61,
	0xa9, 0x33, 0x3b, 0x5b, 0xa5, 0x4e, 0x35, 0x56, 0x1f, 0xc2, 0xf2, 0x21, 0x17, 0xca, 0xe4, 0x9c,
	0x6c, 0xbe, 0xce, 0xfa, 0x81, 0x2a, 0x86, 0x1a, 0xc2, 0xa4, 0x84, 0x31, 0x53, 0xb9, 0xb3, 0x51,
	0x92, 0xab, 0xbf, 0xc0, 0x07, 0xb0, 0x86, 0x39, 0x52, 0x22, 0xf5, 0x90, 0x8b, 0x2d, 0x64, 0x46,
	0xc4, 0x62, 0xe3, 0x7b, 0x00, 0xb9, 0xf1, 0xa5, 0x23, 0x7d, 0x0c, 0xb0, 0xcf, 0x87, 0x5c, 0x5b,
	0x5d, 0xa1, 0xeb, 0x8f, 0xeb, 0x6a, 0x4c, 0xbe, 0xff, 0x4f, 0x00, 0x00, 0x00, 0xff, 0xff, 0xbb,
	0xe9, 0xeb, 0x65, 0x81, 0x0e, 0x00, 0x00,
}
