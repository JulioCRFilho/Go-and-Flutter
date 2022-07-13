package model

type Task struct {
	Id          int    `bson:"_id"`
	Name        string `json:"name" binding:"required"`
	done        bool   `default:"false"`
	DueDate     string `json:"dueDate" binding:"required"`
	CreatedDate string `json:"createdDate" binding:"required"`
}

func (t *Task) GetDone() bool {
	return t.done
}

func (t *Task) SetDone(b bool) {
	t.done = b
}
