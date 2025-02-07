package db

import (
	"fmt"
	"log"
	"os"

	"ticketing-system/models"

	"github.com/joho/godotenv" // Load environment variables from .env file
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

// ConnectDB establishes a connection to the database.
func ConnectDB() {
	// Load environment variables from the .env file
	err := godotenv.Load()
	if err != nil {
		log.Println("Warning: Could not load .env file. Make sure environment variables are set.")
	}

	// Construct the DSN (Data Source Name)
	dsn := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%s sslmode=disable",
		getEnv("DB_HOST", "localhost"),
		getEnv("DB_USER", "postgres"),
		getEnv("DB_PASSWORD", ""),
		getEnv("DB_NAME", "ticketing_system"),
		getEnv("DB_PORT", "5432"),
	)

	// Connect to the PostgreSQL database
	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect to the database:", err)
	}

	log.Println("Database connection established successfully.")
}

// MigrateTables automigrates the database tables based on the models.
func MigrateTables() {
	err := DB.AutoMigrate(
		&models.User{},
		&models.Venue{},
		&models.Event{},
		&models.TicketType{},
		&models.Ticket{},
		&models.Payment{},
		&models.Organizer{},
	)
	if err != nil {
		log.Fatal("Failed to migrate tables:", err)
	}

	log.Println("Database tables migrated successfully.")
}

// getEnv retrieves environment variables or returns a default value if not set.
func getEnv(key, defaultValue string) string {
	value := os.Getenv(key)
	if value == "" {
		return defaultValue
	}
	return value
}
