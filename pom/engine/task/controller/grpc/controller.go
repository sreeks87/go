package controller

import (
	"context"
	"pom/engine/domain"
	"pom/engine/task/controller/grpc/tasks_grpc"

	"google.golang.org/grpc"
)

type server struct {
	taskService domain.ITaskService
}

func NewTaskServiceGRPC(gserver *grpc.Server, svc domain.ITaskService) {
	taskServer := &server{
		taskService: svc,
	}

	tasks_grpc.RegisterTaskHandlerServer(gserver, taskServer)
}

func (s *server) transformTaskToGRPC(t *domain.Task) *tasks_grpc.SingleTask {
	s.taskService.Log(s.taskService.GetUser(), "Transforming task data to GRPC data")
	if t == nil {
		return nil
	}
	task := &tasks_grpc.SingleTask{
		ID:          t.ID,
		Description: t.Description,
		State:       t.State,
	}
	return task
}

func (s *server) transformGRPCToTask(t *tasks_grpc.SingleTask) *domain.Task {
	s.taskService.Log(s.taskService.GetUser(), "Transforming GRPC data to task data")
	task := &domain.Task{
		ID:          t.ID,
		Description: t.Description,
		State:       t.State,
	}
	return task
}

func (s *server) AddTask(c context.Context, in *tasks_grpc.SingleTask) (*tasks_grpc.SingleTask, error) {
	s.taskService.Log(s.taskService.GetUser(), "Add task controller")
	task := s.transformGRPCToTask(in)
	t, e := s.taskService.Add(task)
	return s.transformTaskToGRPC(t), e
}

func (s *server) FetchTask(c context.Context, in *tasks_grpc.TaskIDRequest) (*tasks_grpc.ListTasks, error) {
	s.taskService.Log(s.taskService.GetUser(), "Fetch task controller")
	var tasks []*domain.Task
	if in.Id == "all" {
		tasks, _ = s.taskService.ShowTasks(nil)
	} else {
		tasks, _ = s.taskService.ShowTasks(in.Id)
	}
	responseArray := make([]*tasks_grpc.SingleTask, len(tasks))
	for i, task := range tasks {
		responseArray[i] = s.transformTaskToGRPC(task)
	}
	result := &tasks_grpc.ListTasks{
		Tasks: responseArray,
	}
	return result, nil
}

func (s *server) StateUpdate(c context.Context, in *tasks_grpc.TaskIDRequest) (*tasks_grpc.SingleTask, error) {
	s.taskService.Log(s.taskService.GetUser(), "StateUpdate task controller")
	return nil, nil
}

func (s *server) DeleteTask(c context.Context, in *tasks_grpc.TaskIDRequest) (*tasks_grpc.SingleTask, error) {
	s.taskService.Log(s.taskService.GetUser(), "Delete task controller")
	t, _ := s.taskService.ShowTasks(in.Id)
	s.taskService.DeleteTask(in.Id)
	// todo: Dont do this, to be removed
	return s.transformTaskToGRPC(t[0]), nil
}
