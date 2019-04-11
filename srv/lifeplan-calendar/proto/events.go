package calendar

const (
	// SingleInstance update/delete single event
	SingleInstance = 1
	// AllInstances update/delete all events
	AllInstances = 2
	// FutureInstance update/delete future events
	FutureInstance = 3
	// STATECHANGE update state on tasks
	STATECHANGE = 1
	// OTHERCHANGE update other attributes on tasks
	OTHERCHANGE = 2
	// TODO state of task
	TODO = 1
	// INPROGRESS state of task
	INPROGRESS = 2
	// DONE state of task
	DONE = 3
	// MISSED state of task
	MISSED = 4
)
