package models

import (
	"time"

	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Name         string `json:"name" binding:"required,max=150"`
	Email        string `json:"client_email" binding:"required,email"`
	PasswordHash string `json:"password_hash"`
	Role         string `json:"role"`
}

type Venue struct {
	gorm.Model

	Name     string `json:"name" binding:"required,max=150"`
	Location string `json:"location" binding:"required,max=150"`
	Capacity int    `json:"capacity"`
}

type Organizer struct {
	gorm.Model
	FirstName   string `json:"first_name" binding:"required,max=150"`
	LastName    string `json:"last_name" binding:"required,max=150"`
	PhoneNumber string `json:"phone_number" binding:"required,max=150"`
	Email       string `json:"email" binding:"required,max=150"`
}
type Event struct {
	gorm.Model
	Name        string    `json:"name" binding:"required,max=150"`
	Description string    `json:"description" binding:"required,max=350"`
	VenueID     uint      `json:"venue_id" binding:"required"`
	OrganizerID uint      `json:"organizer_id" binding:"required"`
	StartTime   time.Time `json:"start_time" binding:"required"`
	EndTime     time.Time `json:"end_time" binding:"required"`
	IsFree      bool      `json:"is_free" binding:"required"`
	ImageURL    string    `json:"image_url" binding:"required"`
	TicketType  []TicketType
}

type TicketType struct {
	gorm.Model
	EventID           uint      `json:"event_id" binding:"required"`
	Name              string    `json:"name" binding:"required,max=150"`
	Price             float64   `json:"price" binding:"required"`
	AvailableQuantity int       `json:"available_quantity"`
	StartSaleTime     time.Time `json:"start_sale_time"`
	EndSaleTime       time.Time `json:"end_sale_time"`
}

type Ticket struct {
	gorm.Model
	TicketTypeID uint      `json:"tickettype_id"`
	UserID       uint      `json:"user_id"`
	PurchaseTime time.Time `json:"purchase_time"`
	Status       string    `json:"status"`
	SeatNumber   string    `json:"seat_number"`
}

type Payment struct {
	gorm.Model
	UserID        uint      `json:"user_id"`
	TicketID      uint      `json:"ticket_id"`
	Amount        float64   `json:"amount"`
	PaymentMethod string    `json:"payment_method"`
	Status        string    `json:"status"`
	PaymentTime   time.Time `json:"payment_time"`
}
