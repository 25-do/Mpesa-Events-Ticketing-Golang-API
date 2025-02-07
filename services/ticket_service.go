package services

import (
	"ticketing-system/models"
	"ticketing-system/repositories"
)

// TicketServiceInterface defines the methods the TicketService must implement
type TicketServiceInterface interface {
	CreateTicket(ticket *models.Ticket) (*models.Ticket, error)
	GetAllTickets() ([]models.Ticket, error)
	GetSingleTicket(id uint) ([]models.Ticket, error)
	UpdateTicket(id uint, ticket models.Ticket) (*models.Ticket, error)
	DeleteTicket(id uint) (*models.Ticket, error)
}

// TicketService implements TicketServiceInterface
type TicketService struct {
	TicketRepo repositories.TicketRepositoryInterface
}

func NewTicketService(ticketRepo repositories.TicketRepositoryInterface) TicketServiceInterface {
	return &TicketService{TicketRepo: ticketRepo}
}

func (s *TicketService) CreateTicket(ticket *models.Ticket) (*models.Ticket, error) {
	return s.TicketRepo.CreateTicket(ticket)
}

func (s *TicketService) GetAllTickets() ([]models.Ticket, error) {
	return s.TicketRepo.GetAllTickets()
}

func (s *TicketService) GetSingleTicket(id uint) ([]models.Ticket, error) {
	return s.TicketRepo.GetSingleTicket(id)
}

func (s *TicketService) UpdateTicket(id uint, ticket models.Ticket) (*models.Ticket, error) {
	return s.TicketRepo.UpdateTicket(id, ticket)
}

func (s *TicketService) DeleteTicket(id uint) (*models.Ticket, error) {
	return s.TicketRepo.DeleteTicket(id)
}
