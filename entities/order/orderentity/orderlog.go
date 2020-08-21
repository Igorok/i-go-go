package orderentity

import "time"

// OrderLog is entity of history of events from order live circle
type OrderLog struct {
	ID     string    `validate:"omitempty"`
	IDOrd  string    `validate:"required"`
	IDUsr  string    `validate:"omitempty"`
	Date   time.Time `validate:"required"`
	Status string    `validate:"required"`
}
