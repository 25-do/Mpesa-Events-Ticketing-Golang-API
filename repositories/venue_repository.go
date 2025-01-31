package repositories

import (
	"ticketing-system/models"

	"gorm.io/gorm"
)

type VenueRepositoryInterface interface {
	CreateVenue(event *models.Event) (*models.Event, error)
	GetAllVenues() ([]models.Event, error)
	// GetSingleVenue(id uint) (*models.Event, error)
	// UpdateVenue(id uint) (*models.Event, error)
	// DeleteVenue(id uint) (*models.Event, error)
}

type venueRepository struct {
	DB *gorm.DB
}

func NewVenueRepository(db *gorm.DB) VenueRepositoryInterface {
	return &venueRepository{DB: db}
}
func (dc *venueRepository) CreateVenue(event *models.Event) (*models.Event, error) {
	if err := dc.DB.Create(&event).Error; err != nil {
		return nil, err
	}

	return event, nil
}

func (dc *venueRepository) GetAllVenues() ([]models.Event, error) {
	var events []models.Event
	if err := dc.DB.Find(&events).Error; err != nil {

		return events, err
	}

	return events, nil
}
