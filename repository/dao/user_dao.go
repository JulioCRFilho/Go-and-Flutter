package dao

import (
	"firstProject/model"
	util "firstProject/repository"
	"go.mongodb.org/mongo-driver/bson"
)

var userDB = "userDB"

func CreateUser(user model.User) error {
	ctx, cancel := util.Context()

	defer cancel()

	collection := getCollection(userDB)

	if res, err := collection.InsertOne(ctx, user); err != nil {
		return err
	} else {
		print("\n\ninserted ID:", res.InsertedID, "\n\n")
		return nil
	}
}

func GetUsers() (*[]model.User, error) {
	ctx, cancel := util.Context()

	defer cancel()

	collection := getCollection(userDB)

	var users []model.User

	if res, err := collection.Find(ctx, bson.M{}); err != nil {
		return nil, err
	} else {
		if err2 := res.All(ctx, &users); err2 != nil {
			return nil, err2
		}

		return &users, nil
	}
}
