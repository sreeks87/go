package domain

type Student struct {
	ID        string `json:"id"`
	FirstName string `json:"firstname"`
	LastName  string `json:"lastname"`
	Course    string `json:"course"`
	Age       int    `json:"age"`
}

type Service interface {
	Validate(*Student) error
	Register(*Student) (*Student, error)
	Fetch(string) (*Student, error)
}

type Repository interface {
	Save(*Student) (*Student, error)
	Get(string) (*Student, error)
}
