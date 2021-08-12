package domain

type Task struct {
	ID          string `json:"id"`
	Description string `json:"description"`
	State       string `json:"state"`
}

type ITaskService interface {
	Add(*Task) error
	Run(string) error
	Pause(string) error
	ShowTasks(interface{}) (*[]Task, error)
}

type ITaskRepo interface {
	Get(string) (*Task, error)
	Insert(*Task) error
	Update(*Task) error
	GetAll() (*[]Task, error)
}

type ISerialize interface {
	Serialize(Task) []byte
	Deserialize([]byte) Task
}
