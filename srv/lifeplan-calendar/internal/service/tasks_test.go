package service

import (
	"context"
	"fmt"
	"testing"

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
	// test data
	task := "Clean room"
	category := "Some Category"

	rsp := &tasks.TaskResponse{}

	// create testing request
	req := &tasks.Task{
		State:    tasks.TODO,
		Task:     task,
		Category: category,
	}
	err := suite.service.CreateTask(context.TODO(), req, rsp)

	assert.Nil(suite.T(), err)
	assert.NotNil(suite.T(), rsp.Task)
	assert.Equal(suite.T(), req.Task, rsp.Task.Task)
	assert.Equal(suite.T(), req.Category, rsp.Task.Category)
	assert.Equal(suite.T(), req.State, rsp.Task.State)
}

func (suite *TasksTestSuite) TestCreateSubTask() {
	// test data
	task := "Clean room"
	category := "Some Category"

	rsp := &tasks.TaskResponse{}

	// create testing request
	req := &tasks.Task{
		State:      tasks.TODO,
		Task:       task,
		Category:   category,
		Categoryid: "EventID",
	}
	err := suite.service.CreateTask(context.TODO(), req, rsp)

	assert.Nil(suite.T(), err)
	assert.NotNil(suite.T(), rsp.Task)

	sub := &tasks.Task{
		State:      tasks.TODO,
		Task:       "Pick up clothes",
		Category:   category,
		Parent:     rsp.Task.Id,
		Categoryid: "EventID",
	}

	suite.service.CreateTask(context.TODO(), sub, rsp)
	assert.Nil(suite.T(), err)
	assert.NotNil(suite.T(), rsp.Task)

	sub = &tasks.Task{
		State:      tasks.TODO,
		Task:       "Pick up clothes 2",
		Category:   category,
		Parent:     rsp.Task.Id,
		Categoryid: "EventID",
	}

	suite.service.CreateTask(context.TODO(), sub, rsp)
	assert.Nil(suite.T(), err)
	assert.NotNil(suite.T(), rsp.Task)

	sub = &tasks.Task{
		State:      tasks.TODO,
		Task:       "Pick up clothes 3",
		Category:   category,
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
		Category:   category,
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
		Category:   category,
		Parent:     rsp4.Task.Id,
		Categoryid: "EventID",
	}

	rsp5 := &tasks.TaskResponse{}
	suite.service.CreateTask(context.TODO(), sub, rsp5)
	assert.Nil(suite.T(), err)
	assert.NotNil(suite.T(), rsp.Task)

	tasksReq := &tasks.TasksRequest{
		Categoryid: "EventID",
	}

	tasksRsp := &tasks.TasksTree{}
	err = suite.service.GetTasks(context.TODO(), tasksReq, tasksRsp)
	printRec(tasksRsp.Nodes)
	m := make(map[string]int32)
	tasks.StateMap(tasksRsp.Nodes, m)
	fmt.Println(m)
}

func printRec(taskNodes []*tasks.TreeNode) {
	for _, treeNode := range taskNodes {
		fmt.Println(len(treeNode.Subtasks))
		fmt.Println(treeNode.Task)
		if len(treeNode.Subtasks) > 0 {
			printRec(treeNode.Subtasks)
		}
	}
}
