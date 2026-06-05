package service

import (
	"context"
	"errors"
	"fmt"
	"server/internal/domain"
	productexc "server/internal/domain/exceptions/product"
	"server/internal/util/pagination"
	"time"
)

type productService struct {
	repo    domain.ProductRepository
	timeout time.Duration
}

func NewProductService(repo domain.ProductRepository) domain.ProductService {
	return &productService{
		repo:    repo,
		timeout: 2 * time.Second,
	}
}

func (s *productService) CreateProduct(ctx context.Context, product *domain.Product) (*domain.Product, error) {
	ctx, cancel := context.WithTimeout(ctx, s.timeout)
	defer cancel()

	if product.Price < 5 {
		return nil, productexc.NewProductNotBeLessThanFiveError()
	}

	return s.repo.CreateProduct(ctx, product)
}

func (s *productService) GetAllProducts(ctx context.Context, params pagination.Params) (*pagination.Response[domain.Product], error) {
	ctx, cancel := context.WithTimeout(ctx, s.timeout)
	defer cancel()

	products, total, err := s.repo.FindAllProducts(ctx, params)

	if err != nil {
		return nil, err
	}

	return &pagination.Response[domain.Product]{
		Data:    products,
		Total:   total,
		Limit:   params.Limit,
		Offset:  params.Offset,
		HasMore: int64(params.Offset+params.Limit) < total,
	}, nil
}

func (s *productService) GetProductById(ctx context.Context, id int) (*domain.Product, error) {
	ctx, cancel := context.WithTimeout(ctx, s.timeout)
	defer cancel()

	product, err := s.repo.FindProductById(ctx, id)

	if err != nil {
		if errors.Is(err, domain.ErrProductNotFound) {
			return nil, productexc.NewProductNotFoundError()
		}
		return nil, fmt.Errorf("get product by id: %w", err)
	}

	return product, nil
}

func (s *productService) UpdateByProduct(ctx context.Context, id int, product *domain.Product) (*domain.Product, error) {
	ctx, cancel := context.WithTimeout(ctx, s.timeout)
	defer cancel()

	updatedProduct, err := s.repo.UpdateByProductId(ctx, id, product)

	if err != nil {
		if errors.Is(err, domain.ErrProductNotFound) {
			return nil, productexc.NewProductNotFoundError()
		}
		return nil, fmt.Errorf("update product by id: %w", err)
	}

	return updatedProduct, nil
}

func (s *productService) DeleteByProduct(ctx context.Context, id int) error {
	ctx, cancel := context.WithTimeout(ctx, s.timeout)
	defer cancel()

	if err := s.repo.DeleteByProductId(ctx, id); err != nil {
		if errors.Is(err, domain.ErrProductNotFound) {
			return productexc.NewProductNotFoundError()
		}
		return fmt.Errorf("delete product by id: %w", err)
	}

	return nil
}
