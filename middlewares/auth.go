package middlewares

import (
	"firstProject/repository/auth"
	"github.com/gin-gonic/gin"
	"strings"
)

func Auth() gin.HandlerFunc {
	return func(c *gin.Context) {
		token := c.GetHeader("Authorization")

		if splitBearer := strings.Split(token, "Bearer "); len(splitBearer) != 2 {
			c.JSON(401, gin.H{
				"error": "lacking bearer token",
			})
			c.Abort()
			return

		} else {
			authToken := splitBearer[1]

			if authToken == "" {
				c.JSON(401, gin.H{
					"error": "lacking access token",
				})
				c.Abort()
				return
			}

			if err := auth.ValidateToken(authToken); err != nil {
				c.JSON(401, gin.H{
					"error": err.Error(),
				})
				c.Abort()
				return
			}

			c.Next()
		}
	}
}
