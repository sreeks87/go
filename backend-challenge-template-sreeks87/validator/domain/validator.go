package domain

// i could have avoided duplicating struct with same fields
// state, Record and validateRequest are somewhat similar i could have composed the validateRequest
// with tick struct but since the request response is already defined and the test binary would send it in in this format
// i am going ahead with duplication over compsoition of stucts.

type State struct {
	GameID string `json:"gameId"`
	Width  int    `json:"width"`
	Height int    `json:"height"`
	Score  int    `json:"score"`
	Fruit  Fruit  `json:"fruit"`
	Snake  Snake  `json:"snake"`
}

type Fruit struct {
	X int `json:"x"`
	Y int `json:"y"`
}

type Snake struct {
	X    int `json:"x"`
	Y    int `json:"y"`
	VelX int `json:"velX"` // X velocity of the snake (-1, 0, 1)
	VelY int `json:"velY"` // Y velocity of the snake (-1, 0, 1)
}

type ValidateRequest struct {
	GameID string `json:"gameId"`
	Width  int    `json:"width"`
	Height int    `json:"height"`
	Score  int    `json:"score"`
	Fruit  Fruit  `json:"fruit"`
	Snake  Snake  `json:"snake"`
	Ticks  []Tick `json:"ticks"`
}

type Tick struct {
	VelX int `json:"velX"`
	VelY int `json:"velY"`
}

type Validator interface {
	Validate(*ValidateRequest) (State, error)
	NewGame(int, int) (State, error)
}
