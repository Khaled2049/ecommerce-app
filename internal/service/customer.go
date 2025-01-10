// internal/service/customer.go

package service

import (
	"context"

	"github.com/Khaled2049/ecommerce-app/internal/domain"
	"github.com/Khaled2049/ecommerce-app/internal/repository/interfaces"
)

type CustomerService struct {
	repo interfaces.CustomerRepository
}

func NewCustomerService(repo interfaces.CustomerRepository) *CustomerService {
	return &CustomerService{
		repo: repo,
	}
}

func (s *CustomerService) CreateCustomer(ctx context.Context, input *domain.CustomerCreate) (*domain.Customer, error) {
	// Add business logic validation here
	if err := validateCustomerInput(input); err != nil {
		return nil, err
	}

	// Check if email already exists
	// Add more business logic as needed

	return s.repo.Create(ctx, input)
}

func (s *CustomerService) GetCustomers(ctx context.Context) ([]domain.Customer, error) {
	// Add business logic as needed

	return s.repo.List(ctx, 10, 0)
}

func validateCustomerInput(input *domain.CustomerCreate) error {
	// Implement validation logic
	return nil
}
