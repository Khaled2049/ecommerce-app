package postgres

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/Khaled2049/ecommerce-app/internal/domain"
)

type paymentRepository struct {
	db *sql.DB
}

func NewPaymentRepository(db *sql.DB) *paymentRepository {
	return &paymentRepository{db: db}
}

func (r *paymentRepository) Create(ctx context.Context, p *domain.PaymentCreate) (*domain.Payment, error) {
	query := `
		INSERT INTO payments (order_id, amount, payment_type)
		VALUES ($1, $2, $3)
		RETURNING payment_id, order_id, amount, payment_type, created_at
	`

	payment := &domain.Payment{}
	err := r.db.QueryRowContext(
		ctx,
		query,
		p.OrderID,
		p.Amount,
		p.PaymentType,
	).Scan(
		&payment.ID,
		&payment.OrderID,
		&payment.Amount,
		&payment.PaymentType,
		&payment.PaymentDate,
	)

	if err != nil {
		return nil, fmt.Errorf("error creating payment: %w", err)
	}

	return payment, nil
}

func (r *paymentRepository) GetByID(ctx context.Context, paymentID int64) (*domain.Payment, error) {
	payment := &domain.Payment{}
	query := `
		SELECT payment_id, order_id, amount, payment_type, created_at
		FROM payments
		WHERE payment_id = $1
	`

	err := r.db.QueryRowContext(ctx, query, paymentID).Scan(
		&payment.ID,
		&payment.OrderID,
		&payment.Amount,
		&payment.PaymentType,
		&payment.PaymentDate,
	)

	if err != nil {
		return nil, fmt.Errorf("error getting payment: %w", err)
	}

	return payment, nil
}

func (r *paymentRepository) Update(ctx context.Context, p *domain.PaymentUpdate) (*domain.Payment, error) {
	query := `
		UPDATE payments
		SET order_id = $1, amount = $2, payment_type = $3
		WHERE payment_id = $4
		RETURNING payment_id, order_id, amount, payment_type, created_at
	`

	payment := &domain.Payment{}
	err := r.db.QueryRowContext(
		ctx,
		query,
		p.OrderID,
		p.Amount,
		p.PaymentType,
		p.ID,
	).Scan(
		&payment.ID,
		&payment.OrderID,
		&payment.Amount,
		&payment.PaymentType,
		&payment.PaymentDate,
	)

	if err != nil {
		return nil, fmt.Errorf("error updating payment: %w", err)
	}

	return payment, nil
}

func (r *paymentRepository) Delete(ctx context.Context, paymentID int64) error {
	query := `
		DELETE FROM payments
		WHERE payment_id = $1
	`

	_, err := r.db.ExecContext(ctx, query, paymentID)
	if err != nil {
		return fmt.Errorf("error deleting payment: %w", err)
	}

	return nil
}

func (r *paymentRepository) List(ctx context.Context, limit, offset int) ([]domain.Payment, error) {
	query := `
		SELECT payment_id, order_id, amount, payment_type, created_at
		FROM payments
		LIMIT $1 OFFSET $2
	`

	rows, err := r.db.QueryContext(ctx, query, limit, offset)
	if err != nil {
		return nil, fmt.Errorf("error listing payments: %w", err)
	}
	defer rows.Close()

	payments := []domain.Payment{}
	for rows.Next() {
		var p domain.Payment
		err := rows.Scan(
			&p.ID,
			&p.OrderID,
			&p.Amount,
			&p.PaymentType,
			&p.PaymentDate,
		)
		if err != nil {
			return nil, fmt.Errorf("error listing payments: %w", err)
		}
		payments = append(payments, p)
	}

	if err = rows.Err(); err != nil {
		return nil, fmt.Errorf("error listing payments: %w", err)
	}

	return payments, nil
}
