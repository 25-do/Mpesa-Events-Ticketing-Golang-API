package routers

import (
	"ticketing-system/controllers"

	"github.com/gin-gonic/gin"
)

func RegisterUserRoutes(routerGroup *gin.RouterGroup, userController *controllers.UserController) {
	userRoutes := routerGroup.Group("/users")
	{
		userRoutes.GET("/", userController.GetAllUsers)    // Route to get all users
		userRoutes.GET("/:id", userController.GetUserByID) // Route to get a user by ID
	}
}
