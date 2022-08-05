package main

import (
	"firstProject/middlewares"
	"firstProject/repository/dao"
	"firstProject/repository/db"
	"firstProject/routes"
	"github.com/gin-contrib/cors"
	"github.com/gin-contrib/static"
	"github.com/gin-gonic/gin"
)

func main() {
	db.CreateClient()
	dao.GetDatabases()

	r := gin.Default()
	r.Use(cors.Default())
	r.Use(middlewares.Cache())
	r.Use(middlewares.Writer())

	routes.Setup(r)

	r.Use(static.Serve("/", static.LocalFile("./static", false)))

	err := r.Run()
	if err != nil {
		panic(err)
	}
}
