package routes

import (
	"firstProject/model"
	"firstProject/repository"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func DefineTaskRoutes(c *gin.Engine) {
	c.GET("/tasks", getTasks)
	r := c.Group("/task")
	{
		r.GET(":id", getTask)
		r.POST("", createTask)
		r.PUT("", updateTask)
		r.DELETE(":id", deleteTask)
	}
}

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
		return
	}

	if task, err2 := repository.GetTask(id); err2 != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"error": err2.Error(),
		})
	} else {
		c.JSON(http.StatusOK, gin.H{
			"task": task,
		})
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
			"error": "Id inválido",
		})
		return
	}

	if v, err := strconv.Atoi(id); err == nil {
		err2 := repository.DeleteTask(v)

		if err2 != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err2.Error(),
			})
		} else {
			c.String(200, "Success")
		}
	} else {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
	}
}

func updateTask(c *gin.Context) {
	var task model.Task
	err := c.ShouldBind(&task)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
	} else {
		err2 := repository.UpdateTask(task)

		if err2 != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err2.Error(),
			})
		} else {
			c.String(200, "Updated")
		}
	}
}
