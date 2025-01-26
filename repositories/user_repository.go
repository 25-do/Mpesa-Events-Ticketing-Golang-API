package repositories

import (
	"ticketing-system/db"
	"ticketing-system/models"
)

func CreateUser(user *models.User) error {
	return db.DB.Create(user).Error
}
