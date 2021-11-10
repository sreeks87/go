package repository_test

import (
	"io/ioutil"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/sumup-challenges/backend-challenge-template-sreeks87/validator/domain"
	"github.com/sumup-challenges/backend-challenge-template-sreeks87/validator/repository"
)

const FILE_PATH = "mockDB.json"

func setup() {
	b := []byte(`{
		"1": "{\"GameID\":\"1\",\"Width\":5,\"Height\":5,\"Score\":0,\"Fruit\":{\"x\":0,\"y\":2},\"Snake\":{\"x\":0,\"y\":0,\"velX\":1,\"velY\":0},\"PreviousPosition\":{\"velX\":0,\"velY\":0}}",
		"2": "{\"GameID\":\"2\",\"Width\":100,\"Height\":100,\"Score\":0,\"Fruit\":{\"x\":85,\"y\":20},\"Snake\":{\"x\":0,\"y\":0,\"velX\":1,\"velY\":0},\"PreviousPosition\":{\"velX\":0,\"velY\":0}}",
		"maxid": "2"
	   }`)
	err := ioutil.WriteFile(FILE_PATH, b, 0644)
	if err != nil {
		panic(err)
	}
}

func teardown() {
	err := os.Remove(FILE_PATH)
	if err != nil {
		panic(err)
	}
}
func TestGet(t *testing.T) {
	setup()
	defer teardown()
	d := repository.NewJSONFileRepository(FILE_PATH)
	res, _ := d.Get("1")
	assert.Equal(t, res.GameID, "1")
}

func TestInsert(t *testing.T) {
	setup()
	defer teardown()
	d := repository.NewJSONFileRepository(FILE_PATH)
	r := domain.Record{
		GameID: "123",
		Width:  6,
		Height: 6,
		Score:  1,
		Fruit: domain.Fruit{
			X: 2,
			Y: 3,
		}, Snake: domain.Snake{
			X:    1,
			Y:    4,
			VelX: 4,
			VelY: 3,
		},
	}
	res, _ := d.Insert(&r)
	assert.Equal(t, res.Width, 6)
}

func TestUpdate(t *testing.T) {
	setup()
	defer teardown()
	d := repository.NewJSONFileRepository(FILE_PATH)
	r := domain.Record{
		GameID: "1",
		Width:  6,
		Height: 6,
		Score:  1,
		Fruit: domain.Fruit{
			X: 2,
			Y: 3,
		}, Snake: domain.Snake{
			X:    1,
			Y:    4,
			VelX: 4,
			VelY: 3,
		},
	}
	res, _ := d.Update(&r)
	assert.Equal(t, res.Width, 5)
}

func TestGetCurrentFruitPostFromDB(t *testing.T) {
	setup()
	defer teardown()
	d := repository.NewJSONFileRepository(FILE_PATH)
	res, _ := d.GetCurrentFruitPostFromDB("1")
	assert.Equal(t, res.X, 0)
}

func TestGetLastSnakePos(t *testing.T) {
	setup()
	defer teardown()
	d := repository.NewJSONFileRepository(FILE_PATH)
	res, _ := d.GetLastSnakePos("1")
	assert.Equal(t, res.X, 0)
}
