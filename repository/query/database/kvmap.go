package database

import (
	"strconv"

	"github.com/sreeks87/repository/query/domain"
)

type KVMap struct {
	DB map[string]int
}

func NewKVMap(m map[string]int) domain.DB {
	return &KVMap{
		DB: m,
	}
}

func (db *KVMap) GetStat() (*domain.StatResponse, error) {
	var stat domain.Statistics
	var statList []domain.Statistics
	var statResp domain.StatResponse
	for k, v := range db.DB {
		stat.Query = k
		stat.Count = strconv.Itoa(v)
		statList = append(statList, stat)
	}
	statResp.Items = statList
	return &statResp, nil
}

func (db *KVMap) UpdateStat(param string) error {
	if _, Ok := db.DB[param]; Ok {
		db.DB[param] += 1
	}
	db.DB[param] = 1
	return nil
}
