package service

import (
	"context"

	"github.com/Khaled2049/ecommerce-app/internal/domain"
	"github.com/Khaled2049/ecommerce-app/internal/repository/interfaces"
)

type PaymentService struct {
	repo interfaces.PaymentRepository
}

func NewPaymentService(repo interfaces.PaymentRepository) *PaymentService {
	return &PaymentService{
		repo: repo,
	}
}

func (s *PaymentService) CreatePayment(ctx context.Context, input *domain.PaymentCreate) (*domain.Payment, error) {
	return s.repo.Create(ctx, input)
}

func (s *PaymentService) GetPayments(ctx context.Context) ([]domain.Payment, error) {
	return s.repo.List(ctx, 10, 0)
}

func (s *PaymentService) GetPaymentByID(ctx context.Context, id int64) (*domain.Payment, error) {
	return s.repo.GetByID(ctx, id)
}

func (s *PaymentService) UpdatePayment(ctx context.Context, input *domain.PaymentUpdate) (*domain.Payment, error) {
	return s.repo.Update(ctx, input)
}

func (s *PaymentService) DeletePayment(ctx context.Context, id int64) error {
	return s.repo.Delete(ctx, id)
}
