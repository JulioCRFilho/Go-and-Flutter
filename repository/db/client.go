package db

import (
	util "firstProject/repository"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var Client *mongo.Client

func CreateClient() {
	var err error
	conn := LoadAppConfig()

	serverAPIOptions := options.ServerAPI(options.ServerAPIVersion1)
	clientOptions := options.Client().ApplyURI(conn).SetServerAPIOptions(serverAPIOptions)

	ctx, cancel := util.Context()

	defer cancel()

	Client, err = mongo.Connect(ctx, clientOptions)
	if err != nil {
		panic(err)
	}
}
