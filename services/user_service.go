package services

import (
	"ticketing-system/models"
	"ticketing-system/repositories"
)

func CreateUser(user *models.User) error {
	return repositories.CreateUser(user)
}
