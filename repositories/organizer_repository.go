package repositories

import (
	"ticketing-system/models"

	"gorm.io/gorm"
)

type OrganizerRepositoryInterface interface {
	CreateOrganizer(organizer *models.Organizer) (*models.Organizer, error)
	GetAllOrganizers() ([]models.Organizer, error)
	GetSingleOrganizer(id uint) (*models.Organizer, error)
	UpdateOrganizer(id uint, organizer models.Organizer) (*models.Organizer, error)
	DeleteOrganizer(id uint) (*models.Organizer, error)
}

type organizerRepository struct {
	DB *gorm.DB
}

func NewOrganizerRepository(db *gorm.DB) OrganizerRepositoryInterface {
	return &organizerRepository{DB: db}
}
func (dc *organizerRepository) CreateOrganizer(organizer *models.Organizer) (*models.Organizer, error) {
	if err := dc.DB.Create(&organizer).Error; err != nil {
		return nil, err
	}

	return organizer, nil
}

func (dc *organizerRepository) GetAllOrganizers() ([]models.Organizer, error) {
	var organizers []models.Organizer
	if err := dc.DB.Find(&organizers).Error; err != nil {

		return organizers, err
	}

	return organizers, nil
}

func (dc *organizerRepository) GetSingleOrganizer(id uint) (*models.Organizer, error) {
	var organizer models.Organizer
	if err := dc.DB.Where("id = ?", id).First(&organizer).Error; err != nil {
		return nil, err
	}
	return &organizer, nil
}

func (dc *organizerRepository) UpdateOrganizer(id uint, organizer models.Organizer) (*models.Organizer, error) {

	if err := dc.DB.Model(&organizer).Where("id=?", id).Updates(map[string]interface{}{
		"first_name":           organizer.FirstName,
		"last_name":            organizer.LastName,
		"phone_number":         organizer.PhoneNumber,
		"email":                organizer.Email,
		"till_pay_bill_number": organizer.TillPayBillNumber,
		"account_reference":    organizer.AccountReference}).Scan(&organizer).Error; err != nil {
		return nil, err

	}
	return &organizer, nil
}

func (dc *organizerRepository) DeleteOrganizer(id uint) (*models.Organizer, error) {
	var organizer models.Organizer
	if err := dc.DB.Where("id = ?", id).Delete(&organizer).Error; err != nil {
		return nil, err
	}

	return &organizer, nil
}
