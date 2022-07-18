package routes

import (
	"firstProject/model"
	"firstProject/repository/dao"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func DefineUserRoutes(c *gin.Engine) {
	c.GET("/users", getUsers)
	group := c.Group("/user")
	{
		group.GET(":id", getUser)
		group.PUT("", updateUser)
		group.POST("", createUser)
		group.DELETE(":id", deleteUser)
	}
}

func getUsers(c *gin.Context) {
	if users, err := dao.GetUsers(); err != nil {
		c.String(400, err.Error())
	} else {
		c.JSON(200, gin.H{
			"users": users,
		})
	}
}

func getUser(c *gin.Context) {
	id := c.Param("id")

	if user, err := dao.GetUser(id); err != nil {
		c.String(400, err.Error())
	} else {
		c.JSON(200, gin.H{
			"user": user,
		})
	}
}

func createUser(c *gin.Context) {
	var user model.User

	if err := c.ShouldBind(&user); err != nil {
		c.JSON(400, err.Error())
	}

	user.Id = primitive.NewObjectID()

	if err := dao.CreateUser(user); err != nil {
		c.JSON(400, err.Error())
	} else {
		c.JSON(200, "success")
	}
}

func deleteUser(c *gin.Context) {
	id := c.Param("id")

	if err := dao.DeleteUser(id); err != nil {
		c.JSON(400, err.Error())
	} else {
		c.JSON(200, "success")
	}
}

func updateUser(c *gin.Context) {
	var user model.User

	if err := c.ShouldBind(&user); err != nil {
		c.String(400, err.Error())
	}

	if err := dao.UpdateUser(user); err != nil {
		c.String(400, err.Error())
	} else {
		c.String(200, "success")
	}
}
