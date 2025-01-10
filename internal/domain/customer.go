package domain

import (
	"encoding/json"
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
	Name        string          `json:"name" validate:"required"`
	Email       string          `json:"email" validate:"required,email"`
	Phone       string          `json:"phone" validate:"required"`
	Preferences json.RawMessage `json:"preferences,omitempty"`
}
