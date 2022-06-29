package main

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-contrib/static"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.Use(cors.Default())

	r.GET("/tasks", getTasks)
	r.GET("/task/:id", getTask)
	r.POST("/task", createTask)
	r.PUT("/task", updateTask)
	r.DELETE("/task/:id", deleteTask)

	r.Use(static.Serve("/", static.LocalFile("./static", false)))

	err := r.Run()
	if err != nil {
		panic(err)
	}
}
