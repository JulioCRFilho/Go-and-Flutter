package dao

import (
	"errors"
	"firstProject/model"
	util "firstProject/repository"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

var dbName = "tasks"

func GetTasks(id string) (*[]model.Task, error) {
	ctx, cancel := util.Context()

	defer cancel()

	tasks := getCollection(dbName)

	if tasks == nil {
		return nil, errors.New("collection not retrievable")
	}

	hexId, errHex := primitive.ObjectIDFromHex(id)

	if errHex != nil {
		return nil, errHex
	}

	var tasksBson []model.Task

	filter := bson.D{{"user_id", hexId}}

	if task, err := tasks.Find(ctx, filter); err != nil {

		return nil, err
	} else {
		if err2 := task.All(ctx, &tasksBson); err2 != nil {
			print("err2", err2.Error())
			return nil, err2
		}

		return &tasksBson, nil
	}
}

func GetTask(taskId string, userId string) *model.Task {
	ctx, cancel := util.Context()

	defer cancel()

	tasks := getCollection(dbName)

	var task model.Task

	userObjId, err2 := primitive.ObjectIDFromHex(userId)

	if err2 != nil {
		print("err2:", err2.Error())
		return nil
	}

	taskObjId, err3 := primitive.ObjectIDFromHex(taskId)

	if err3 != nil {
		print("err3 :", err3.Error())
		return nil
	}

	filter := bson.D{{"_id", taskObjId}, {"user_id", userObjId}}
	if err := tasks.FindOne(ctx, filter).Decode(&task); err != nil {
		return nil
	} else {
		return &task
	}
}

func CreateTask(task model.Task) error {
	ctx, cancel := util.Context()

	defer cancel()

	tasks := getCollection(dbName)

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

	tasks := getCollection(dbName)

	filter := bson.D{{"_id", task.Id}, {"user_id", task.UserId}}

	if _, err := tasks.ReplaceOne(ctx, filter, task); err != nil {
		return err
	} else {
		return nil
	}
}

func DeleteTask(id string, userId string) error {
	ctx, cancel := util.Context()

	defer cancel()

	objID, err := primitive.ObjectIDFromHex(id)

	if err != nil {
		return err
	}

	userID, err3 := primitive.ObjectIDFromHex(userId)

	if err3 != nil {
		return err3
	}

	filter := bson.D{{"_id", objID}, {"user_id", userID}}

	tasks := getCollection(dbName)

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
