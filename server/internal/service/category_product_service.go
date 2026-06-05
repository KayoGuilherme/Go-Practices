package service

import (
	"context"
	"errors"
	"time"

	"server/internal/domain"

	"gorm.io/gorm"
)

type categoryProductService struct {
	linkRepo     domain.CategoryProductRepository
	productRepo  domain.ProductRepository
	categoryRepo domain.CategoryRepository
	timeout      time.Duration
}

func NewCategoryProductService(
	linkRepo domain.CategoryProductRepository,
	productRepo domain.ProductRepository,
	categoryRepo domain.CategoryRepository,
) domain.CategoryProductService {
	return &categoryProductService{
		linkRepo:     linkRepo,
		productRepo:  productRepo,
		categoryRepo: categoryRepo,
		timeout:      2 * time.Second,
	}
}

func (s *categoryProductService) AttachCategory(ctx context.Context, productID, categoryID int) error {
	ctx, cancel := context.WithTimeout(ctx, s.timeout)
	defer cancel()

	if err := s.ensureProduct(ctx, productID); err != nil {
		return err
	}
	if err := s.ensureCategory(ctx, categoryID); err != nil {
		return err
	}

	pid := int64(productID)
	cid := int64(categoryID)

	exists, err := s.linkRepo.Exists(ctx, pid, cid)
	if err != nil {
		return err
	}
	if exists {
		return domain.ErrLinkAlreadyExists
	}

	if err := s.linkRepo.Link(ctx, pid, cid); err != nil {
		if errors.Is(err, gorm.ErrDuplicatedKey) {
			return domain.ErrLinkAlreadyExists
		}
		return err
	}
	return nil
}

func (s *categoryProductService) DetachCategory(ctx context.Context, productID, categoryID int) error {
	ctx, cancel := context.WithTimeout(ctx, s.timeout)
	defer cancel()

	if err := s.ensureProduct(ctx, productID); err != nil {
		return err
	}
	if err := s.ensureCategory(ctx, categoryID); err != nil {
		return err
	}

	return s.linkRepo.Unlink(ctx, int64(productID), int64(categoryID))
}

func (s *categoryProductService) GetCategoriesByProduct(ctx context.Context, productID int) ([]domain.Category, error) {
	ctx, cancel := context.WithTimeout(ctx, s.timeout)
	defer cancel()

	if err := s.ensureProduct(ctx, productID); err != nil {
		return nil, err
	}

	return s.linkRepo.ListCategoriesByProduct(ctx, int64(productID))
}

func (s *categoryProductService) GetProductsByCategory(ctx context.Context, categoryID int) ([]domain.Product, error) {
	ctx, cancel := context.WithTimeout(ctx, s.timeout)
	defer cancel()

	if err := s.ensureCategory(ctx, categoryID); err != nil {
		return nil, err
	}

	return s.linkRepo.ListProductsByCategory(ctx, int64(categoryID))
}

func (s *categoryProductService) ensureProduct(ctx context.Context, id int) error {
	product, err := s.productRepo.FindProductById(ctx, id)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return domain.ErrProductNotFound
		}
		return err
	}
	if product == nil || product.ID == 0 {
		return domain.ErrProductNotFound
	}
	return nil
}

func (s *categoryProductService) ensureCategory(ctx context.Context, id int) error {
	category, err := s.categoryRepo.FindByCategoryId(ctx, id)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return domain.ErrCategoryNotFound
		}
		return err
	}
	if category == nil || category.ID == 0 {
		return domain.ErrCategoryNotFound
	}
	return nil
}
