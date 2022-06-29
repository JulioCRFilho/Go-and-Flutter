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
	var task model.Task

	if err := c.ShouldBind(&task); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
	} else {
		repository.CreateTask(task)

		c.JSON(http.StatusCreated, gin.H{
			"status": "success",
		})
	}
}

func deleteTask(c *gin.Context) {
	id := c.Param("id")

	if id == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Invalid id",
		})
	} else {
		if v, err := strconv.Atoi(id); err == nil {
			err2 := repository.DeleteTask(v)

			if err2 != nil {
				c.JSON(http.StatusBadRequest, gin.H{
					"error": "Task not found",
				})
			} else {
				c.String(200, "Success")
			}
		} else {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": "Invalid id",
			})
		}
	}
}
