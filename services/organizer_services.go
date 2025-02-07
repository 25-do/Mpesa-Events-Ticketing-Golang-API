package services

import (
	"ticketing-system/models"
	"ticketing-system/repositories"
)

// OrganizerServiceInterface defines the methods the OrganizerService must implement
type OrganizerServiceInterface interface {
	CreateOrganizer(organizer *models.Organizer) (*models.Organizer, error)
	GetAllOrganizers() ([]models.Organizer, error)
	GetSingleOrganizer(id uint) ([]models.Organizer, error)
	UpdateOrganizer(id uint, organizer models.Organizer) (*models.Organizer, error)
	DeleteOrganizer(id uint) (*models.Organizer, error)
}

// OrganizerService implements OrganizerServiceInterface
type OrganizerService struct {
	OrganizerRepo repositories.OrganizerRepositoryInterface
}

func NewOrganizerService(organizerRepo repositories.OrganizerRepositoryInterface) OrganizerServiceInterface {
	return &OrganizerService{OrganizerRepo: organizerRepo}
}

func (s *OrganizerService) CreateOrganizer(organizer *models.Organizer) (*models.Organizer, error) {
	return s.OrganizerRepo.CreateOrganizer(organizer)
}

func (s *OrganizerService) GetAllOrganizers() ([]models.Organizer, error) {
	return s.OrganizerRepo.GetAllOrganizers()
}

func (s *OrganizerService) GetSingleOrganizer(id uint) ([]models.Organizer, error) {
	return s.OrganizerRepo.GetSingleOrganizer(id)
}

func (s *OrganizerService) UpdateOrganizer(id uint, organizer models.Organizer) (*models.Organizer, error) {
	return s.OrganizerRepo.UpdateOrganizer(id, organizer)
}

func (s *OrganizerService) DeleteOrganizer(id uint) (*models.Organizer, error) {
	return s.OrganizerRepo.DeleteOrganizer(id)
}
