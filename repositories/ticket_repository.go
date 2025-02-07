package repositories

import (
	"ticketing-system/models"

	"gorm.io/gorm"
)

type TicketRepositoryInterface interface {
	CreateTicket(ticket *models.Ticket) (*models.Ticket, error)
	GetAllTickets() ([]models.Ticket, error)
	GetSingleTicket(id uint) ([]models.Ticket, error)
	UpdateTicket(id uint, ticket models.Ticket) (*models.Ticket, error)
	DeleteTicket(id uint) (*models.Ticket, error)
}

type ticketRepository struct {
	DB *gorm.DB
}

func NewTicketRepository(db *gorm.DB) TicketRepositoryInterface {
	return &ticketRepository{DB: db}
}
func (dc *ticketRepository) CreateTicket(ticket *models.Ticket) (*models.Ticket, error) {
	if err := dc.DB.Create(&ticket).Error; err != nil {
		return nil, err
	}

	return ticket, nil
}

func (dc *ticketRepository) GetAllTickets() ([]models.Ticket, error) {
	var tickets []models.Ticket
	if err := dc.DB.Find(&tickets).Error; err != nil {

		return tickets, err
	}

	return tickets, nil
}

func (dc *ticketRepository) GetSingleTicket(id uint) ([]models.Ticket, error) {
	var ticket []models.Ticket
	if err := dc.DB.Where("id = ?", id).First(&ticket).Error; err != nil {
		return nil, err
	}
	return ticket, nil
}

func (dc *ticketRepository) UpdateTicket(id uint, ticket models.Ticket) (*models.Ticket, error) {

	if err := dc.DB.Model(&ticket).Where("id=?", id).Updates(map[string]interface{}{
		"purchase_time": ticket.PurchaseTime,
		"status":        ticket.Status,
		"seat_number":   ticket.SeatNumber}).Scan(&ticket).Error; err != nil {
		return nil, err

	}
	return &ticket, nil
}

func (dc *ticketRepository) DeleteTicket(id uint) (*models.Ticket, error) {
	var ticket models.Ticket
	if err := dc.DB.Where("id = ?", id).Delete(&ticket).Error; err != nil {
		return nil, err
	}

	return &ticket, nil
}
