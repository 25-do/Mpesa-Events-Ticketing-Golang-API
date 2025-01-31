package controllers

import (
	"net/http"
	"ticketing-system/services"

	"github.com/gin-gonic/gin"
)

type EventController struct {
	EventService services.EventServiceInterface
}

func NewEventController(eventService services.EventServiceInterface) *EventController {
	return &EventController{EventService: eventService}
}

func (uc *EventController) GetAllEvents(c *gin.Context) {
	event, err := uc.EventService.GetAllEvents()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve users"})
		return
	}
	c.JSON(http.StatusOK, event)
}
