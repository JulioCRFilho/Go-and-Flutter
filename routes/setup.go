package routes

import "github.com/gin-gonic/gin"

func Setup(router *gin.Engine) {
	DefineTaskRoutes(router)
	DefineUserRoutes(router)
	DefineTokenRoutes(router)
}
