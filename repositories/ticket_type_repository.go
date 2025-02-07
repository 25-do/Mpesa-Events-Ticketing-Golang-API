package repositories

import (
	"ticketing-system/models"

	"gorm.io/gorm"
)

type TicketTypeRepositoryInterface interface {
	CreateTicketType(ticketype *models.TicketType) (*models.TicketType, error)
	GetAllTicketTypes() ([]models.TicketType, error)
	GetSingleTicketType(id uint) ([]models.TicketType, error)
	UpdateTicketType(id uint, ticketype models.TicketType) (*models.TicketType, error)
	DeleteTicketType(id uint) (*models.TicketType, error)
}

type ticketypeRepository struct {
	DB *gorm.DB
}

func NewTicketTypeRepository(db *gorm.DB) TicketTypeRepositoryInterface {
	return &ticketypeRepository{DB: db}
}
func (dc *ticketypeRepository) CreateTicketType(ticketype *models.TicketType) (*models.TicketType, error) {
	if err := dc.DB.Create(&ticketype).Error; err != nil {
		return nil, err
	}

	return ticketype, nil
}

func (dc *ticketypeRepository) GetAllTicketTypes() ([]models.TicketType, error) {
	var ticketypes []models.TicketType
	if err := dc.DB.Find(&ticketypes).Error; err != nil {

		return ticketypes, err
	}

	return ticketypes, nil
}

func (dc *ticketypeRepository) GetSingleTicketType(id uint) ([]models.TicketType, error) {
	var ticketype []models.TicketType
	if err := dc.DB.Where("id = ?", id).First(&ticketype).Error; err != nil {
		return nil, err
	}
	return ticketype, nil
}

func (dc *ticketypeRepository) UpdateTicketType(id uint, ticketype models.TicketType) (*models.TicketType, error) {

	if err := dc.DB.Model(&ticketype).Where("id=?", id).Updates(map[string]interface{}{
		"name":               ticketype.Name,
		"price":              ticketype.Price,
		"available_quantity": ticketype.AvailableQuantity,
		"start_sale_time":    ticketype.StartSaleTime,
		"end_sale_time":      ticketype.EndSaleTime}).Scan(&ticketype).Error; err != nil {
		return nil, err

	}
	return &ticketype, nil
}

func (dc *ticketypeRepository) DeleteTicketType(id uint) (*models.TicketType, error) {
	var ticketype models.TicketType
	if err := dc.DB.Where("id = ?", id).Delete(&ticketype).Error; err != nil {
		return nil, err
	}

	return &ticketype, nil
}
