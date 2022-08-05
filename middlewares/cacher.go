package middlewares

import (
	"github.com/gin-gonic/gin"
	"io/ioutil"
)

var cache []byte

func Cache() gin.HandlerFunc {
	return func(c *gin.Context) {
		if bytes, err := ioutil.ReadAll(c.Request.Body); err != nil {
			c.String(500, err.Error())
			c.Abort()
		} else {
			cache = bytes
			print("cached body:", string(cache))
			c.Next()
		}
	}
}

func StringCache() string {
	return string(cache)
}

func RawCache() []byte {
	return cache
}
