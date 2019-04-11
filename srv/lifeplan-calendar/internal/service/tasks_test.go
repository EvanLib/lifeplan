package service

import (
	"context"
	"testing"
	"time"

	"github.com/evanlib/lifeplan/srv/lifeplan-calendar/internal/config"
	"github.com/evanlib/lifeplan/srv/lifeplan-calendar/internal/database"
	tasks "github.com/evanlib/lifeplan/srv/lifeplan-calendar/proto"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

type TasksTestSuite struct {
	suite.Suite
	service *CalendarService
}

func Test_Tasks(t *testing.T) {
	suite.Run(t, new(TasksTestSuite))
}

func (suite *TasksTestSuite) SetupTest() {
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

func (suite *TasksTestSuite) TearDownTest() {
	if err := suite.service.db.Drop(); err != nil {
		suite.FailNow("Database deletion failed", "%v", err)
	}
	suite.service.db.Close()
}

func (suite *TasksTestSuite) TestCreateTask() {
	// create testing request
	rsp := &tasks.TaskResponse{}
	req := &tasks.Task{
		State:    tasks.TODO,
		Task:     "Clean room",
		Category: "Some Category",
	}
	err := suite.service.CreateTask(context.TODO(), req, rsp)

	assert.Nil(suite.T(), err)
	assert.NotNil(suite.T(), rsp.Task)
	assert.Equal(suite.T(), req.Task, rsp.Task.Task)
	assert.Equal(suite.T(), req.Category, rsp.Task.Category)
	assert.Equal(suite.T(), req.State, rsp.Task.State)
}

func (suite *TasksTestSuite) TestCreateSubTask() {
	// create testing request
	rsp := &tasks.TaskResponse{}
	req := &tasks.Task{
		State:      tasks.TODO,
		Task:       "Clean room",
		Category:   "Some Category",
		Categoryid: "EventID",
	}
	err := suite.service.CreateTask(context.TODO(), req, rsp)

	assert.Nil(suite.T(), err)
	assert.NotNil(suite.T(), rsp.Task)

	sub := &tasks.Task{
		State:      tasks.TODO,
		Task:       "Pick up clothes",
		Category:   "Some Category",
		Parent:     rsp.Task.Id,
		Categoryid: "EventID",
	}

	suite.service.CreateTask(context.TODO(), sub, rsp)
	assert.Nil(suite.T(), err)
	assert.NotNil(suite.T(), rsp.Task)

	sub = &tasks.Task{
		State:      tasks.TODO,
		Task:       "Pick up clothes 2",
		Category:   "Some Category",
		Parent:     rsp.Task.Id,
		Categoryid: "EventID",
	}

	suite.service.CreateTask(context.TODO(), sub, rsp)
	assert.Nil(suite.T(), err)
	assert.NotNil(suite.T(), rsp.Task)

	sub = &tasks.Task{
		State:      tasks.TODO,
		Task:       "Pick up clothes 3",
		Category:   "Some Category",
		Parent:     rsp.Task.Id,
		Categoryid: "EventID",
	}

	rsp3 := &tasks.TaskResponse{}
	suite.service.CreateTask(context.TODO(), sub, rsp3)
	assert.Nil(suite.T(), err)
	assert.NotNil(suite.T(), rsp.Task)

	sub = &tasks.Task{
		State:      tasks.TODO,
		Task:       "Pick up clothes 4",
		Category:   "Some Category",
		Parent:     rsp.Task.Id,
		Categoryid: "EventID",
	}

	rsp4 := &tasks.TaskResponse{}
	suite.service.CreateTask(context.TODO(), sub, rsp4)
	assert.Nil(suite.T(), err)
	assert.NotNil(suite.T(), rsp.Task)

	sub = &tasks.Task{
		State:      tasks.DONE,
		Task:       "Pick up clothes 5",
		Category:   "Some Category",
		Parent:     rsp4.Task.Id,
		Categoryid: "EventID",
	}

	rsp5 := &tasks.TaskResponse{}
	suite.service.CreateTask(context.TODO(), sub, rsp5)
	assert.Nil(suite.T(), err)
	assert.NotNil(suite.T(), rsp.Task)

	categoryTime := time.Now().UTC()
	tasksReq := &tasks.TasksRequest{
		Categoryid: "EventID",
	}

	tasksRsp := &tasks.TasksTree{}
	err = suite.service.GetTasks(context.TODO(), tasksReq, tasksRsp)
	m := make(map[string]int32)
	tasks.StateMap(tasksRsp.Nodes, m)

	// state map testing
	for i := range m {
		m[i] = tasks.DONE
	}
	updateReq := &tasks.TaskList{
		Categoryid:        "EventID",
		Categorytimestamp: categoryTime,
		TaskStateMap:      m,
	}
	mapRsp := &tasks.TasksTree{}
	err = suite.service.UpdateTasksState(context.TODO(), updateReq, mapRsp)

	//change some stuff to make it look cool :D
	m[rsp4.Task.Id] = tasks.MISSED
	m[rsp5.Task.Id] = tasks.MISSED
	m[rsp3.Task.Id] = tasks.MISSED
	updateReq = &tasks.TaskList{
		Categoryid:        "EventID",
		Categorytimestamp: categoryTime,
		TaskStateMap:      m,
	}

	mapRsp = &tasks.TasksTree{}
	err = suite.service.UpdateTasksState(context.TODO(), updateReq, mapRsp)
	assert.Nil(suite.T(), err)
}

func (suite *TasksTestSuite) TestDeleteTask() {
	// create testing request
	rsp := &tasks.TaskResponse{}
	req := &tasks.Task{
		State:      tasks.TODO,
		Task:       "DELETE TASK",
		Category:   "SOME CATEGORY",
		Categoryid: "EventID",
	}
	err := suite.service.CreateTask(context.TODO(), req, rsp)

	assert.Nil(suite.T(), err)
	assert.NotNil(suite.T(), rsp.Task)
	topid := rsp.Task.Id

	sub := &tasks.Task{
		State:      tasks.TODO,
		Task:       "Pick up clothes",
		Category:   "SOME CATEGORY",
		Parent:     topid,
		Categoryid: "EventID",
	}

	suite.service.CreateTask(context.TODO(), sub, rsp)
	assert.Nil(suite.T(), err)
	assert.NotNil(suite.T(), rsp.Task)

	sub = &tasks.Task{
		State:      tasks.TODO,
		Task:       "Pick up clothes 2",
		Category:   "SOME CATEGORY",
		Parent:     topid,
		Categoryid: "EventID",
	}

	suite.service.CreateTask(context.TODO(), sub, rsp)
	assert.Nil(suite.T(), err)
	assert.NotNil(suite.T(), rsp.Task)

	sub = &tasks.Task{
		State:      tasks.TODO,
		Task:       "Pick up clothes 3",
		Category:   "SOME CATEGORY",
		Parent:     topid,
		Categoryid: "EventID",
	}

	rsp3 := &tasks.TaskResponse{}
	suite.service.CreateTask(context.TODO(), sub, rsp3)
	assert.Nil(suite.T(), err)
	assert.NotNil(suite.T(), rsp.Task)

	reqByID := &tasks.FincByIdRequest{
		Id: topid,
	}
	suite.service.DeleteTask(context.TODO(), reqByID, nil)
}

func (suite *TasksTestSuite) TestUpdateTask() {
	// create testing request
	rsp := &tasks.TaskResponse{}
	req := &tasks.Task{
		State:      tasks.TODO,
		Task:       "UPDATE UPDATE THIS TASK",
		Category:   "UPDATE THIS CATEGORY",
		Categoryid: "EventID",
	}
	err := suite.service.CreateTask(context.TODO(), req, rsp)
	assert.Nil(suite.T(), err)
	assert.NotNil(suite.T(), rsp.Task)

	// update task
	updateRsp := &tasks.TaskResponse{}
	rsp.Task.Task = "THE TASK HAS BEE UPDATED"
	rsp.Task.Category = "THE CATEGORY IS UPDATED"

	err = suite.service.UpdateTask(context.TODO(), rsp.Task, updateRsp)
	assert.Nil(suite.T(), err)
	assert.NotNil(suite.T(), updateRsp.Task)
	assert.Equal(suite.T(), rsp.Task.Task, updateRsp.Task.Task)
	assert.Equal(suite.T(), rsp.Task.Category, updateRsp.Task.Category)
}
