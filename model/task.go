package model

import "go.mongodb.org/mongo-driver/bson/primitive"

type Task struct {
	Id          primitive.ObjectID `bson:"_id"`
	UserId      primitive.ObjectID `bson:"user_id"`
	Name        string             `bson:"name" binding:"required"`
	Done        bool               `bson:"done" default:"false"`
	DueDate     string             `bson:"dueDate" binding:"required"`
	CreatedDate string             `bson:"createdDate" binding:"required"`
}
