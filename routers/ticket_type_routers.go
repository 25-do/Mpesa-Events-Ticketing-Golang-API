package routers

import (
	"ticketing-system/controllers"

	"github.com/gin-gonic/gin"
)

func TicketTypeRoutes(router *gin.RouterGroup, tickettypeController *controllers.TicketTypeController) {
	tickettypeRoutes := router.Group("/tickettype")
	{
		tickettypeRoutes.GET("/", tickettypeController.GetAllTicketTypes)
		tickettypeRoutes.POST("/create", tickettypeController.CreateTicketType)
		tickettypeRoutes.GET("/getSingle", tickettypeController.GetSingleTicketType)
		tickettypeRoutes.PUT("/update", tickettypeController.UpdateTicketType)
		tickettypeRoutes.DELETE("/delete", tickettypeController.DeleteTicketType)
	}
}
