package middlewares

import (
	"fmt"
	"github.com/gin-gonic/gin"
)

func Writer() gin.HandlerFunc {
	return func(c *gin.Context) {
		fmt.Printf(
			"\n\nRequest: %s\n"+
				"url: %s\n"+
				"header: %s\n"+
				"body: %s\n"+
				"\n\n",
			c.Request.Method,
			c.Request.RequestURI,
			c.Request.Header,
			StringCache(),
		)
	}
}
