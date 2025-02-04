package controllers

import (
	"net/http"
	"strconv"
	"ticketing-system/models"
	"ticketing-system/services"

	"github.com/gin-gonic/gin"
)

type VenueController struct {
	VenueService services.VenueServiceInterface
}

func NewVenueController(venueService services.VenueServiceInterface) *VenueController {
	return &VenueController{VenueService: venueService}

}

func (ctrl *VenueController) CreateVenue(c *gin.Context) {
	var venue models.Venue
	if err := c.ShouldBindJSON(&venue); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	newVenue, err := ctrl.VenueService.CreateVenue(&venue)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, newVenue)
}

func (ctrl *VenueController) GetAllVenues(c *gin.Context) {
	venue, err := ctrl.VenueService.GetAllVenues()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve venues"})
		return
	}
	c.JSON(http.StatusOK, venue)

}

func (ctrl *VenueController) GetSingleVenue(c *gin.Context) {
	idParam := c.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err})
		return
	}
	venue, err := ctrl.VenueService.GetSingleVenue(uint(id))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrive venue"})
		return
	}
	c.JSON(http.StatusOK, venue)
}

func (ctrl *VenueController) UpdateVenue(c *gin.Context) {
	id := c.Param("id")
	venue, err := ctrl.VenueService.UpdateVenue(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update venue"})
		return
	}
	c.JSON(http.StatusOK, venue)
}

func (ctrl *VenueController) DeleteVenue(c *gin.Context) {
	id := c.Param("id")
	venue, err := ctrl.VenueService.DeleteVenue(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete venue"})
		return
	}
	c.JSON(http.StatusOK, venue)
}
