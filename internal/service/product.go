package service

import (
	"context"

	"github.com/Khaled2049/ecommerce-app/internal/domain"
	"github.com/Khaled2049/ecommerce-app/internal/repository/interfaces"
)

type ProductService struct {
	repo interfaces.ProductRepository
}

func NewProductService(repo interfaces.ProductRepository) *ProductService {
	return &ProductService{
		repo: repo,
	}
}

func (s *ProductService) CreateProduct(ctx context.Context, product *domain.ProductCreate) (*domain.Product, error) {
	return s.repo.Create(ctx, product)
}

func (s *ProductService) GetProductByID(ctx context.Context, id uint64) (*domain.Product, error) {
	return s.repo.GetByID(ctx, id)
}

func (s *ProductService) UpdateProduct(ctx context.Context, product *domain.ProductUpdate) (*domain.Product, error) {
	return s.repo.Update(ctx, product)
}

func (s *ProductService) DeleteProduct(ctx context.Context, id uint64) error {
	return s.repo.Delete(ctx, id)
}

func (s *ProductService) ListProducts(ctx context.Context, limit, offset int) ([]*domain.Product, error) {
	return s.repo.List(ctx, limit, offset)
}
