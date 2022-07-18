package dao

import (
	"errors"
	"firstProject/model"
	util "firstProject/repository"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
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

func GetUser(id string) (*model.User, error) {
	ctx, cancel := util.Context()

	defer cancel()

	collection := getCollection(userDB)

	var user model.User

	objId, err := primitive.ObjectIDFromHex(id)

	if err != nil {
		return nil, err
	}

	filter := bson.D{{"_id", objId}}

	if err = collection.FindOne(ctx, filter).Decode(&user); err != nil {
		return nil, err
	} else {
		return &user, nil
	}
}

func DeleteUser(id string) error {
	ctx, cancel := util.Context()

	defer cancel()

	collection := getCollection(userDB)

	objId, err := primitive.ObjectIDFromHex(id)

	if err != nil {
		return err
	}

	filter := bson.D{{"_id", objId}}

	var res *mongo.DeleteResult
	if res, err = collection.DeleteOne(ctx, filter); err != nil {
		return err
	}

	if res.DeletedCount == 1 {
		return nil
	} else {
		return errors.New(fmt.Sprintf("deleted %d documents", res.DeletedCount))
	}
}

func UpdateUser(user model.User) error {
	ctx, cancel := util.Context()

	defer cancel()

	collection := getCollection(userDB)

	filter := bson.D{{"_id", user.Id}}

	if _, err := collection.ReplaceOne(ctx, filter, user); err != nil {
		return err
	} else {
		return nil
	}
}
