package postgres

import (
	"context"
	"database/sql"
	"time"

	"github.com/Khaled2049/ecommerce-app/internal/domain"
)

type customerRepository struct {
	db *sql.DB
}

func NewCustomerRepository(db *sql.DB) *customerRepository {
	return &customerRepository{db: db}
}

func (r *customerRepository) Create(ctx context.Context, c *domain.CustomerCreate) (*domain.Customer, error) {
	query := `
        INSERT INTO customers (name, email, phone, preferences)
        VALUES ($1, $2, $3, $4)
        RETURNING id, name, email, phone, created_at, is_active, preferences`

	customer := &domain.Customer{}
	err := r.db.QueryRowContext(
		ctx,
		query,
		c.Name,
		c.Email,
		c.Phone,
		c.Preferences,
	).Scan(
		&customer.ID,
		&customer.Name,
		&customer.Email,
		&customer.Phone,
		&customer.CreatedAt,
		&customer.IsActive,
		&customer.Preferences,
	)

	if err != nil {
		return nil, err
	}

	return customer, nil
}

func (r *customerRepository) List(ctx context.Context, limit, offset int) ([]domain.Customer, error) {
	// query := `
	// 	SELECT id, name, email, phone, created_at, is_active, preferences
	// 	FROM customers
	// 	ORDER BY id DESC
	// 	LIMIT $1 OFFSET $2`

	// rows, err := r.db.QueryContext(ctx, query, limit, offset)
	// if err != nil {
	// 	return nil, err
	// }
	// defer rows.Close()

	// var customers []domain.Customer
	// for rows.Next() {
	// 	c := domain.Customer{}
	// 	err := rows.Scan(
	// 		&c.ID,
	// 		&c.Name,
	// 		&c.Email,
	// 		&c.Phone,
	// 		&c.CreatedAt,
	// 		&c.IsActive,
	// 		&c.Preferences,
	// 	)
	// 	if err != nil {
	// 		return nil, err
	// 	}

	// 	customers = append(customers, c)
	// }

	customers := []domain.Customer{
		{
			ID:          1,
			Name:        "John Doe",
			Email:       "1@2.com",
			CreatedAt:   time.Date(2021, time.January, 1, 0, 0, 0, 0, time.UTC),
			IsActive:    true,
			Preferences: nil,
		},
	}

	return customers, nil
}
