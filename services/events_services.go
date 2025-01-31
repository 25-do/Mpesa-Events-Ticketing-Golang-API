package services

import (
	"ticketing-system/models"
	"ticketing-system/repositories"
)

type EventServiceInterface interface {
	GetAllEvents() ([]models.Event, error)
	CreateEvent(event *models.Event) (*models.Event, error)
}

type EventService struct {
	EventRepo repositories.EventRepositoryInterface
}

func NewEventService(eventRepo repositories.EventRepositoryInterface) EventServiceInterface {
	return &EventService{EventRepo: eventRepo}
}

func (dc *EventService) GetAllEvents() ([]models.Event, error) {
	return dc.EventRepo.GetAllEvents()

}

func (dc *EventService) CreateEvent(event *models.Event) (*models.Event, error) {
	return dc.EventRepo.CreateEvent(event)
}
