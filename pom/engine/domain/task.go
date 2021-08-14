package domain

type Task struct {
	ID          string
	Description string
	State       string
}

type ITaskService interface {
	Add(*Task) error
	Run(string) error
	Pause(string) error
	ShowTasks(interface{}) ([]*Task, error)
	DeleteTask(interface{}) (interface{}, error)
}

type ITaskRepo interface {
	Get(string) (*Task, error)
	Insert(*Task) error
	Update(*Task) error
	Delete(string) error
	GetAll() ([]*Task, error)
}

type ISerialize interface {
	Serialize(Task) []byte
	Deserialize([]byte) Task
}
