package domain

type Generator interface {
	GenerateFruitPosition(int, int, int, int) (int, int, error)
	GenerateID() int
}
