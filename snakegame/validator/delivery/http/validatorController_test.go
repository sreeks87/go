package delivery_test

import (
	"bytes"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
	delivery "github.com/sumup-challenges/backend-challenge-template-sreeks87/validator/delivery/http"
	"github.com/sumup-challenges/backend-challenge-template-sreeks87/validator/service/mocks"
)

func TestGetNewGameInvalidWandH(t *testing.T) {
	mockSvc := new(mocks.SnakeGameValidator)
	c := delivery.ValidatorController{ValidatorSvc: mockSvc}
	req, err := http.NewRequest("GET", "/new", nil)
	if err != nil {
		t.Fatal(err)
	}
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(c.GetNewGame)
	handler.ServeHTTP(rr, req)
	// Check the status code is what we expect.
	assert.Equal(t, rr.Code, 400)
}

func TestGetNewGameInvalidWidth(t *testing.T) {
	mockSvc := new(mocks.SnakeGameValidator)
	c := delivery.ValidatorController{ValidatorSvc: mockSvc}
	req, err := http.NewRequest("GET", "/new?w=sddd&h=4", nil)
	if err != nil {
		t.Fatal(err)
	}
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(c.GetNewGame)
	handler.ServeHTTP(rr, req)
	// Check the status code is what we expect.
	assert.Equal(t, rr.Code, 400)
}

func TestGetNewGameInvalidHeight(t *testing.T) {
	mockSvc := new(mocks.SnakeGameValidator)
	c := delivery.ValidatorController{ValidatorSvc: mockSvc}
	req, err := http.NewRequest("GET", "/new?w=4&h=fjvnfjn", nil)
	if err != nil {
		t.Fatal(err)
	}
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(c.GetNewGame)
	handler.ServeHTTP(rr, req)
	// Check the status code is what we expect.
	assert.Equal(t, rr.Code, 400)
}

func TestGetNewGameValid(t *testing.T) {
	mockSvc := new(mocks.SnakeGameValidator)
	c := delivery.ValidatorController{ValidatorSvc: mockSvc}
	req, err := http.NewRequest("GET", "/new?w=4&h=4", nil)
	if err != nil {
		t.Fatal(err)
	}
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(c.GetNewGame)
	handler.ServeHTTP(rr, req)
	// Check the status code is what we expect.
	assert.Equal(t, rr.Code, 200)
}

func TestValidate(t *testing.T) {
	mockSvc := new(mocks.SnakeGameValidator)
	c := delivery.ValidatorController{ValidatorSvc: mockSvc}
	r := []byte(
		`{"gameId": "1", 
		"width": 5, 
		"height": 5, 
		"score": 0, 
		"fruit": {"x": 4, "y": 4}, 
		"snake": {"x": 0, "y": 0, "velX": 1, "velY": 0},
		"ticks":[
		{"velX":1,"velY":0},
		{"velX":1,"velY":0},
		{"velX":1,"velY":0},
		{"velX":0,"velY":1},
		{"velX":0,"velY":1},
		{"velX":0,"velY":1},
		{"velX":0,"velY":1},
		{"velX":0,"velY":1},
		]
		}`)
	req, err := http.NewRequest("POST", "/validate", bytes.NewBuffer(r))
	if err != nil {
		t.Fatal(err)
	}
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(c.GetNewGame)
	handler.ServeHTTP(rr, req)
	// Check the status code is what we expect.
	assert.Equal(t, rr.Code, 400)
}
