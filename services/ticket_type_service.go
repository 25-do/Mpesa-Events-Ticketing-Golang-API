package services

import (
	"ticketing-system/models"
	"ticketing-system/repositories"
)

// TicketTypeServiceInterface defines the methods the TicketTypeService must implement
type TicketTypeServiceInterface interface {
	CreateTicketType(ticketype *models.TicketType) (*models.TicketType, error)
	GetAllTicketTypes() ([]models.TicketType, error)
	GetSingleTicketType(id uint) ([]models.TicketType, error)
	UpdateTicketType(id uint, ticketype models.TicketType) (*models.TicketType, error)
	DeleteTicketType(id uint) (*models.TicketType, error)
}

// TicketTypeService implements TicketTypeServiceInterface
type TicketTypeService struct {
	TicketTypeRepo repositories.TicketTypeRepositoryInterface
}

func NewTicketTypeService(ticketypeRepo repositories.TicketTypeRepositoryInterface) TicketTypeServiceInterface {
	return &TicketTypeService{TicketTypeRepo: ticketypeRepo}
}

func (s *TicketTypeService) CreateTicketType(ticketype *models.TicketType) (*models.TicketType, error) {
	return s.TicketTypeRepo.CreateTicketType(ticketype)
}

func (s *TicketTypeService) GetAllTicketTypes() ([]models.TicketType, error) {
	return s.TicketTypeRepo.GetAllTicketTypes()
}

func (s *TicketTypeService) GetSingleTicketType(id uint) ([]models.TicketType, error) {
	return s.TicketTypeRepo.GetSingleTicketType(id)
}

func (s *TicketTypeService) UpdateTicketType(id uint, ticketype models.TicketType) (*models.TicketType, error) {
	return s.TicketTypeRepo.UpdateTicketType(id, ticketype)
}

func (s *TicketTypeService) DeleteTicketType(id uint) (*models.TicketType, error) {
	return s.TicketTypeRepo.DeleteTicketType(id)
}
