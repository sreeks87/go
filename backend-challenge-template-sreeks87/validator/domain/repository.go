package domain

type Record struct {
	GameID           string
	Width            int
	Height           int
	Score            int
	Fruit            Fruit
	Snake            Snake
	PreviousPosition Tick
}

type ValidatorRepo interface {
	Get(string) (*Record, error)
	Insert(*Record) (*Record, error)
	Update(*Record) (*Record, error)
	GetCurrentFruitPostFromDB(string) (Fruit, error)
	GetLastSnakePos(string) (Snake, error)
}
