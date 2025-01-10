// internal/service/customer.go

package service

import (
	"context"
	"errors"
	"fmt"

	"github.com/Khaled2049/ecommerce-app/internal/domain"
	"github.com/Khaled2049/ecommerce-app/internal/repository/interfaces"
	"github.com/go-playground/validator/v10"
)

type CustomerService struct {
	repo     interfaces.CustomerRepository
	validate *validator.Validate
}

func NewCustomerService(repo interfaces.CustomerRepository) *CustomerService {
	return &CustomerService{
		repo:     repo,
		validate: validator.New(),
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

func (s *CustomerService) GetCustomerByID(ctx context.Context, id int64) (*domain.Customer, error) {
	// Add business logic as needed

	return s.repo.GetByID(ctx, id)
}

func (s *CustomerService) UpdateCustomer(ctx context.Context, input *domain.CustomerUpdate) (*domain.Customer, error) {
	// Validate input
	if err := s.validate.Struct(input); err != nil {
		return nil, fmt.Errorf("invalid input: %w", err)
	}

	// Check if email is already taken by another customer
	existing, err := s.repo.GetByEmail(ctx, input.Email)
	if err != nil && !errors.Is(err, domain.ErrCustomerNotFound) {
		return nil, err
	}
	if existing != nil && existing.ID != input.ID {
		return nil, domain.ErrEmailAlreadyExists
	}

	// Update the customer
	updated, err := s.repo.Update(ctx, input)
	if err != nil {
		return nil, err
	}

	return updated, nil
}

func (s *CustomerService) DeleteCustomer(ctx context.Context, id int64) error {
	// Optional: Check if customer exists first
	_, err := s.repo.GetByID(ctx, id)
	if err != nil {
		return err
	}

	if err := s.repo.Delete(ctx, id); err != nil {
		return fmt.Errorf("failed to delete customer: %w", err)
	}

	return nil
}
func validateCustomerInput(input *domain.CustomerCreate) error {
	// Implement validation logic
	return nil
}
