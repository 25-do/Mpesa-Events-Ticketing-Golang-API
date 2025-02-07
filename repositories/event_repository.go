package repositories

import (
	"ticketing-system/models"

	"gorm.io/gorm"
)

type EventRepositoryInterface interface {
	GetAllEvents() ([]models.Event, error)
	CreateEvent(event *models.Event) (*models.Event, error)
	GetSingleEvent(id uint) ([]models.Event, error)
	UpdateEvent(id uint, event models.Event) (*models.Event, error)
	DeleteEvent(id uint) (*models.Event, error)
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

func (dc *eventRepository) GetSingleEvent(id uint) ([]models.Event, error) {
	var event []models.Event
	if err := dc.DB.Where("id = ?", id).First(&event).Error; err != nil {
		return nil, err
	}
	return event, nil
}

func (dc *eventRepository) UpdateEvent(id uint, event models.Event) (*models.Event, error) {

	if err := dc.DB.Model(&event).Where("id=?", id).Updates(map[string]interface{}{
		"name":         event.Name,
		"description":  event.Description,
		"venue_id":     event.VenueID,
		"organizer_id": event.OrganizerID,
		"start_time":   event.StartTime,
		"end_time":     event.EndTime,
		"is_free":      event.IsFree,
		"image_url":    event.ImageURL}).Scan(&event).Error; err != nil {
		return nil, err

	}
	return &event, nil
}

func (dc *eventRepository) DeleteEvent(id uint) (*models.Event, error) {
	var event models.Event
	if err := dc.DB.Where("id = ?", id).Delete(&event).Error; err != nil {
		return nil, err
	}

	return &event, nil
}
