package routers

import (
	"ticketing-system/controllers"

	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	r := gin.Default()

	// User routes
	userGroup := r.Group("/users")
	{
		userGroup.POST("/", controllers.CreateUser)
	}

	return r
}
