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

type Event struct {
	gorm.Model
	Name         string    `json:"name" binding:"required,max=150"`
	Description  string    `json:"description" binding:"required,max=350"`
	VenueID      uint      `json:"venue_id"`
	OrganizerID  uint      `json:"organizer_id"`
	StartTime    time.Time `json:"start_time" binding:"required"`
	EndTime      time.Time `json:"end_time" binding:"required"`
	IsFree       bool      `json:"is_free"`
	TicketPrice  float64   `json:"ticket_price" binding:"required"`
	TotalTickets int       `json:"total_tickets"`
	ImageURL     string    `json:"image_url" binding:"required"`
}

type TicketType struct {
	gorm.Model
	EventID           uint      `json:"event_id"`
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
