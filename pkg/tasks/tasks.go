package tasks

type Tasks struct{}

func New() (*Tasks, error) {
	return &Tasks{}, nil
}
