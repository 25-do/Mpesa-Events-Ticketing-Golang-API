package main

import (
	"ticketing-system/db"
	"ticketing-system/routers"
)

func main() {
	// Connect to the database
	db.ConnectDB()

	// Migrate tables
	db.MigrateTables()

	// Setup router
	r := routers.SetupRouter()
	r.Run(":8080") // Run server on port 8080
}
