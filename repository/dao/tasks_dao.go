package dao

import (
	"context"
	"errors"
	"firstProject/model"
	util "firstProject/repository"
	"firstProject/repository/db"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"time"
)

var dbName = "tasks"

func GetTasks() (*[]model.Task, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)

	defer cancel()

	tasks := getCollection()

	if tasks == nil {
		return nil, errors.New("collection not retrievable")
	}

	var tasksBson []model.Task

	if task, err := tasks.Find(ctx, bson.M{}); err != nil {

		return nil, err
	} else {
		if err2 := task.All(ctx, &tasksBson); err2 != nil {
			print("err2", err2.Error())
			return nil, err2
		}

		return &tasksBson, nil
	}
}

func GetTask(id string) *model.Task {
	ctx, cancel := util.Context()

	defer cancel()

	tasks := getCollection()

	var task model.Task

	objId, err2 := primitive.ObjectIDFromHex(id)

	if err2 != nil {
		print("err2 :", err2.Error())
		return nil
	}

	if err := tasks.FindOne(ctx, bson.D{{"_id", objId}}).Decode(&task); err != nil {
		print("erro findOne:", err.Error())
		return nil
	} else {
		return &task
	}
}

func CreateTask(task model.Task) error {
	ctx, cancel := util.Context()

	defer cancel()

	tasks := getCollection()

	if tasks == nil {
		return errors.New("falha ao recuperar a collection de Tasks")
	}

	task.Id = primitive.NewObjectID()

	if _, err := tasks.InsertOne(ctx, task); err != nil {
		return err
	} else {
		return nil
	}
}

func UpdateTask(task model.Task) error {
	ctx, cancel := util.Context()

	defer cancel()

	tasks := getCollection()

	if tasks == nil {
		return errors.New("collection not retrievable")
	}

	filter := bson.D{{"_id", task.Id}}

	if _, err := tasks.ReplaceOne(ctx, filter, task); err != nil {
		return err
	} else {
		return nil
	}
}

func DeleteTask(id string) error {
	ctx, cancel := util.Context()

	defer cancel()

	objID, err := primitive.ObjectIDFromHex(id)

	if err != nil {
		return err
	}

	filter := bson.D{{"_id", objID}}

	tasks := getCollection()

	if res, err2 := tasks.DeleteOne(ctx, filter); err2 != nil {
		return err2
	} else {
		if count := res.DeletedCount; count != 1 {
			return errors.New("document not found")
		} else {
			return nil
		}
	}
}

func getCollection() *mongo.Collection {
	tasksDB := db.Client.Database(dbName)

	if tasksDB == nil {
		return nil
	}

	tasks := tasksDB.Collection(dbName)

	if tasks == nil {
		return nil
	}

	return tasks
}
