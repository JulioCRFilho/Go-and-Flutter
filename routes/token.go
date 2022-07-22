package routes

import (
	"firstProject/repository/auth"
	"firstProject/repository/dao"
	"github.com/gin-gonic/gin"
)

func DefineTokenRoutes(c *gin.Engine) {
	group := c.Group("/auth")
	{
		group.POST("/token", generateToken)
		group.POST("/validate", validateToken)
	}
}

func generateToken(c *gin.Context) {
	id := c.Param("id")
	pass := c.Param("pass")

	var token auth.JWT

	if err := c.ShouldBind(&token); err != nil {
		c.String(500, err.Error())
		return
	}

	if user, err := dao.GetUser(id); err != nil {
		c.String(500, err.Error())
	} else {
		if err2 := user.CheckPassword(pass); err2 != nil {
			c.String(401, err2.Error())
			return
		}

		if stringToken, err3 := token.GenerateToken(); err3 != nil {
			c.String(500, err3.Error())
		} else {
			c.JSON(200, gin.H{
				"token": stringToken,
			})
		}
	}

}

func validateToken(c *gin.Context) {

}
