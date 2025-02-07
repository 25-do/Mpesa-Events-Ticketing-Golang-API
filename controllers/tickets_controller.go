package controllers

import (
	"fmt"
	"net/http"
	"strconv"
	"ticketing-system/models"
	"ticketing-system/services"

	"github.com/gin-gonic/gin"
)

type TicketController struct {
	TicketService services.TicketServiceInterface
}

func NewTicketController(ticketService services.TicketServiceInterface) *TicketController {
	return &TicketController{TicketService: ticketService}

}

func (ctrl *TicketController) CreateTicket(c *gin.Context) {
	var ticket models.Ticket
	if err := c.ShouldBindJSON(&ticket); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	newTicket, err := ctrl.TicketService.CreateTicket(&ticket)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, newTicket)
}

func (ctrl *TicketController) GetAllTickets(c *gin.Context) {
	ticket, err := ctrl.TicketService.GetAllTickets()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve tickets"})
		return
	}
	c.JSON(http.StatusOK, ticket)

}

func (ctrl *TicketController) GetSingleTicket(c *gin.Context) {
	idParam := c.Query("id")
	fmt.Println(idParam)
	id, err := strconv.Atoi(idParam)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	ticket, err := ctrl.TicketService.GetSingleTicket(uint(id))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrive ticket"})
		return
	}
	c.JSON(http.StatusOK, ticket)
}

func (ctrl *TicketController) UpdateTicket(c *gin.Context) {
	var ven models.Ticket
	idParam := c.Query("id")
	fmt.Println(idParam)
	id, err := strconv.Atoi(idParam)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "an error occured check your Query parameters"})
		return
	}
	if err := c.ShouldBindJSON(&ven); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	ticket, err := ctrl.TicketService.UpdateTicket(uint(id), ven)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "an error occured while updating the ticket"})
		return
	}
	c.JSON(http.StatusOK, ticket)
}

func (ctrl *TicketController) DeleteTicket(c *gin.Context) {
	idParam := c.Query("id")
	fmt.Println(idParam)
	id, err := strconv.Atoi(idParam)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	ticket, err := ctrl.TicketService.DeleteTicket(uint(id))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete ticket"})
		return
	}
	c.JSON(http.StatusOK, ticket)
}
