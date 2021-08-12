package repository

import (
	"encoding/json"
	"pom/engine/domain"

	"github.com/schollz/jsonstore"
)

type jsonRepo struct {
	Path string
}

func NewFileRepository(p string) domain.ITaskRepo {
	return &jsonRepo{
		Path: p,
	}
}

func (j *jsonRepo) Get(id string) (*domain.Task, error) {
	js, err := jsonstore.Open(j.Path)
	if err != nil {
		panic(err)
	}
	var task domain.Task
	err = js.Get(id, &task)
	if err != nil {
		return nil, err
	}
	return &task, nil
}
func (j *jsonRepo) Insert(task *domain.Task) error {
	js, err := jsonstore.Open(j.Path)
	if err != nil {
		panic(err)
	}
	if err := js.Set(task.ID, &task); err != nil {
		panic(err)
	}
	if err := jsonstore.Save(js, j.Path); err != nil {
		return err
	}
	return nil
}
func (j *jsonRepo) Update(task *domain.Task) error {
	existingTask, err := j.Get(task.ID)
	if err != nil {
		return err
	}
	existingTask.ID = task.ID
	existingTask.Description = task.Description
	existingTask.State = task.State
	if err := j.Insert(existingTask); err != nil {
		return err
	}
	return nil
}

func (j *jsonRepo) GetAll() (*[]domain.Task, error) {
	js, err := jsonstore.Open(j.Path)
	if err != nil {
		return nil, err
	}
	var task domain.Task
	var tasks []domain.Task
	allTasks := js.GetAll(nil)
	for _, v := range allTasks {
		_ = json.Unmarshal(v, &task)
		tasks = append(tasks, task)
	}
	return &tasks, nil
}
