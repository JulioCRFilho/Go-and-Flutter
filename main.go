package main

import (
	"context"
	"firstProject/repository/db"
	"firstProject/routes"
	"github.com/gin-contrib/cors"
	"github.com/gin-contrib/static"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"time"
)

func main() {
	r := gin.Default()
	r.Use(cors.Default())

	conn := db.LoadAppConfig()

	serverAPIOptions := options.ServerAPI(options.ServerAPIVersion1)
	clientOptions := options.Client().ApplyURI(conn).SetServerAPIOptions(serverAPIOptions)

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	client, dbErr := mongo.Connect(ctx, clientOptions)
	if dbErr != nil {
		panic(dbErr)
	}

	print(client, "\n\nchegou aqui")

	routes.DefineTaskRoutes(r)

	r.Use(static.Serve("/", static.LocalFile("./static", false)))

	err := r.Run()
	if err != nil {
		panic(err)
	}
}
