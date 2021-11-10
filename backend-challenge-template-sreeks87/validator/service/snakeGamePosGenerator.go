package service

import (
	"math/rand"
	"time"

	"github.com/sumup-challenges/backend-challenge-template-sreeks87/validator/domain"
)

type generator struct {
}

func NewFoodPositionService() domain.Generator {
	return &generator{}
}
func (g *generator) GenerateFruitPosition(width int, height int, snakeX int, snakeY int) (int, int, error) {
	if width == 0 && height == 0 {
		return 0, 0, domain.ValidationError{ErrCode: 400, ErrorDesc: "cant create snake board for 0/0"}
	}
	// it should not be equal to the curret position of the snake
findposition:
	rand.Seed(time.Now().UnixNano())
	x := rand.Intn(width)
	y := rand.Intn(height)
	if x == snakeX && y == snakeY {
		goto findposition
	}
	return x, y, nil
}

func (g *generator) GenerateID() int {
	rand.Seed(time.Now().UnixNano())
	return rand.Intn(500)
}
