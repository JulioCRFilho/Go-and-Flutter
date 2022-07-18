package routes

import (
	"firstProject/repository/dao"
	"github.com/gin-gonic/gin"
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
}

func createUser(c *gin.Context) {

}

func deleteUser(c *gin.Context) {

}

func updateUser(c *gin.Context) {

}
