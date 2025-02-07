package repositories

import (
	"ticketing-system/models"

	"gorm.io/gorm"
)

type VenueRepositoryInterface interface {
	CreateVenue(venue *models.Venue) (*models.Venue, error)
	GetAllVenues() ([]models.Venue, error)
	GetSingleVenue(id uint) ([]models.Venue, error)
	UpdateVenue(id uint, venue models.Venue) (*models.Venue, error)
	DeleteVenue(id uint) (*models.Venue, error)
}

type venueRepository struct {
	DB *gorm.DB
}

func NewVenueRepository(db *gorm.DB) VenueRepositoryInterface {
	return &venueRepository{DB: db}
}
func (dc *venueRepository) CreateVenue(venue *models.Venue) (*models.Venue, error) {
	if err := dc.DB.Create(&venue).Error; err != nil {
		return nil, err
	}

	return venue, nil
}

func (dc *venueRepository) GetAllVenues() ([]models.Venue, error) {
	var venues []models.Venue
	if err := dc.DB.Find(&venues).Error; err != nil {

		return venues, err
	}

	return venues, nil
}

func (dc *venueRepository) GetSingleVenue(id uint) ([]models.Venue, error) {
	var venue []models.Venue
	if err := dc.DB.Where("id = ?", id).First(&venue).Error; err != nil {
		return nil, err
	}
	return venue, nil
}

func (dc *venueRepository) UpdateVenue(id uint, venue models.Venue) (*models.Venue, error) {

	if err := dc.DB.Model(&venue).Where("id=?", id).Updates(map[string]interface{}{
		"name":     venue.Name,
		"location": venue.Location,
		"capacity": venue.Capacity}).Scan(&venue).Error; err != nil {
		return nil, err

	}
	return &venue, nil
}

func (dc *venueRepository) DeleteVenue(id uint) (*models.Venue, error) {
	var venue models.Venue
	if err := dc.DB.Where("id = ?", id).Delete(&venue).Error; err != nil {
		return nil, err
	}

	return &venue, nil
}
