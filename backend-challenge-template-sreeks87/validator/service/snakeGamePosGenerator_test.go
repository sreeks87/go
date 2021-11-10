package service_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/sumup-challenges/backend-challenge-template-sreeks87/validator/domain"
	"github.com/sumup-challenges/backend-challenge-template-sreeks87/validator/service"
)

func TestGenerateFruitPositionERROR(t *testing.T) {
	s := service.NewFoodPositionService()
	_, _, e := s.GenerateFruitPosition(0, 0, 1, 1)
	assert.Equal(t, e.(domain.ValidationError).ErrCode, 400)
}
func TestGenerateFruitPosition(t *testing.T) {
	s := service.NewFoodPositionService()
	_, _, e := s.GenerateFruitPosition(5, 5, 1, 1)
	assert.Equal(t, e, nil)
}
func TestGenerateFruitPosition2(t *testing.T) {
	s := service.NewFoodPositionService()
	v := s.GenerateID()
	assert.LessOrEqual(t, v, 500)
}
