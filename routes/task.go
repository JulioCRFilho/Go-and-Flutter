package routes

import (
	"firstProject/model"
	"firstProject/repository/dao"
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
	if tasks, err := dao.GetTasks(); err != nil {
		c.String(500, err.Error())
	} else {
		c.JSON(http.StatusOK, gin.H{
			"tasks": tasks,
		})
	}
}

func getTask(c *gin.Context) {
	id := c.Param("id")

	if task := dao.GetTask(id); task == nil {
		c.JSON(http.StatusNotFound, gin.H{
			"error": "task não encontrada",
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
		if err2 := dao.CreateTask(task); err2 != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err2.Error(),
			})
			return
		}

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
		err2 := dao.DeleteTask(v)

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
		err2 := dao.UpdateTask(task)

		if err2 != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err2.Error(),
			})
		} else {
			c.String(200, "Updated")
		}
	}
}
