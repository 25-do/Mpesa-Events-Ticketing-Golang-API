package services

import (
	"ticketing-system/models"
	"ticketing-system/repositories"
)

type PaymentServiceInterface interface {
	GetAllPayments() ([]models.Payment, error)
	CreatePayment(payment *models.Payment) (*models.Payment, error)
	GetSinglePayment(id uint) ([]models.Payment, error)
	UpdatePayment(id uint, payment models.Payment) (*models.Payment, error)
	DeletePayment(id uint) (*models.Payment, error)
}

type PaymentService struct {
	PaymentRepo repositories.PaymentRepositoryInterface
}

func NewPaymentService(paymentRepo repositories.PaymentRepositoryInterface) PaymentServiceInterface {
	return &PaymentService{PaymentRepo: paymentRepo}
}

func (dc *PaymentService) GetAllPayments() ([]models.Payment, error) {
	return dc.PaymentRepo.GetAllPayments()

}

func (dc *PaymentService) CreatePayment(payment *models.Payment) (*models.Payment, error) {
	return dc.PaymentRepo.CreatePayment(payment)
}

func (s *PaymentService) GetSinglePayment(id uint) ([]models.Payment, error) {
	return s.PaymentRepo.GetSinglePayment(id)
}

func (s *PaymentService) UpdatePayment(id uint, payment models.Payment) (*models.Payment, error) {
	return s.PaymentRepo.UpdatePayment(id, payment)
}

func (s *PaymentService) DeletePayment(id uint) (*models.Payment, error) {
	return s.PaymentRepo.DeletePayment(id)
}
