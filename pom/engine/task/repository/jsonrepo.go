package repository

import (
	"encoding/json"
	"io/ioutil"
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
func (j *jsonRepo) Insert2(task *domain.Task) error {
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

func (j *jsonRepo) Update(task *domain.Task) (*domain.Task, error) {
	existingTask, err := j.Get(task.ID)
	if err != nil {
		return nil, err
	}
	existingTask.ID = task.ID
	existingTask.Description = task.Description
	existingTask.State = task.State
	tsk, err := j.Insert(existingTask)
	if err != nil {
		return nil, err
	}
	return tsk, nil
}

func (j *jsonRepo) GetAll() ([]*domain.Task, error) {
	if !j.RepoExists() {
		nullArray := make([]*domain.Task, 0)
		return nullArray, nil
	}
	js, err := jsonstore.Open(j.Path)
	if err != nil {
		return nil, err
	}

	var tasks []*domain.Task
	re := regexp.MustCompile("[0-9]+")
	allTasks := js.GetAll(re)
	for _, v := range allTasks {
		var task *domain.Task
		_ = json.Unmarshal(v, &task)
		tasks = append(tasks, task)
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

func (j *jsonRepo) Insert(task *domain.Task) (*domain.Task, error) {
	// js := new(jsonstore.JSONStore)
	m, _ := j.GetMaxID()
	tasks, e := j.GetAll()
	if e != nil {
		panic(e)
	}
	m += 1
	task.ID = strconv.Itoa(m)
	tasks = append(tasks, task)
	// file write here
	fileData := make(map[string]string)
	for _, v := range tasks {
		str, _ := json.Marshal(v)
		fileData[v.ID] = string(str)
	}
	fileData["maxid"] = strconv.Itoa(m)
	file, _ := json.MarshalIndent(fileData, "", " ")
	_ = ioutil.WriteFile(j.Path, file, 0644)
	// read the written file and update the maxid
	return task, nil
}
