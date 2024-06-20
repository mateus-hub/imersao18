package domain

import "errors"

type TicketType string

// Errors
var (
	ErrInvalidTicketType = errors.New("invalid ticket type")
)

const (
	TicketTypeHalf TicketType = "half" // Half-price ticket
	TicketTypeFull TicketType = "full" // Full-price ticket
)

// Ticket represents a ticket for an event.
type Ticket struct {
	Id         string
	EventId    string
	Spot       *Spot
	TicketType TicketType
	Price      float64
}

// IsValidTicketType checks if a ticket type is valid.
func IsValidTicketType(ticketType TicketType) bool {
	return ticketType == TicketTypeHalf || ticketType == TicketTypeFull
}

// CalculatePrice calculates the price based on the ticket type.
func (t *Ticket) CalculatePrice() {
	if t.TicketType == TicketTypeHalf {
		t.Price /= 2
	}
}

// Validate checks if the ticket data is valid.
func (t *Ticket) Validate() error {
	if t.Price <= 0 {
		return ErrInvalidTicketType
	}
	return nil
}
