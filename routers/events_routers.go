package routers

import (
	"ticketing-system/controllers"

	"github.com/gin-gonic/gin"
)

func EventRoutes(router *gin.RouterGroup, eventController *controllers.EventController) {
	userRoutes := router.Group("/events")
	{
		userRoutes.GET("/", eventController.GetAllEvents) // Route to get all users
	}
}
