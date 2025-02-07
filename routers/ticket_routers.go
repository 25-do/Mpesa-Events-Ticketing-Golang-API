package routers

import (
	"ticketing-system/controllers"

	"github.com/gin-gonic/gin"
)

func TicketRoutes(router *gin.RouterGroup, ticketController *controllers.TicketController) {
	ticketRoutes := router.Group("/ticket")
	{
		ticketRoutes.GET("/", ticketController.GetAllTickets)
		ticketRoutes.POST("/create", ticketController.CreateTicket)
		ticketRoutes.GET("/getSingle", ticketController.GetSingleTicket)
		ticketRoutes.PUT("/update", ticketController.UpdateTicket)
		ticketRoutes.DELETE("/delete", ticketController.DeleteTicket)
	}
}
