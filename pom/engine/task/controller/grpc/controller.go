package controller

import (
	"context"
	"pom/engine/domain"
	"pom/engine/task/controller/grpc/tasks_grpc"

	"github.com/golang/protobuf/protoc-gen-go/grpc"
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

func (s *server) transformTaskRequestToGRPC(t *domain.Task) *tasks_grpc.SingleTask {
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

func (s *server) transformGRPCResponseToResponse(t *tasks_grpc.SingleTask) *domain.Task {
	task := &domain.Task{
		ID:          t.ID,
		Description: t.Description,
		State:       t.State,
	}
	return task
}

func (s *server) AddTask(c context.Context, in *tasks_grpc.SingleTask) (*tasks_grpc.TaskWithErrorResponse, error) {

	return nil, nil
}
func (s *server) FetchTask(in *tasks_grpc.TaskIDRequest, stream tasks_grpc.TaskHandler_FetchTaskServer) error {
	return nil
}

// func (s *server) FetchTask(c context.Context, in *tasks_grpc.TaskIDRequest) (tasks_grpc.TaskHandler_FetchTaskServer, error) {
// return nil, nil
// }

func (s *server) StateUpdate(c context.Context, in *tasks_grpc.TaskIDRequest) (*tasks_grpc.TaskWithErrorResponse, error) {
	return nil, nil
}
func (s *server) DeleteTask(c context.Context, in *tasks_grpc.TaskIDRequest) (*tasks_grpc.TaskWithErrorResponse, error) {
	return nil, nil
}
