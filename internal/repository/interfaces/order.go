package interfaces

import (
	"context"

	"github.com/Khaled2049/ecommerce-app/internal/domain"
)


type OrderRepository interface {
	Create(ctx context.Context, order *domain.OrderCreate) (*domain.Order, error)
	GetByID(ctx context.Context, id uint64) (*domain.Order, error)
	GetByCustomerID(ctx context.Context, customerID uint64) ([]*domain.Order, error)
	Update(ctx context.Context, order *domain.Order) (*domain.Order, error)
	Delete(ctx context.Context, id uint64) error
	List(ctx context.Context, limit, offset int) ([]*domain.Order, error)
}
