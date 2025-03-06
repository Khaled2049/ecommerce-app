package interfaces

import (
	"context"

	"github.com/Khaled2049/ecommerce-app/internal/domain"
)

type ProductRepository interface {
	Create(ctx context.Context, product *domain.ProductCreate) (*domain.Product, error)
	GetByID(ctx context.Context, id uint64) (*domain.Product, error)
	Update(ctx context.Context, product *domain.ProductUpdate) (*domain.Product, error)
	Delete(ctx context.Context, id uint64) error
	List(ctx context.Context, limit, offset int) ([]*domain.Product, error)
}
