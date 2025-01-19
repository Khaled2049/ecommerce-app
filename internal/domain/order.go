package domain

import "time"

// OrderStatus represents the current state of an order
type OrderStatus string

const (
	OrderStatusPending   OrderStatus = "pending"
	OrderStatusPaid      OrderStatus = "paid"
	OrderStatusShipped   OrderStatus = "shipped"
	OrderStatusDelivered OrderStatus = "delivered"
	OrderStatusCanceled  OrderStatus = "canceled"
)

// Order represents the order domain model
type Order struct {
	OrderID     uint64      `json:"order_id" db:"order_id"`
	CustomerID  uint64      `json:"customer_id" db:"customer_id"`
	TotalAmount float64     `json:"total_amount" db:"total_amount"`
	Status      OrderStatus `json:"status" db:"status"`
	OrderDate   time.Time   `json:"order_date" db:"order_date"`
}

// OrderCreate represents the data needed to create a new order
type OrderCreate struct {
	CustomerID  uint64      `json:"customer_id" validate:"required"`
	TotalAmount float64     `json:"total_amount" validate:"required,gt=0"`
	Status      OrderStatus `json:"status"`
}

