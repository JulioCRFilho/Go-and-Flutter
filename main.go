package main

import (
	"firstProject/routes"
	"github.com/gin-contrib/cors"
	"github.com/gin-contrib/static"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.Use(cors.Default())

	routes.DefineTaskRoutes(r)

	r.Use(static.Serve("/", static.LocalFile("./static", false)))

	err := r.Run()
	if err != nil {
		panic(err)
	}
}
