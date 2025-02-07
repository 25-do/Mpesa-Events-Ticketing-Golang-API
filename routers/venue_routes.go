package routers

import (
	"ticketing-system/controllers"

	"github.com/gin-gonic/gin"
)

func VenueRoutes(router *gin.RouterGroup, venueController *controllers.VenueController) {
	venueRoutes := router.Group("/venue")
	{
		venueRoutes.GET("/", venueController.GetAllVenues)
		venueRoutes.POST("/create", venueController.CreateVenue)
		venueRoutes.GET("/getSingle", venueController.GetSingleVenue)
		venueRoutes.PUT("/update", venueController.UpdateVenue)
		venueRoutes.DELETE("/delete", venueController.DeleteVenue)
	}
}
