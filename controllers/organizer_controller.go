package controllers

import (
	"fmt"
	"net/http"
	"strconv"
	"ticketing-system/models"
	"ticketing-system/services"

	"github.com/gin-gonic/gin"
)

type OrganizerController struct {
	OrganizerService services.OrganizerServiceInterface
}

func NewOrganizerController(organizerService services.OrganizerServiceInterface) *OrganizerController {
	return &OrganizerController{OrganizerService: organizerService}

}

func (ctrl *OrganizerController) CreateOrganizer(c *gin.Context) {
	var organizer models.Organizer
	if err := c.ShouldBindJSON(&organizer); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	newOrganizer, err := ctrl.OrganizerService.CreateOrganizer(&organizer)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, newOrganizer)
}

func (ctrl *OrganizerController) GetAllOrganizers(c *gin.Context) {
	organizer, err := ctrl.OrganizerService.GetAllOrganizers()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve organizers"})
		return
	}
	c.JSON(http.StatusOK, organizer)

}

func (ctrl *OrganizerController) GetSingleOrganizer(c *gin.Context) {
	idParam := c.Query("id")
	fmt.Println(idParam)
	id, err := strconv.Atoi(idParam)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	organizer, err := ctrl.OrganizerService.GetSingleOrganizer(uint(id))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrive organizer"})
		return
	}
	c.JSON(http.StatusOK, organizer)
}

func (ctrl *OrganizerController) UpdateOrganizer(c *gin.Context) {
	var ven models.Organizer
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
	organizer, err := ctrl.OrganizerService.UpdateOrganizer(uint(id), ven)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "an error occured while updating the organizer"})
		return
	}
	c.JSON(http.StatusOK, organizer)
}

func (ctrl *OrganizerController) DeleteOrganizer(c *gin.Context) {
	idParam := c.Query("id")
	fmt.Println(idParam)
	id, err := strconv.Atoi(idParam)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	organizer, err := ctrl.OrganizerService.DeleteOrganizer(uint(id))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete organizer"})
		return
	}
	c.JSON(http.StatusOK, organizer)
}
