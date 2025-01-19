package domain

import (
	"encoding/json"
	"errors"
	"time"
)

type Customer struct {
	ID          int64           `json:"id"`
	Name        string          `json:"name"`
	Email       string          `json:"email"`
	Phone       string          `json:"phone"`
	CreatedAt   time.Time       `json:"created_at"`
	IsActive    bool            `json:"is_active"`
	Preferences json.RawMessage `json:"preferences,omitempty"`
}

type CustomerCreate struct {
    Name        string                 `json:"name" validate:"required"`
    Email       string                 `json:"email" validate:"required,email"`
    Phone       string                 `json:"phone" validate:"required"`
    Preferences map[string]interface{} `json:"preferences,omitempty"`
}

type CustomerUpdate struct {
	ID          int64           `json:"id" validate:"required"`
	Name        string          `json:"name" validate:"required"`
	Email       string          `json:"email" validate:"required,email"`
	Phone       string          `json:"phone" validate:"required"`
	Preferences json.RawMessage `json:"preferences,omitempty"`
}

var (
	ErrCustomerNotFound   = errors.New("customer not found")
	ErrEmailAlreadyExists = errors.New("email already exists")
)
