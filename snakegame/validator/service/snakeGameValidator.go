package service

import (
	"math"

	"github.com/sumup-challenges/backend-challenge-template-sreeks87/validator/domain"
)

// The snake validator strcut that holds the repo object and the generator object
type snakeGameValidator struct {
	snakeValidatorRepo domain.ValidatorRepo
	generator          domain.Generator
}

func NewSnakeGameValidator(repo domain.ValidatorRepo, gen domain.Generator) domain.Validator {
	return &snakeGameValidator{
		snakeValidatorRepo: repo,
		generator:          gen,
	}
}

func (s *snakeGameValidator) NewGame(width int, height int) (domain.State, error) {
	x, y, _ := s.generator.GenerateFruitPosition(width, height, -1, -1)
	// Disabling the DB interaction for now.
	// Puposefully keeping the commented code for future discussion

	record := domain.Record{
		GameID: "0",
		Width:  width,
		Height: height,
		Score:  0,
		Fruit: domain.Fruit{
			X: x,
			Y: y,
		},
		Snake: domain.Snake{
			X:    0,
			Y:    0,
			VelX: 1,
			VelY: 0,
		},
		PreviousPosition: domain.Tick{VelX: 0, VelY: 0},
	}
	state, e := s.Save(record)
	if e != nil {
		return state, e
	}
	return state, nil
}
func (s *snakeGameValidator) Validate(request *domain.ValidateRequest) (domain.State, error) {
	// pass current position of the snake here, so that the new food position is not == snake postion

	// Disabling the DB interaction for now.
	// Puposefully keeping the commented code for future discussion
	// existingRecord, e := s.snakeValidatorRepo.Get(request.GameID)
	// if e != nil {
	// 	return domain.State{}, e
	// }

	// currentFruitPos := request.Fruit
	// currentFruitPosFromDB := existingRecord.Fruit
	// if !(currentFruitPos == currentFruitPosFromDB) {
	// 	return domain.State{}, errors.New("fruit position doesnt match the previously generated position")
	// }
	var state domain.State
	newSnakePos, e := validateAndGetNewSnakePosition(request)
	if e != nil {
		// Generate state object here
		state = domain.State{
			GameID: request.GameID,
			Width:  request.Width,
			Height: request.Height,
			Score:  request.Score,
			Fruit: domain.Fruit{
				X: request.Fruit.X,
				Y: request.Fruit.Y,
			},
			Snake: domain.Snake{
				X:    newSnakePos.X,
				Y:    newSnakePos.Y,
				VelX: request.Snake.VelX,
				VelY: request.Snake.VelY,
			},
		}
		return state, e
	}
	x, y, _ := s.generator.GenerateFruitPosition(request.Width, request.Height, newSnakePos.X, newSnakePos.Y)
	state = domain.State{
		GameID: request.GameID,
		Width:  request.Width,
		Height: request.Height,
		Score:  request.Score + 1,
		Fruit: domain.Fruit{
			X: x,
			Y: y,
		},
		Snake: domain.Snake{
			X:    newSnakePos.X,
			Y:    newSnakePos.Y,
			VelX: 0,
			VelY: 1,
		},
	}
	return state, nil
}

func validateAndGetNewSnakePosition(request *domain.ValidateRequest) (domain.Snake, error) {
	snakeX := request.Snake.X
	snakeY := request.Snake.Y
	fruitX := request.Fruit.X
	fruitY := request.Fruit.Y

	prevTick := domain.Tick{
		VelX: 0,
		VelY: 0,
	}
	for _, tick := range request.Ticks {
		// diagonal move
		if math.Abs(float64(tick.VelX))*math.Abs(float64(tick.VelY)) == 1 {
			// invalid velocity return error
			// at this point service should not be aware of the http codes or
			// what mechanism is used to call service, but for simplicity im returning http code as error code
			return request.Snake, &domain.ValidationError{ErrCode: 418, ErrorDesc: "invalid diagonal velocity"}
		}
		// look for invalid velocity if !(-1,0,1) return error
		if math.Abs(float64(tick.VelX)) > 1 || math.Abs(float64(tick.VelY)) > 1 {
			// invalid velocity return error
			return request.Snake, &domain.ValidationError{ErrCode: 418, ErrorDesc: "invalid velocity"}
		}
		// look for 180 degree turns
		if (tick.VelX*prevTick.VelX < 0) || (tick.VelY*prevTick.VelY < 0) {
			// invalid turn return error
			return request.Snake, &domain.ValidationError{ErrCode: 418, ErrorDesc: "invalid turn"}
		}
		snakeX += tick.VelX
		snakeY += tick.VelY
		// game over scenario
		if snakeX > request.Width || snakeX < 0 || snakeY > request.Height || snakeY < 0 {
			return request.Snake, &domain.ValidationError{ErrCode: 418, ErrorDesc: "game over"}
		}
		prevTick = tick
	}

	if snakeX == fruitX && snakeY == fruitY {
		newSnakePos := domain.Snake{
			X:    snakeX,
			Y:    snakeY,
			VelX: request.Snake.VelX,
			VelY: request.Snake.VelY,
		}
		return newSnakePos, nil
	}

	return request.Snake, &domain.ValidationError{ErrCode: 404, ErrorDesc: "fruit not found"}

}

func (s *snakeGameValidator) Save(record domain.Record) (domain.State, error) {
	result, err := s.snakeValidatorRepo.Insert(&record)
	if err != nil {
		return domain.State{}, err
	}

	state := domain.State{
		GameID: result.GameID,
		Width:  result.Width,
		Height: result.Height,
		Score:  result.Score,
		Fruit:  result.Fruit,
		Snake:  result.Snake,
	}

	return state, nil

}

// The below functions are not really used.,
// butthese will be helper DB functions if we wish to
// validate the request vs state in the DB
// this helps the system to verify any manipulation to the request the clent might have done

func (s *snakeGameValidator) GetCurrentFruitPosFromDB(id string) (domain.Fruit, error) {
	s.snakeValidatorRepo.GetCurrentFruitPostFromDB(id)
	return domain.Fruit{}, nil
}

// func (s *snakeGameValidator) GetCurrentSnakePos([]domain.Tick) (*domain.Snake, error) {
// 	s.snakeValidatorRepo.GetLastSnakePos(id)
// 	return nil, nil
// }

func (s *snakeGameValidator) GetLastSnakePos(id string) (*domain.Snake, error) {
	s.snakeValidatorRepo.GetLastSnakePos(id)
	return &domain.Snake{}, nil
}
