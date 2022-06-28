package model

type Task struct {
	Id          int    `json:"id"`
	Name        string `json:"name"`
	done        bool
	DueDate     string `json:"dueDate"`
	CreatedDate string `json:"createdDate"`
}

func (t *Task) GetDone() bool {
	return t.done
}

func (t *Task) SetDone(b bool) {
	t.done = b
}
