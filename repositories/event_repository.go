package repositories

import (
	"ticketing-system/db"
	"ticketing-system/models"
)

func CreateEvent(event *models.Event) error {
	return db.DB.Create(&event).Error
}
