package postgres

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/Khaled2049/ecommerce-app/internal/domain"
)

type ProductRepository struct {
	db *sql.DB
}

func NewProductRepository(db *sql.DB) *ProductRepository {
	return &ProductRepository{
		db: db,
	}
}

func (r *ProductRepository) Create(ctx context.Context, product *domain.ProductCreate) (*domain.Product, error) {
	query := `
		INSERT INTO products (name, description, price, stock)
		VALUES ($1, $2, $3, $4)
		RETURNING product_id, name, description, price, stock, created_at, updated_at`

	p := &domain.Product{}
	err := r.db.QueryRowContext(ctx, query,
		product.Name,
		product.Description,
		product.Price,
	).Scan(
		&p.ID,
		&p.Name,
		&p.Description,
		&p.Price,
		&p.CreatedAt,
		&p.UpdatedAt,
	)

	if err != nil {
		return nil, fmt.Errorf("error creating product: %w", err)
	}

	return p, nil
}

func (r *ProductRepository) GetByID(ctx context.Context, id uint64) (*domain.Product, error) {
	query := `
		SELECT product_id, name, description, price, stock, created_at, updated_at
		FROM products
		WHERE product_id = $1`

	p := &domain.Product{}
	err := r.db.QueryRowContext(ctx, query, id).Scan(
		&p.ID,
		&p.Name,
		&p.Description,
		&p.Price,
		&p.CreatedAt,
		&p.UpdatedAt,
	)

	if err != nil {
		return nil, fmt.Errorf("error getting product: %w", err)
	}

	return p, nil
}

func (r *ProductRepository) Update(ctx context.Context, product *domain.ProductUpdate) (*domain.Product, error) {
	query := `
		UPDATE products
		SET name = $1, description = $2, price = $3, stock = $4
		WHERE product_id = $5
		RETURNING product_id, name, description, price, stock, created_at, updated_at`

	p := &domain.Product{}
	err := r.db.QueryRowContext(ctx, query,
		product.Name,
		product.Description,
		product.Price,
		product.ID,
		product.UpdatedAt,
	).Scan(
		&p.ID,
		&p.Name,
		&p.Description,
		&p.Price,
		&p.CreatedAt,
		&p.UpdatedAt,
	)

	if err != nil {
		return nil, fmt.Errorf("error updating product: %w", err)
	}

	return p, nil
}

func (r *ProductRepository) Delete(ctx context.Context, id uint64) error {
	query := `DELETE FROM products WHERE product_id = $1`

	_, err := r.db.ExecContext(ctx, query, id)
	if err != nil {
		return fmt.Errorf("error deleting product: %w", err)
	}

	return nil
}

func (r *ProductRepository) List(ctx context.Context, limit, offset int) ([]*domain.Product, error) {
	query := `
		SELECT product_id, name, description, price, stock, created_at, updated_at
		FROM products
		LIMIT $1 OFFSET $2`

	rows, err := r.db.QueryContext(ctx, query, limit, offset)
	if err != nil {
		return nil, fmt.Errorf("error listing products: %w", err)
	}
	defer rows.Close()

	products := make([]*domain.Product, 0)
	for rows.Next() {
		p := &domain.Product{}
		err := rows.Scan(
			&p.ID,
			&p.Name,
			&p.Description,
			&p.Price,
			&p.CreatedAt,
			&p.UpdatedAt,
		)
		if err != nil {
			return nil, fmt.Errorf("error scanning product: %w", err)
		}
		products = append(products, p)
	}

	return products, nil
}
