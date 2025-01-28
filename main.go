package main

import (
	"ticketing-system/controllers"
	"ticketing-system/db"
	"ticketing-system/repositories"
	"ticketing-system/routers"
	"ticketing-system/services"

	"github.com/gin-gonic/gin"
)

func main() {
	// Connect to the database
	db.ConnectDB()
	db.MigrateTables()

	// Initialize repository, service, and controller
	userRepo := repositories.NewUserRepository(db.DB)
	userService := services.NewUserService(userRepo)
	userController := controllers.NewUserController(userService)

	// Initialize Gin router
	r := gin.Default()

	// Register routes
	api := r.Group("/api")
	routers.RegisterUserRoutes(api, userController) // Register user routes

	// Start the server
	r.Run(":8080")
}
