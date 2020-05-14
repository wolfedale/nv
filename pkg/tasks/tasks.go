package tasks

type TaskInFace interface {
	Create() error
	Edit() error
	Delete() error
}

type Tasks struct {
	Task []Task
}

type Task struct {
	Name string
}

func New() (*Tasks, error) {
	return &Tasks{}, nil
}

func (t *Tasks) Create() error {
	return nil
}

func (t *Tasks) Edit() error {
	return nil
}

func (t *Tasks) Delete() error {
	return nil
}
