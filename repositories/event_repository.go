package repositories

import (
	"ticketing-system/models"

	"gorm.io/gorm"
)

type EventRepositoryInterface interface {
	GetAllEvents() ([]models.Event, error)
	CreateEvent(event *models.Event) (*models.Event, error)
}

type eventRepository struct {
	DB *gorm.DB
}

func NewEventRepository(db *gorm.DB) EventRepositoryInterface {
	return &eventRepository{DB: db}
}

func (dc *eventRepository) GetAllEvents() ([]models.Event, error) {
	var events []models.Event
	err := dc.DB.Find(&events).Error
	return events, err
}

func (dc *eventRepository) CreateEvent(event *models.Event) (*models.Event, error) {
	if err := dc.DB.Create(event).Error; err != nil {
		return nil, err
	}
	return event, nil
}
