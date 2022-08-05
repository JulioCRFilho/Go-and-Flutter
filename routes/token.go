package routes

import (
	"firstProject/middlewares"
	"firstProject/model"
	"firstProject/repository/auth"
	"firstProject/repository/dao"
	"github.com/gin-gonic/gin"
	json2 "github.com/goccy/go-json"
)

func DefineTokenRoutes(c *gin.Engine) {
	group := c.Group("/auth")
	{
		group.POST("/token", generateToken)
	}
}

func generateToken(c *gin.Context) {
	var token model.Token

	if err := json2.Unmarshal(middlewares.RawCache(), &token); err != nil {
		c.String(500, err.Error())
		return
	}

	if user, err := dao.GetByEmail(token.Email); err != nil {
		c.String(500, err.Error())
	} else {
		if err2 := user.CheckPassword(token.Pass); err2 != nil {
			c.String(401, err2.Error())
			return
		}

		if stringToken, err3 := auth.GenerateToken(user.Email, user.Pass); err3 != nil {
			c.String(500, err3.Error())
		} else {
			c.JSON(200, gin.H{
				"token": stringToken,
			})
		}
	}
}
