package service

import (
	"context"

	"github.com/Khaled2049/ecommerce-app/internal/domain"
	"github.com/Khaled2049/ecommerce-app/internal/repository/interfaces"
	"github.com/go-playground/validator/v10"
)

type OrderService struct {
	repo     interfaces.OrderRepository
	validate *validator.Validate
}

func NewOrderService(repo interfaces.OrderRepository) *OrderService {
	return &OrderService{
		repo:     repo,
		validate: validator.New(),
	}
}

func (s *OrderService) CreateOrder(ctx context.Context, order *domain.OrderCreate) (*domain.Order, error) {
	if err := s.validate.Struct(order); err != nil {
		return nil, err
	}
	return s.repo.Create(ctx, order)
}

func (s *OrderService) GetOrderByID(ctx context.Context, id uint64) (*domain.Order, error) {
	return s.repo.GetByID(ctx, id)
}

func (s *OrderService) UpdateOrder(ctx context.Context, order *domain.Order) (*domain.Order, error) {
	return s.repo.Update(ctx, order)
}

func (s *OrderService) DeleteOrder(ctx context.Context, id uint64) error {
	return s.repo.Delete(ctx, id)
}

func (s *OrderService) GetOrders(ctx context.Context, limit, offset int) ([]*domain.Order, error) {
	return s.repo.List(ctx, limit, offset)
}


