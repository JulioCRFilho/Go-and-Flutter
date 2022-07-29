package model

import "go.mongodb.org/mongo-driver/bson/primitive"

type Task struct {
	Id          primitive.ObjectID `bson:"_id"`
	UserId      primitive.ObjectID `bson:"user_id"`
	Name        string             `bson:"name" binding:"required"`
	done        bool               `default:"false"`
	DueDate     string             `bson:"dueDate" binding:"required"`
	CreatedDate string             `bson:"createdDate" binding:"required"`
}

func (t *Task) Done() bool {
	return t.done
}

func (t *Task) SetDone(b bool) {
	t.done = b
}
