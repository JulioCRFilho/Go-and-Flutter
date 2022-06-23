package main

import (
	"firstProject/repository"
	"github.com/gin-gonic/gin"
	"net/http"
)

func home(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"tasks": repository.GetTasks(),
	})
}
