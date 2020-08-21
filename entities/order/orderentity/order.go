package orderentity

import "time"

// Order is entity of order
type Order struct {
	ID           string     `validate:"omitempty"`
	IDCl         string     `validate:"required"`
	IDEst        string     `validate:"required"`
	Cart         []CartItem `validate:"gt=0"`
	Price        int        `validate:"required,gt=0"`
	DateCreate   time.Time  `validate:"required"`
	DateCustomer time.Time  `validate:"required"`
	Comment      string     `validate:"omitempty"`
	Address      string     `validate:"required"`
	Status       string     `validate:"required"`
}
