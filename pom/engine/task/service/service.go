package service

import (
	"errors"
	"pom/engine/domain"
)

type taskService struct {
	taskRepo domain.ITaskRepo
}

func NewTaskService(taskRepository domain.ITaskRepo) domain.ITaskService {
	return &taskService{
		taskRepo: taskRepository,
	}
}

func (t *taskService) Add(task *domain.Task) error {
	if err := t.taskRepo.Insert(task); err != nil {
		return err
	}
	return nil
}
func (t *taskService) Run(id string) error {
	return nil
}
func (t *taskService) Pause(id string) error {
	return nil
}
func (t *taskService) ShowTasks(id interface{}) ([]*domain.Task, error) {
	var tasks []*domain.Task
	if id != nil {
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
	if e := t.taskRepo.Delete(id.(string)); e != nil {
		return nil, errors.New("Delete failed.")
	}
	return "Deleted Task " + id.(string), nil
}
