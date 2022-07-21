package model

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"golang.org/x/crypto/bcrypt"
)

type User struct {
	Id    primitive.ObjectID `bson:"_id"`
	Name  string             `bson:"name" binding:"required"`
	Email string             `bson:"email" binding:"required" valid:"email" unique:"true"`
	Pass  string             `bson:"pass" binding:"required"`
	Doc   string             `bson:"doc" binding:"required" unique:"true"`
}

func (user *User) HashPassword() error {
	if bytes, err := bcrypt.GenerateFromPassword([]byte(user.Pass), 14); err != nil {
		return err
	} else {
		user.Pass = string(bytes)
		return nil
	}
}

func (user *User) CheckPassword(providedPass string) error {
	if err := bcrypt.CompareHashAndPassword([]byte(user.Pass), []byte(providedPass)); err != nil {
		return err
	}
	return nil
}
