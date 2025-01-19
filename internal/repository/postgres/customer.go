package postgres

import (
	"context"
	"database/sql"
	"encoding/json"
	"fmt"

	"github.com/Khaled2049/ecommerce-app/internal/domain"
)

type customerRepository struct {
	db *sql.DB
}

func NewCustomerRepository(db *sql.DB) *customerRepository {
	return &customerRepository{db: db}
}

func (r *customerRepository) Create(ctx context.Context, c *domain.CustomerCreate) (*domain.Customer, error) {
    // Convert preferences to JSON
    preferencesJSON, err := json.Marshal(c.Preferences)
    if err != nil {
        return nil, fmt.Errorf("error marshaling preferences: %w", err)
    }

    query := `
        INSERT INTO customers (name, email, phone, preferences)
        VALUES ($1, $2, $3, $4)
        RETURNING customer_id, name, email, phone, created_at, is_active, preferences`
    
    customer := &domain.Customer{}
    err = r.db.QueryRowContext(
        ctx, 
        query,
        c.Name,
        c.Email,
        c.Phone,
        preferencesJSON, // Now passing properly formatted JSON
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
        return nil, fmt.Errorf("error creating customer: %w", err)
    }
    
    return customer, nil
}

func (r *customerRepository) List(ctx context.Context, limit, offset int) ([]domain.Customer, error) {
	query := `
		SELECT customer_id, name, email, phone, created_at, is_active
		FROM customers
		ORDER BY customer_id DESC
		LIMIT $1 OFFSET $2`

	rows, err := r.db.QueryContext(ctx, query, limit, offset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var customers []domain.Customer
	for rows.Next() {
		c := domain.Customer{}
		err := rows.Scan(
			&c.ID,
			&c.Name,
			&c.Email,
			&c.Phone,
			&c.CreatedAt,
			&c.IsActive,
			// &c.Preferences,
		)
		if err != nil {
			return nil, err
		}

		customers = append(customers, c)
	}

	return customers, nil
}

func (r *customerRepository) GetByID(ctx context.Context, id int64) (*domain.Customer, error) {
	query := `
		SELECT customer_id, name, email, phone, created_at, is_active
		FROM customers
		WHERE customer_id = $1`

	c := &domain.Customer{}
	err := r.db.QueryRowContext(ctx, query, id).Scan(
		&c.ID,
		&c.Name,
		&c.Email,
		&c.Phone,
		&c.CreatedAt,
		&c.IsActive,
		// &c.Preferences,
	)
	if err != nil {
		return nil, err
	}

	return c, nil
}

func (r *customerRepository) Update(ctx context.Context, c *domain.CustomerUpdate) (*domain.Customer, error) {
	query := `
        UPDATE customers 
        SET name = $1, email = $2, phone = $3, preferences = $4
        WHERE customer_id = $5
        RETURNING customer_id, name, email, phone, created_at, is_active, preferences`

	var updated domain.Customer
	err := r.db.QueryRowContext(
		ctx,
		query,
		c.Name,
		c.Email,
		c.Phone,
		c.Preferences,
		c.ID,
	).Scan(
		&updated.ID,
		&updated.Name,
		&updated.Email,
		&updated.Phone,
		&updated.CreatedAt,
		&updated.IsActive,
		&updated.Preferences,
	)
	if err == sql.ErrNoRows {
		return nil, domain.ErrCustomerNotFound
	}
	if err != nil {
		return nil, fmt.Errorf("failed to update customer: %w", err)
	}

	return &updated, nil
}

func (r *customerRepository) GetByEmail(ctx context.Context, email string) (*domain.Customer, error) {
	query := `
		SELECT customer_id, name, email, phone, created_at, is_active
		FROM customers
		WHERE email = $1`

	c := &domain.Customer{}
	err := r.db.QueryRowContext(ctx, query, email).Scan(
		&c.ID,
		&c.Name,
		&c.Email,
		&c.Phone,
		&c.CreatedAt,
		&c.IsActive,
		// &c.Preferences,
	)
	if err == sql.ErrNoRows {
		return nil, domain.ErrCustomerNotFound
	}
	if err != nil {
		return nil, err
	}

	return c, nil
}

func (r *customerRepository) Delete(ctx context.Context, id int64) error {
	query := `
        DELETE FROM customers
        WHERE customer_id = $1
        RETURNING customer_id` // Use RETURNING to check if row existed

	result := r.db.QueryRowContext(ctx, query, id)
	if err := result.Scan(&id); err != nil {
		if err == sql.ErrNoRows {
			return domain.ErrCustomerNotFound
		}
		return fmt.Errorf("failed to delete customer: %w", err)
	}

	return nil
}
