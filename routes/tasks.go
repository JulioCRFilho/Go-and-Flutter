package routes

import (
	"firstProject/middlewares"
	"firstProject/model"
	"firstProject/repository/dao"
	"github.com/gin-gonic/gin"
	json2 "github.com/goccy/go-json"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"net/http"
)

func DefineTaskRoutes(c *gin.Engine) {
	c.GET("/tasks/:user_id", getTasks)
	r := c.Group("/task").Use(middlewares.Auth())
	{
		r.GET(":user_id/:task_id", getTask)
		r.POST(":user_id", createTask)
		r.PUT(":user_id", updateTask)
		r.DELETE(":user_id/:task_id", deleteTask)
	}
}

func getTasks(c *gin.Context) {
	id := c.Param("user_id")

	if tasks, err := dao.GetTasks(id); err != nil {
		c.String(500, err.Error())
	} else {
		c.JSON(http.StatusOK, gin.H{
			"tasks": tasks,
		})
	}
}

func getTask(c *gin.Context) {
	taskId := c.Param("task_id")
	userId := c.Param("user_id")

	if task := dao.GetTask(taskId, userId); task == nil {
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
	userId := c.Param("user_id")

	var task model.Task

	if err := json2.Unmarshal(middlewares.RawCache(), &task); err != nil {
		c.JSON(500, gin.H{
			"error": err.Error(),
		})
		print("error unmarshal:", err.Error())
		return
	}

	task.Done = false

	if objId, err3 := primitive.ObjectIDFromHex(userId); err3 != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err3.Error(),
		})
	} else {
		task.UserId = objId

		if err2 := dao.CreateTask(task); err2 != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err2.Error(),
			})
			return
		}

		c.String(200, "success")
	}
}

func deleteTask(c *gin.Context) {
	taskId := c.Param("task_id")
	userId := c.Param("user_id")

	if err := dao.DeleteTask(taskId, userId); err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
	} else {
		c.String(200, "success")
	}
}

func updateTask(c *gin.Context) {
	userId := c.Param("user_id")

	var task model.Task
	err := json2.Unmarshal(middlewares.RawCache(), &task)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
	} else {
		if objId, err3 := primitive.ObjectIDFromHex(userId); err3 != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err3.Error(),
			})
		} else {
			task.UserId = objId

			err2 := dao.UpdateTask(task)

			if err2 != nil {
				c.JSON(http.StatusBadRequest, gin.H{
					"error": err2.Error(),
				})
			} else {
				c.String(200, "success")
			}
		}
	}
}
