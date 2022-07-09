package dao

import (
	"context"
	"firstProject/repository/db"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"time"
)

func GetDatabases() {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)

	defer cancel()

	if dbs, err := db.Client.ListDatabaseNames(ctx, bson.M{}); err != nil {
		fmt.Printf("Failed to get databases: %s", err.Error())
	} else {
		fmt.Println(dbs)
	}
}
