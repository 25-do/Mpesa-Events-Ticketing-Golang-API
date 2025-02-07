package routers

import (
	"ticketing-system/controllers"

	"github.com/gin-gonic/gin"
)

func EventRoutes(router *gin.RouterGroup, eventController *controllers.EventController) {
	eventRoutes := router.Group("/event")
	{
		eventRoutes.GET("/", eventController.GetAllEvents)
		eventRoutes.POST("/create", eventController.CreateEvent)
		eventRoutes.GET("/getSingle", eventController.GetSingleEvent)
		eventRoutes.PUT("/update", eventController.UpdateEvent)
		eventRoutes.DELETE("/delete", eventController.DeleteEvent)
	}
}
