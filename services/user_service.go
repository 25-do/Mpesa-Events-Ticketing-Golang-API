package services

import (
	"errors"
	"ticketing-system/models"
	"ticketing-system/repositories"
)

// UserServiceInterface defines the methods the UserService must implement
type UserServiceInterface interface {
	GetAllUsers() ([]models.User, error)
	GetUserByID(id uint) (*models.User, error)
}

// UserService implements UserServiceInterface
type UserService struct {
	UserRepo repositories.UserRepositoryInterface
}

func NewUserService(userRepo repositories.UserRepositoryInterface) UserServiceInterface {
	return &UserService{UserRepo: userRepo}
}

func (s *UserService) GetAllUsers() ([]models.User, error) {
	return s.UserRepo.FindAll()
}

func (s *UserService) GetUserByID(id uint) (*models.User, error) {
	user, err := s.UserRepo.FindByID(id)
	if err != nil {
		return nil, errors.New("user not found")
	}
	return user, nil
}
