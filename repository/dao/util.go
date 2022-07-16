package dao

import (
	"firstProject/repository/db"
	"go.mongodb.org/mongo-driver/mongo"
)

func getCollection(dbName string) *mongo.Collection {
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
