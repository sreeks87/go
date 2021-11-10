package mocks

import (
	"github.com/stretchr/testify/mock"
	"github.com/sumup-challenges/backend-challenge-template-sreeks87/validator/domain"
)

type JSONRepo struct {
	mock.Mock
}

func (j *JSONRepo) Get(string) (*domain.Record, error) {
	return &domain.Record{
		GameID: "123",
		Width:  5,
		Height: 5,
		Score:  1,
		Fruit: domain.Fruit{
			X: 1,
			Y: 1,
		},
		Snake: domain.Snake{
			X:    2,
			Y:    3,
			VelX: 0,
			VelY: 1,
		},
		PreviousPosition: domain.Tick{
			VelX: 2,
			VelY: 2,
		},
	}, nil
}
func (j *JSONRepo) Insert(*domain.Record) (*domain.Record, error) {
	return &domain.Record{
		GameID: "123",
		Width:  5,
		Height: 5,
		Score:  1,
		Fruit: domain.Fruit{
			X: 1,
			Y: 1,
		},
		Snake: domain.Snake{
			X:    2,
			Y:    3,
			VelX: 0,
			VelY: 1,
		},
		PreviousPosition: domain.Tick{
			VelX: 2,
			VelY: 2,
		},
	}, nil
}
func (j *JSONRepo) Update(*domain.Record) (*domain.Record, error) {
	return &domain.Record{
		GameID: "123",
		Width:  5,
		Height: 5,
		Score:  1,
		Fruit: domain.Fruit{
			X: 1,
			Y: 1,
		},
		Snake: domain.Snake{
			X:    2,
			Y:    3,
			VelX: 0,
			VelY: 1,
		},
		PreviousPosition: domain.Tick{
			VelX: 2,
			VelY: 2,
		},
	}, nil
}
func (j *JSONRepo) GetCurrentFruitPostFromDB(string) (domain.Fruit, error) {
	return domain.Fruit{
		X: 1,
		Y: 1,
	}, nil
}
func (j *JSONRepo) GetLastSnakePos(string) (domain.Snake, error) {
	return domain.Snake{
		X:    2,
		Y:    3,
		VelX: 0,
		VelY: 1,
	}, nil
}

func (j *JSONRepo) GetMaxID() (int, error) {
	return 1, nil
}
func (j *JSONRepo) RepoExists() bool {
	return true
}

func (j *JSONRepo) GetAll() ([]*domain.Record, error) {
	var res []*domain.Record
	r := &domain.Record{
		GameID: "123",
		Width:  5,
		Height: 5,
		Score:  1,
		Fruit: domain.Fruit{
			X: 1,
			Y: 1,
		},
		Snake: domain.Snake{
			X:    2,
			Y:    3,
			VelX: 0,
			VelY: 1,
		},
		PreviousPosition: domain.Tick{
			VelX: 2,
			VelY: 2,
		},
	}
	res = append(res, r)
	return res, nil
}
