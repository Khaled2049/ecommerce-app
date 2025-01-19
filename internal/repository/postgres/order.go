package postgres

import (
	"context"
	"database/sql"
	"errors"
	"fmt"

	"github.com/Khaled2049/ecommerce-app/internal/domain"
)

type OrderRepository struct {
	db *sql.DB
}

func NewOrderRepository(db *sql.DB) *OrderRepository {
	return &OrderRepository{
		db: db,
	}
}

func (r *OrderRepository) Create(ctx context.Context, o *domain.OrderCreate) (*domain.Order, error) {
	query := `
		INSERT INTO orders (customer_id, total_amount, status, order_date)
		VALUES ($1, $2, $3, CURRENT_TIMESTAMP)
		RETURNING order_id, customer_id, total_amount, status, order_date`

	order := &domain.Order{}
	err := r.db.QueryRowContext(ctx, query,
		o.CustomerID,
		o.TotalAmount,
		domain.OrderStatusPending,
	).Scan(
		&order.OrderID,
		&order.CustomerID,
		&order.TotalAmount,
		&order.Status,
		&order.OrderDate,
	)

	if err != nil {
		return nil, fmt.Errorf("error creating order: %w", err)
	}

	return order, nil
}

func (r *OrderRepository) GetByID(ctx context.Context, orderID uint64) (*domain.Order, error) {
	order := &domain.Order{}
	query := `
		SELECT order_id, customer_id, total_amount, status, order_date
		FROM orders
		WHERE order_id = $1`

	err := r.db.QueryRowContext(ctx, query, orderID).Scan(
		&order.OrderID,
		&order.CustomerID,
		&order.TotalAmount,
		&order.Status,
		&order.OrderDate,
	)

	if err == sql.ErrNoRows {
		return nil, errors.New("order not found")
	}

	if err != nil {
		return nil, err
	}

	return order, nil
}

func (r *OrderRepository) GetByCustomerID(ctx context.Context, customerID uint64) ([]*domain.Order, error) {
	query := `
		SELECT order_id, customer_id, total_amount, status, order_date
		FROM orders
		WHERE customer_id = $1
		ORDER BY order_date DESC`

	rows, err := r.db.QueryContext(ctx, query, customerID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var orders []*domain.Order
	for rows.Next() {
		order := &domain.Order{}
		err := rows.Scan(
			&order.OrderID,
			&order.CustomerID,
			&order.TotalAmount,
			&order.Status,
			&order.OrderDate,
		)
		if err != nil {
			return nil, err
		}
		orders = append(orders, order)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return orders, nil
}

func (r *OrderRepository) Update(ctx context.Context, order *domain.Order) (*domain.Order, error) {
	query := `
		UPDATE orders 
		SET customer_id = $1, total_amount = $2, status = $3, order_date = $4
		WHERE order_id = $5
		RETURNING order_id, customer_id, total_amount, status, order_date`

	err := r.db.QueryRowContext(ctx, query,
		order.CustomerID,
		order.TotalAmount,
		order.Status,
		order.OrderDate,
		order.OrderID,
	).Scan(&order.OrderID, &order.CustomerID, &order.TotalAmount, &order.Status, &order.OrderDate)

	if err == sql.ErrNoRows {
		return nil, errors.New("order not found")
	}
	if err != nil {
		return nil, err
	}

	return order, nil
}

func (r *OrderRepository) Delete(ctx context.Context, id uint64) error {
	query := `DELETE FROM orders WHERE order_id = $1`

	result, err := r.db.ExecContext(ctx, query, id)
	if err != nil {
		return err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return err
	}

	if rowsAffected == 0 {
		return errors.New("order not found")
	}

	return nil
}

func (r *OrderRepository) List(ctx context.Context, limit, offset int) ([]*domain.Order, error) {
	query := `
		SELECT order_id, customer_id, total_amount, status, order_date
		FROM orders
		ORDER BY order_date DESC`

	rows, err := r.db.QueryContext(ctx, query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var orders []*domain.Order
	for rows.Next() {
		order := &domain.Order{}
		err := rows.Scan(
			&order.OrderID,
			&order.CustomerID,
			&order.TotalAmount,
			&order.Status,
			&order.OrderDate,
		)
		if err != nil {
			return nil, err
		}
		orders = append(orders, order)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return orders, nil
}
