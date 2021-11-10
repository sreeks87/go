package service_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/sumup-challenges/backend-challenge-template-sreeks87/validator/domain"
	mockrepo "github.com/sumup-challenges/backend-challenge-template-sreeks87/validator/repository/mocks"
	"github.com/sumup-challenges/backend-challenge-template-sreeks87/validator/service"
	mocksvc "github.com/sumup-challenges/backend-challenge-template-sreeks87/validator/service/mocks"
)

func TestNewGame(t *testing.T) {
	mockRepo := new(mockrepo.JSONRepo)
	mockGen := new(mocksvc.Generator)
	s := service.NewSnakeGameValidator(mockRepo, mockGen)
	res, _ := s.NewGame(1, 1)
	assert.Equal(t, res.GameID, "123")
}

func TestValidate(t *testing.T) {
	mockRepo := new(mockrepo.JSONRepo)
	mockGen := new(mocksvc.Generator)
	s := service.NewSnakeGameValidator(mockRepo, mockGen)
	mockReq := &domain.ValidateRequest{
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
			VelX: 0,
			VelY: 1,
		},
	}
	res, _ := s.Validate(mockReq)
	assert.Equal(t, res.GameID, "123")
}
