package services

import (
	"ticketing-system/models"
	"ticketing-system/repositories"
)

// VenueServiceInterface defines the methods the VenueService must implement
type VenueServiceInterface interface {
	CreateVenue(venue *models.Venue) (*models.Venue, error)
	GetAllVenues() ([]models.Venue, error)
	GetSingleVenue(id uint) ([]models.Venue, error)
	UpdateVenue(id uint, venue models.Venue) (*models.Venue, error)
	DeleteVenue(id uint) (*models.Venue, error)
}

// VenueService implements VenueServiceInterface
type VenueService struct {
	VenueRepo repositories.VenueRepositoryInterface
}

func NewVenueService(venueRepo repositories.VenueRepositoryInterface) VenueServiceInterface {
	return &VenueService{VenueRepo: venueRepo}
}

func (s *VenueService) CreateVenue(venue *models.Venue) (*models.Venue, error) {
	return s.VenueRepo.CreateVenue(venue)
}

func (s *VenueService) GetAllVenues() ([]models.Venue, error) {
	return s.VenueRepo.GetAllVenues()
}

func (s *VenueService) GetSingleVenue(id uint) ([]models.Venue, error) {
	return s.VenueRepo.GetSingleVenue(id)
}

func (s *VenueService) UpdateVenue(id uint, venue models.Venue) (*models.Venue, error) {
	return s.VenueRepo.UpdateVenue(id, venue)
}

func (s *VenueService) DeleteVenue(id uint) (*models.Venue, error) {
	return s.VenueRepo.DeleteVenue(id)
}
