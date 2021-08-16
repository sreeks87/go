package domain

type Task struct {
	ID          string
	Description string
	State       string
}

type ITaskService interface {
	Add(*Task) (*Task, error)
	Run(string) error
	Pause(string) error
	ShowTasks(interface{}) ([]*Task, error)
	DeleteTask(interface{}) (interface{}, error)
	Log(string, string)
	GetUser() string
}

type ITaskRepo interface {
	Get(string) (*Task, error)
	Insert(*Task) (*Task, error)
	Update(*Task) (*Task, error)
	Delete(string) error
	GetAll() ([]*Task, error)
}

type ISerialize interface {
	Serialize(Task) []byte
	Deserialize([]byte) Task
}
