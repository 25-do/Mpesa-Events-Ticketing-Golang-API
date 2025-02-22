package repositories_test

import (
	"fmt"
	"testing"
	"ticketing-system/models"
	"ticketing-system/repositories"
	"time"

	"github.com/stretchr/testify/assert"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

const timestampStr = "2025-02-02T07:05:27.860335+03:00"

func setupTestDB() *gorm.DB {
	db, _ := gorm.Open(sqlite.Open(":memory:"), &gorm.Config{})
	db.AutoMigrate(&models.Event{}) // Create tables in memory
	return db
}

func CreateEvent(db *gorm.DB) {
	repo := repositories.NewEventRepository(db)

	// Parse the string into a time.Time object
	timestamp, err := time.Parse(time.RFC3339, timestampStr)
	if err != nil {
		fmt.Println("Error parsing time:", err)
		return
	}
	repo.CreateEvent(&models.Event{Name: "Soul Fest",
		Description: "soul generation festival",
		VenueID:     5,
		OrganizerID: 1,
		StartTime:   timestamp,
		EndTime:     timestamp,
		IsFree:      true,
		ImageURL:    "https://cloudinary.com"})
}

func TestCreateEvent(t *testing.T) {
	db := setupTestDB()
	repo := repositories.NewEventRepository(db)

	// Parse the string into a time.Time object
	timestamp, err := time.Parse(time.RFC3339, timestampStr)
	if err != nil {
		fmt.Println("Error parsing time:", err)
		return
	}
	events, err := repo.CreateEvent(&models.Event{Name: "Soul Fest",
		Description: "soul generation festival",
		VenueID:     5,
		OrganizerID: 1,
		StartTime:   timestamp,
		EndTime:     timestamp,
		IsFree:      true,
		ImageURL:    "https://cloudinary.com"})

	assert.NoError(t, err)                    // Ensure no error occurred
	assert.NotNil(t, events)                  // Ensure event is not nil
	assert.Equal(t, "Soul Fest", events.Name) // Check if event name is correct

}
func TestGetAllEvents(t *testing.T) {
	db := setupTestDB()

	repo := repositories.NewEventRepository(db)

	timestamp, _ := time.Parse(time.RFC3339, timestampStr)
	repo.CreateEvent(&models.Event{Name: "Soul Fest",
		Description: "soul generation festival",
		VenueID:     5,
		OrganizerID: 1,
		StartTime:   timestamp,
		EndTime:     timestamp,
		IsFree:      true,
		ImageURL:    "https://cloudinary.com"})
	get_events, err := repo.GetAllEvents()

	// Assertions
	assert.NoError(t, err)
	assert.Len(t, get_events, 1)

}

func TestGetAllSingleEvents(t *testing.T) {
	db := setupTestDB()
	CreateEvent(db)
	repo := repositories.NewEventRepository(db)
	events, err := repo.GetSingleEvent(1)

	// Assertions
	assert.NoError(t, err)
	assert.Len(t, events, 1)
	assert.Equal(t, "Soul Fest", events[0].Name)

}
