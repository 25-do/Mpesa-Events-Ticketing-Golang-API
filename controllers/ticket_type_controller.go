package controllers

import (
	"fmt"
	"net/http"
	"strconv"
	"ticketing-system/models"
	"ticketing-system/services"

	"github.com/gin-gonic/gin"
)

type TicketTypeController struct {
	TicketTypeService services.TicketTypeServiceInterface
}

func NewTicketTypeController(tickettypeService services.TicketTypeServiceInterface) *TicketTypeController {
	return &TicketTypeController{TicketTypeService: tickettypeService}

}

func (ctrl *TicketTypeController) CreateTicketType(c *gin.Context) {
	var tickettype models.TicketType
	if err := c.ShouldBindJSON(&tickettype); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	newTicketType, err := ctrl.TicketTypeService.CreateTicketType(&tickettype)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, newTicketType)
}

func (ctrl *TicketTypeController) GetAllTicketTypes(c *gin.Context) {
	tickettype, err := ctrl.TicketTypeService.GetAllTicketTypes()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve tickettypes"})
		return
	}
	c.JSON(http.StatusOK, tickettype)

}

func (ctrl *TicketTypeController) GetSingleTicketType(c *gin.Context) {
	idParam := c.Query("id")
	fmt.Println(idParam)
	id, err := strconv.Atoi(idParam)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	tickettype, err := ctrl.TicketTypeService.GetSingleTicketType(uint(id))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrive tickettype"})
		return
	}
	c.JSON(http.StatusOK, tickettype)
}

func (ctrl *TicketTypeController) UpdateTicketType(c *gin.Context) {
	var ven models.TicketType
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
	tickettype, err := ctrl.TicketTypeService.UpdateTicketType(uint(id), ven)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "an error occured while updating the tickettype"})
		return
	}
	c.JSON(http.StatusOK, tickettype)
}

func (ctrl *TicketTypeController) DeleteTicketType(c *gin.Context) {
	idParam := c.Query("id")
	fmt.Println(idParam)
	id, err := strconv.Atoi(idParam)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	tickettype, err := ctrl.TicketTypeService.DeleteTicketType(uint(id))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete tickettype"})
		return
	}
	c.JSON(http.StatusOK, tickettype)
}
