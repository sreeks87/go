package database

import "github.com/sreeks87/repository/query/domain"

type KVMap struct {
	DB map[string]int
}

func NewKVMap(m map[string]int) domain.DB {
	return &KVMap{
		DB: m,
	}
}

func (db *KVMap) GetStat() (*domain.StatResponse, error) {
	return nil, nil
}
