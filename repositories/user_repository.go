package repositories

import (
	"ticketing-system/models"

	"gorm.io/gorm"
)

// UserRepositoryInterface defines the methods the UserRepository must implement
type UserRepositoryInterface interface {
	FindAll() ([]models.User, error)
	FindByID(id uint) (*models.User, error)
	CreateUser(user *models.User) (*models.User, error)
}

// UserRepository implements UserRepositoryInterface
type UserRepository struct {
	DB *gorm.DB
}

func NewUserRepository(db *gorm.DB) UserRepositoryInterface {
	return &UserRepository{DB: db}
}

func (r *UserRepository) FindAll() ([]models.User, error) {
	var users []models.User
	err := r.DB.Find(&users).Error
	return users, err
}

func (r *UserRepository) FindByID(id uint) (*models.User, error) {
	var user models.User
	err := r.DB.First(&user, id).Error
	return &user, err
}

func (r *UserRepository) CreateUser(user *models.User) (*models.User, error) {
	if err := r.DB.Create(user).Error; err != nil {
		return nil, err
	}
	return user, nil
}
