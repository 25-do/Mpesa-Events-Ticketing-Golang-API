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

	venueRepo := repositories.NewVenueRepository(db.DB)
	venueService := services.NewVenueService(venueRepo)
	venueController := controllers.NewVenueController(venueService)

	eventRepo := repositories.NewEventRepository(db.DB)
	eventService := services.NewEventService(eventRepo)
	eventController := controllers.NewEventController(eventService)

	organizerRepo := repositories.NewOrganizerRepository(db.DB)
	organizerService := services.NewOrganizerService(organizerRepo)
	organizerController := controllers.NewOrganizerController(organizerService)

	tickettypeRepo := repositories.NewTicketTypeRepository(db.DB)
	tickettypeService := services.NewTicketTypeService(tickettypeRepo)
	tickettypeController := controllers.NewTicketTypeController(tickettypeService)

	paymentRepo := repositories.NewPaymentRepository(db.DB)
	paymentService := services.NewPaymentService(paymentRepo)
	paymentController := controllers.NewPaymentController(paymentService)

	// Initialize Gin router
	r := gin.Default()

	// Register routes
	api := r.Group("/api")
	routers.RegisterUserRoutes(api, userController) // Register user routes
	routers.EventRoutes(api, eventController)
	routers.VenueRoutes(api, venueController)
	routers.OrganizerRoutes(api, organizerController)
	routers.TicketTypeRoutes(api, tickettypeController)
	routers.PaymentRoutes(api, paymentController)

	// Start the server
	r.Run(":8080")
}
