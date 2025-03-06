package domain

import "errors"

type Payment struct {
	ID          int64   `json:"id"`
	OrderID     int64   `json:"order_id"`
	Amount      float64 `json:"amount"`
	PaymentType string  `json:"payment_type"`
	CreatedAt   string  `json:"created_at"`
}

type PaymentCreate struct {
	OrderID     int64   `json:"order_id" validate:"required"`
	Amount      float64 `json:"amount" validate:"required"`
	PaymentType string  `json:"payment_type" validate:"required"`
}

type PaymentUpdate struct {
	ID          int64   `json:"id" validate:"required"`
	OrderID     int64   `json:"order_id" validate:"required"`
	Amount      float64 `json:"amount" validate:"required"`
	PaymentType string  `json:"payment_type" validate:"required"`
}

var (
	ErrPaymentNotFound = errors.New("payment not found")
)
