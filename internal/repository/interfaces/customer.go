package interfaces

import (
	"context"

	"github.com/Khaled2049/ecommerce-app/internal/domain"
)

type CustomerRepository interface {
	Create(ctx context.Context, customer *domain.CustomerCreate) (*domain.Customer, error)
	GetByID(ctx context.Context, id int64) (*domain.Customer, error)
	Update(ctx context.Context, customer *domain.CustomerUpdate) (*domain.Customer, error)
	Delete(ctx context.Context, id int64) error
	List(ctx context.Context, limit, offset int) ([]domain.Customer, error)
	GetByEmail(ctx context.Context, email string) (*domain.Customer, error)
}
