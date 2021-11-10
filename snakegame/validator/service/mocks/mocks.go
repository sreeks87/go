package mocks

import (
	"github.com/stretchr/testify/mock"
	"github.com/sumup-challenges/backend-challenge-template-sreeks87/validator/domain"
)

type SnakeGameValidator struct {
	mock.Mock
}

type Generator struct {
	mock.Mock
}

func (v *SnakeGameValidator) Validate(*domain.ValidateRequest) (domain.State, error) {
	return domain.State{
		GameID: "123",
		Width:  5,
		Height: 5,
		Score:  0,
		Fruit: domain.Fruit{
			X: 1,
			Y: 1,
		},
		Snake: domain.Snake{
			X:    1,
			Y:    1,
			VelX: 1,
			VelY: 1,
		},
	}, nil
}
func (v *SnakeGameValidator) NewGame(int, int) (domain.State, error) {
	return domain.State{
		GameID: "123",
		Width:  5,
		Height: 5,
		Score:  0,
		Fruit: domain.Fruit{
			X: 1,
			Y: 1,
		},
		Snake: domain.Snake{
			X:    1,
			Y:    1,
			VelX: 1,
			VelY: 1,
		},
	}, nil
}

func (s *SnakeGameValidator) GetCurrentFruitPosFromDB(id string) (domain.Fruit, error) {
	return domain.Fruit{
		X: 1,
		Y: 1,
	}, nil
}

func (s *SnakeGameValidator) GetCurrentSnakePos([]domain.Tick) (*domain.Snake, error) {
	return &domain.Snake{
		X:    1,
		Y:    1,
		VelX: 1,
		VelY: 1,
	}, nil
}
func (s *SnakeGameValidator) GetLastSnakePos(id string) (*domain.Snake, error) {
	return &domain.Snake{
		X:    1,
		Y:    1,
		VelX: 1,
		VelY: 1,
	}, nil
}

func (s *SnakeGameValidator) Save(record domain.Record) (domain.State, error) {
	return domain.State{
		GameID: "123",
		Width:  5,
		Height: 5,
		Score:  0,
		Fruit: domain.Fruit{
			X: 1,
			Y: 1,
		},
		Snake: domain.Snake{
			X:    1,
			Y:    1,
			VelX: 1,
			VelY: 1,
		},
	}, nil
}

func (g *Generator) GenerateFruitPosition(width int, height int, snakeX int, snakeY int) (int, int, error) {
	return 1, 1, nil
}
func (g *Generator) GenerateID() int {
	return 1
}
