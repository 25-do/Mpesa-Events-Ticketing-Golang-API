package services

import (
	"ticketing-system/models"
	"ticketing-system/repositories"
)

func CreateEvent(event *models.Event) error {
	return repositories.CreateEvent(event)
}
