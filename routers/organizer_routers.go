package routers

import (
	"ticketing-system/controllers"

	"github.com/gin-gonic/gin"
)

func OrganizerRoutes(router *gin.RouterGroup, organizerController *controllers.OrganizerController) {
	organizerRoutes := router.Group("/organizer")
	{
		organizerRoutes.GET("/", organizerController.GetAllOrganizers)
		organizerRoutes.POST("/create", organizerController.CreateOrganizer)
		organizerRoutes.GET("/getSingle", organizerController.GetSingleOrganizer)
		organizerRoutes.PUT("/update", organizerController.UpdateOrganizer)
		organizerRoutes.DELETE("/delete", organizerController.DeleteOrganizer)
	}
}
