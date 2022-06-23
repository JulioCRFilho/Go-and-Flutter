package model

type Task struct {
	Name        string `json:"name"`
	done        bool
	DueDate     string `json:"dueDate"`
	CreatedDate string `json:"createdDate"`
}

type TaskI interface {
	getDone() bool
	setDone(b bool)
}

func (t Task) getDone() bool {
	return t.done
}

func (t Task) setDone(b bool) {
	t.done = b
}
