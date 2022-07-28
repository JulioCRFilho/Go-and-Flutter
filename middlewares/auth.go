package middlewares

import (
	"firstProject/repository/auth"
	"github.com/gin-gonic/gin"
)

func Auth() gin.HandlerFunc {
	return func(c *gin.Context) {
		token := c.GetHeader("Authorization")

		if token == "" {
			c.JSON(401, gin.H{
				"error": "lacking access token",
			})
			c.Abort()
		}

		if err := auth.ValidateToken(token); err != nil {
			c.JSON(401, gin.H{
				"error": err.Error(),
			})
			c.Abort()
		}

		c.Next()
	}
}
