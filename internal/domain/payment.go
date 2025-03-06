package domain

import "errors"

type Payment struct {
	ID            int64   `json:"payment_id"`
	OrderID       int64   `json:"order_id"`
	PaymentID     string  `json:"payment_method_id"`
	Amount        float64 `json:"amount"`
	Status        string  `json:"status"`
	PaymentType   string  `json:"payment_type"`
	PaymentDate   string  `json:"payment_date"`
	TransactionId string  `json:"transaction_id"`
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
