package repository

import (
	"encoding/json"
	"os"
	"pom/engine/domain"
	"regexp"
	"strconv"

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
	js := new(jsonstore.JSONStore)
	m, _ := j.GetMaxID()
	m += 1
	task.ID = strconv.Itoa(m)
	if err := js.Set(task.ID, &task); err != nil {
		panic(err)
	}
	if err := jsonstore.Save(js, j.Path); err != nil {
		return err
	}

	if err := js.Set("maxid", m); err != nil {
		panic(err)
	}
	if err := jsonstore.Save(js, j.Path); err != nil {
		return err
	}
	return nil
}

func (j *jsonRepo) Delete(taskId string) error {
	js, err := jsonstore.Open(j.Path)
	if err != nil {
		panic(err)
	}
	js.Delete(taskId)
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

func (j *jsonRepo) GetAll() ([]*domain.Task, error) {
	js, err := jsonstore.Open(j.Path)
	if err != nil {
		return nil, err
	}
	var task domain.Task
	var tasks []*domain.Task
	re := regexp.MustCompile("[0-9]+")
	allTasks := js.GetAll(re)
	for _, v := range allTasks {
		_ = json.Unmarshal(v, &task)
		tasks = append(tasks, &task)
	}
	return tasks, nil
}

func (j *jsonRepo) GetMaxID() (int, error) {
	if !j.RepoExists() {
		return 0, nil
	}
	js, err := jsonstore.Open(j.Path)
	if err != nil {
		panic(err)
	}
	var maxId int
	err = js.Get("maxid", &maxId)
	if err != nil {
		return 0, err
	}
	return maxId, nil
}

func (j *jsonRepo) RepoExists() bool {
	if _, err := os.Stat(j.Path); err != nil {
		if os.IsNotExist(err) {
			return false
		}
	}
	return true
}
