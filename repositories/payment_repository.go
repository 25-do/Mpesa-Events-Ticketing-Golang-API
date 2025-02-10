package repositories

import (
	"ticketing-system/models"

	"gorm.io/gorm"
)

type PaymentRepositoryInterface interface {
	GetAllPayments() ([]models.Payment, error)
	CreatePayment(payment *models.Payment) (*models.Payment, error)
	GetSinglePayment(id uint) ([]models.Payment, error)
	UpdatePayment(id uint, payment models.Payment) (*models.Payment, error)
	DeletePayment(id uint) (*models.Payment, error)
}

type paymentRepository struct {
	DB *gorm.DB
}

func NewPaymentRepository(db *gorm.DB) PaymentRepositoryInterface {
	return &paymentRepository{DB: db}
}

func (dc *paymentRepository) GetAllPayments() ([]models.Payment, error) {
	var payments []models.Payment
	err := dc.DB.Find(&payments).Error
	return payments, err
}

func (dc *paymentRepository) CreatePayment(payment *models.Payment) (*models.Payment, error) {
	if err := dc.DB.Create(payment).Error; err != nil {
		return nil, err
	}
	return payment, nil
}

func (dc *paymentRepository) GetSinglePayment(id uint) ([]models.Payment, error) {
	var payment []models.Payment
	if err := dc.DB.Where("id = ?", id).First(&payment).Error; err != nil {
		return nil, err
	}
	return payment, nil
}

func (dc *paymentRepository) UpdatePayment(id uint, payment models.Payment) (*models.Payment, error) {

	if err := dc.DB.Model(&payment).Where("id=?", id).Updates(map[string]interface{}{
		"amount":         payment.Amount,
		"payment_method": payment.PaymentMethod,
		"payment_time":   payment.PaymentTime}).Scan(&payment).Error; err != nil {
		return nil, err

	}
	return &payment, nil
}

func (dc *paymentRepository) DeletePayment(id uint) (*models.Payment, error) {
	var payment models.Payment
	if err := dc.DB.Where("id = ?", id).Delete(&payment).Error; err != nil {
		return nil, err
	}

	return &payment, nil
}

func (dc *paymentRepository) OrganizerId(id uint) (*models.Payment, error) {
	var payment models.Payment
	if err := dc.DB.Where("organizer_id = ?", id).First(&payment).Scan(&payment).Error; err != nil {
		return nil, err
	}
	return &payment, nil
}
