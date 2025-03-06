package interfaces

import (
	"context"

	"github.com/Khaled2049/ecommerce-app/internal/domain"
)

type PaymentRepository interface {
	Create(ctx context.Context, payment *domain.PaymentCreate) (*domain.Payment, error)
	GetByID(ctx context.Context, id int64) (*domain.Payment, error)
	Update(ctx context.Context, payment *domain.PaymentUpdate) (*domain.Payment, error)
	Delete(ctx context.Context, id int64) error
	List(ctx context.Context, limit, offset int) ([]domain.Payment, error)
}
