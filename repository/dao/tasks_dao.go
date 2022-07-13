package dao

import (
	"context"
	"errors"
	"firstProject/model"
	util "firstProject/repository"
	"firstProject/repository/db"
	"go.mongodb.org/mongo-driver/bson"
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

func GetTask(id int) *model.Task {
	return nil
}

func CreateTask(task model.Task) error {
	ctx, cancel := util.Context()

	defer cancel()

	tasks := getCollection()

	if tasks == nil {
		return errors.New("falha ao recuperar a collection de Tasks")
	}

	if result, err := tasks.InsertOne(ctx, task); err != nil {
		return err
	} else {
		print(result, "resultado")
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

	if id, err := tasks.UpdateByID(ctx, task.Id, task); err != nil {
		return err
	} else {
		print("id da task atualizada:", id)
		return nil
	}
}

func DeleteTask(id int) error {
	return nil
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
