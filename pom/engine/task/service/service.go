package service

import (
	"errors"
	"pom/engine/domain"
)

type taskService struct {
	taskRepo domain.ITaskRepo
	client   domain.ILogger
	user     string
}

func NewTaskService(taskRepository domain.ITaskRepo, c domain.ILogger, user string) domain.ITaskService {
	return &taskService{
		taskRepo: taskRepository,
		client:   c,
		user:     user,
	}
}

func (t *taskService) Add(task *domain.Task) (*domain.Task, error) {
	t.Log(t.user, "Add task called")
	tsk, err := t.taskRepo.Insert(task)
	if err != nil {
		return nil, err
	}
	return tsk, nil
}
func (t *taskService) Run(id string) error {
	t.Log(t.user, "Run task called")
	return nil
}
func (t *taskService) Pause(id string) error {
	t.Log(t.user, "Pause task called")
	return nil
}
func (t *taskService) ShowTasks(id interface{}) ([]*domain.Task, error) {
	t.Log(t.user, "Show task called")
	var tasks []*domain.Task
	if id != nil {
		t.client.Log(t.user, "Show all task")
		taskId := id.(string)
		task, err := t.taskRepo.Get(taskId)
		if err != nil {
			return nil, err
		}
		tasks = append(tasks, task)
	} else {
		tasks, _ = t.taskRepo.GetAll()
	}
	return tasks, nil
}

func (t *taskService) DeleteTask(id interface{}) (interface{}, error) {
	t.Log(t.user, "Delete task called")
	if e := t.taskRepo.Delete(id.(string)); e != nil {
		return nil, errors.New("delete failed")
	}
	return "Deleted Task " + id.(string), nil
}

func (t *taskService) Log(user string, data string) {
	t.client.Log(t.user, data)
}

func (t *taskService) GetUser() string {
	return t.user
}
