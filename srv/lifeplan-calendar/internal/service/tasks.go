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

// GetTasks retrives tasks and sub tasks from datastore.
// Stores Tasks in RSP TaskTree.
// If request is supplied categoryid and categorytimestamp,
// applies statemap to RSP TaskTree.
func (ts *CalendarService) GetTasks(ctx context.Context, req *tasks.TasksRequest, rsp *tasks.TasksTree) error {
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
	// get statemap
	if !req.Categorytimestamp.IsZero() {
		catList := tasks.TaskList{}
		query := bson.M{"categoryid": req.Categoryid, "categorytimestamp": req.Categorytimestamp}
		err := ts.db.Collection(CollectionTasksList).Find(query).One(&catList)
		if err != nil {
			fmt.Println(err)
			return err
		}
		tasks.ApplyStateMap(rsp.Nodes, catList.TaskStateMap)
	}
	return nil
}

// UpdateTasksState inserts or updates TaskList in datastore.
// TaskStateMap map[taskid]state should only be applied to begining task parent.
func (ts *CalendarService) UpdateTasksState(ctx context.Context, req *tasks.TaskList, rsp *tasks.TasksTree) error {
	query := bson.M{"categoryid": req.Categoryid, "categorytimestamp": req.Categorytimestamp}
	_, err := ts.db.Collection(CollectionTasksList).Upsert(query, req)
	if err != nil {
		return err
	}

	return ts.GetTasks(ctx, &tasks.TasksRequest{Categoryid: req.Categoryid, Categorytimestamp: req.Categorytimestamp}, rsp)
}

// DeleteTask delete task and subtasks based on given taskID.
func (ts *CalendarService) DeleteTask(ctx context.Context, req *tasks.FincByIdRequest, rsp *tasks.EmptyResponse) error {
	// Delete task
	err := ts.db.Collection(CollectionTasks).RemoveId(req.Id)
	if err != nil {
		return err
	}
	// Delete subtask
	query := bson.M{"parent": req.Id}
	_, err = ts.db.Collection(CollectionTasks).RemoveAll(query)
	if err != nil {
		return err
	}

	return nil
}

// UpdateTask updates a given request Tasks returns GetTasks.
func (ts *CalendarService) UpdateTask(ctx context.Context, req *tasks.Task, rsp *tasks.TaskResponse) error {
	//update tasks
	err := ts.db.Collection(CollectionTasks).UpdateId(req.Id, req)
	if err != nil {
		return err
	}

	return ts.GetTask(ctx, &tasks.FincByIdRequest{Id: req.Id}, rsp)
}

// GetTask Retrieves a single task from data store and stores in rps.
func (ts *CalendarService) GetTask(ctx context.Context, req *tasks.FincByIdRequest, rsp *tasks.TaskResponse) error {
	//return task
	task := &tasks.Task{}
	err := ts.db.Collection(CollectionTasks).FindId(req.Id).One(task)
	if err != nil {
		return err
	}
	rsp.Task = task
	return nil
}
