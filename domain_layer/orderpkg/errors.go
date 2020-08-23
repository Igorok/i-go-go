package orderpkg

import "errors"

var (
	ErrOrderWrongData = errors.New("order_wrong_data")
	ErrOrderNotFound  = errors.New("order_not_found")
)
