package repository

import (
	"errors"

	"github.com/sreeks87/school/student/domain"
)

type KVMap struct {
	DB map[string]domain.Student
}

func NewKVMap(m map[string]domain.Student) domain.Repository {
	return &KVMap{
		DB: m,
	}
}

func (k *KVMap) Save(std *domain.Student) (*domain.Student, error) {
	if _, ok := k.DB[std.ID]; ok {
		return nil, errors.New("user already exists")
	}
	k.DB[std.ID] = *std

	return std, nil
}

func (k *KVMap) Get(id string) (*domain.Student, error) {
	if _, ok := k.DB[id]; !ok {
		return nil, errors.New("user not found")
	}
	std := k.DB[id]
	return &std, nil
}
