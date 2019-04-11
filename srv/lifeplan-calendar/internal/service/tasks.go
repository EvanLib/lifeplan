package service

import (
	"context"
	"fmt"

	tasks "github.com/evanlib/lifeplan/srv/lifeplan-calendar/proto"
	"github.com/globalsign/mgo/bson"
)

func filter(ts []*tasks.Task, fn func(tasks *tasks.Task) bool) []*tasks.Task {
	f := []*tasks.Task{}
	for _, c := range ts {
		if fn(c) {
			f = append(f, c)
		}
	}
	return f
}

func recTaskTree(ts []*tasks.Task, treenode *tasks.TreeNode, parent string) {
	subtasks := filter(ts, func(task *tasks.Task) bool { return task.Parent == parent })
	for _, subtask := range subtasks {
		subnode := &tasks.TreeNode{Task: subtask}
		treenode.Subtasks = append(treenode.Subtasks, subnode)
		recTaskTree(ts, subnode, subtask.Id)
	}
}

// CreateTask inserts new task into data store from given request task.
// CreateTaskNode Top
func (ts *CalendarService) CreateTask(ctx context.Context, req *tasks.Task, rsp *tasks.TaskResponse) error {

	task := &tasks.Task{
		Id:         bson.NewObjectId().Hex(),
		Categoryid: req.Categoryid,
		State:      req.State,
		Task:       req.Task,
		Category:   req.Category,
		Parent:     req.Parent,
	}

	err := ts.db.Collection(CollectionTasks).Insert(task)
	if err != nil {
		return err
	}

	rsp.Task = task
	return nil
}

func (ts *CalendarService) GetTasks(ctx context.Context, req *tasks.TasksRequest, rsp *tasks.TasksTree) error {
	fmt.Println(req.Categorytimestamp)
	// get statemap
	if req.Categorytimestamp.IsZero() {
		catList := []*tasks.TaskList{}
		query := bson.M{"categoryid": req.Categoryid, "categorytimestamp": req.Categorytimestamp}
		err := ts.db.Collection(CollectionTasks).Find(query).All(&catList)
		if err != nil {
			return err
		}
	}

	// find all tasks on categoryid
	catTasks := []*tasks.Task{}
	query := bson.M{"categoryid": req.Categoryid}
	err := ts.db.Collection(CollectionTasks).Find(query).All(&catTasks)
	if err != nil {
		return err
	}

	// create tasks tree
	toptasks := filter(catTasks, func(task *tasks.Task) bool { return task.Parent == "" })
	for _, topTask := range toptasks {
		node := &tasks.TreeNode{Task: topTask}
		recTaskTree(catTasks, node, topTask.Id)
		rsp.Nodes = append(rsp.Nodes, node)
	}
	rsp.Categoryid = req.Categoryid

	// apply state map

	return nil
}

// UpdateTaskState inserts or updates TaskList in datastore.
// TaskStateMap map[taskid]state should only be applied to begining task parent.
func (ts *CalendarService) UpdateTaskState(ctx context.Context, req *tasks.TaskList, rsp *tasks.TaskResponse) error {
	// find tasklist for eventid/timestamp\
	selector := &tasks.TaskList{
		Categoryid:        req.Categoryid,
		Categorytimestamp: req.Categorytimestamp,
	}

	_, err := ts.db.Collection(CollectionTasksList).Upsert(selector, req)
	if err != nil {
		return err
	}

	return nil
}
