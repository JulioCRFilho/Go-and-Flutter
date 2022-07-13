package model

import "go.mongodb.org/mongo-driver/bson/primitive"

type Task struct {
	Id          primitive.ObjectID `bson:"_id"`
	Name        string             `json:"name" binding:"required"`
	done        bool               `default:"false"`
	DueDate     string             `json:"dueDate" binding:"required"`
	CreatedDate string             `json:"createdDate" binding:"required"`
}

func (t *Task) Done() bool {
	return t.done
}

func (t *Task) SetDone(b bool) {
	t.done = b
}
