package service

import (
	"context"
	"errors"
	"fmt"
	"server/internal/domain"
	categoryexc "server/internal/domain/exceptions/category"
	"time"
)

type categoryService struct {
	repo    domain.CategoryRepository
	timeout time.Duration
}

func NewCategoryService(repo domain.CategoryRepository) domain.CategoryService {
	return &categoryService{
		repo:    repo,
		timeout: 2 * time.Second,
	}
}

func (c *categoryService) CreateCategory(ctx context.Context, category *domain.Category) (*domain.Category, error) {
	ctx, cancel := context.WithTimeout(ctx, c.timeout)
	defer cancel()

	if _, err := c.repo.CreateCategory(ctx, category); err != nil {
		return nil, err
	}

	return category, nil
}

func (c *categoryService) GetAllCategory(ctx context.Context) ([]domain.Category, error) {
	ctx, cancel := context.WithTimeout(ctx, c.timeout)
	defer cancel()

	category, err := c.repo.FindAllCategory(ctx)

	if err != nil {
		return nil, err
	}
	return category, nil
}

func (c *categoryService) GetCategory(ctx context.Context, id int) (*domain.Category, error) {
	ctx, cancel := context.WithTimeout(ctx, c.timeout)
	defer cancel()

	category, err := c.repo.FindByCategoryId(ctx, id)

	if err != nil {
		if errors.Is(err, domain.ErrCategoryNotFound) {
			return nil, categoryexc.NewCategoryNotFoundError()
		}
		return nil, fmt.Errorf("get category by id: %w", err)
	}

	return category, nil
}
