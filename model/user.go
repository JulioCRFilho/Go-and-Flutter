package model

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"hash"
)

type User struct {
	Id    primitive.ObjectID `bson:"_id"`
	Name  string             `bson:"name" binding:"required"`
	Email string             `bson:"email" binding:"required"`
	Pass  hash.Hash          `bson:"pass" binding:"required"`
	Doc   string             `bson:"doc" binding:"required"`
}
