package main

import (
	"firstProject/model"
	"firstProject/repository"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func getTasks(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"tasks": repository.GetTasks(),
	})
}

func getTask(c *gin.Context) {
	param := c.Param("id")
	id, err := strconv.Atoi(param)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
	} else {
		task, err := repository.GetTask(id)

		if err != nil {
			c.JSON(http.StatusNotFound, gin.H{
				"error": err.Error(),
			})
		} else {
			c.JSON(http.StatusOK, gin.H{
				"task": task,
			})
		}
	}
}

func createTask(c *gin.Context) {
	name := c.Param("name")
	dueDate := c.Param("dueDate")
	createdDate := c.Param("createdDate")

	task := model.Task{
		Id:          len(repository.GetTasks()) + 1,
		Name:        name,
		DueDate:     dueDate,
		CreatedDate: createdDate,
	}

	repository.CreateTask(task)

	c.JSON(http.StatusCreated, gin.H{
		"status": "success",
	})
}
