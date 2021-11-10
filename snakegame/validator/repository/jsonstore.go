package repository

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"os"
	"regexp"
	"strconv"

	"github.com/schollz/jsonstore"
	"github.com/sumup-challenges/backend-challenge-template-sreeks87/validator/domain"
)

type jsonRepo struct {
	Path string
}

func NewJSONFileRepository(path string) domain.ValidatorRepo {
	return &jsonRepo{
		Path: path,
	}
}

func (j *jsonRepo) Get(id string) (*domain.Record, error) {
	js, err := jsonstore.Open(j.Path)
	if err != nil {
		panic(err)
	}

	var record domain.Record
	err = js.Get(id, &record)
	if _, ok := err.(*jsonstore.NoSuchKeyError); ok {
		return &record, nil
	}
	if err != nil {
		return nil, err
	}
	return &record, nil
}

func (j *jsonRepo) Insert(record *domain.Record) (*domain.Record, error) {
	m, _ := j.GetMaxID()
	records, e := j.GetAll()
	if e != nil {
		panic(e)
	}
	m += 1
	record.GameID = strconv.Itoa(m)
	records = append(records, record)
	// file write here
	fileData := make(map[string]string)
	for _, v := range records {
		str, _ := json.Marshal(v)
		fileData[v.GameID] = string(str)
	}
	fileData["maxid"] = strconv.Itoa(m)
	file, _ := json.MarshalIndent(fileData, "", " ")
	_ = ioutil.WriteFile(j.Path, file, 0644)
	// read the written file and update the maxid
	return record, nil
}

func (j *jsonRepo) Update(record *domain.Record) (*domain.Record, error) {
	existing, err := j.Get(record.GameID)
	if err != nil {
		return nil, err
	}
	existing.GameID = record.GameID
	existing.Fruit.X = record.Fruit.X
	existing.Fruit.Y = record.Fruit.Y
	existing.Snake.X = record.Snake.X
	existing.Snake.Y = record.Snake.Y
	rec, err := j.Insert(existing)
	if err != nil {
		return nil, err
	}
	return rec, nil
}

func (j *jsonRepo) GetMaxID() (int, error) {
	maxId := 0
	if !j.RepoExists() {
		return 0, nil
	}
	js, err := jsonstore.Open(j.Path)
	if err != nil {
		return 0, err
	}

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

func (j *jsonRepo) GetCurrentFruitPostFromDB(gameId string) (domain.Fruit, error) {
	rec, err := j.Get(gameId)
	if err != nil {
		return domain.Fruit{}, errors.New("game id not found")
	}
	return rec.Fruit, nil
}

func (j *jsonRepo) GetLastSnakePos(gameId string) (domain.Snake, error) {
	rec, err := j.Get(gameId)
	if err != nil {
		return domain.Snake{}, errors.New("game id not found")
	}
	return rec.Snake, nil
}

func (j *jsonRepo) GetAll() ([]*domain.Record, error) {
	if !j.RepoExists() {
		nullArray := make([]*domain.Record, 0)
		return nullArray, nil
	}
	js, err := jsonstore.Open(j.Path)
	if err != nil {
		return nil, err
	}

	var records []*domain.Record
	re := regexp.MustCompile("[0-9]+")
	allTasks := js.GetAll(re)
	for _, v := range allTasks {
		var record *domain.Record
		_ = json.Unmarshal(v, &record)
		records = append(records, record)
	}
	return records, nil
}
