package repository

import (
	"context"

	"server/internal/domain"
	"server/internal/util/pagination"

	"gorm.io/gorm"
)

type productRepository struct {
	db *gorm.DB
}

func NewProductRepository(db *gorm.DB) *productRepository {
	return &productRepository{db: db}
}

func (repo *productRepository) CreateProduct(ctx context.Context, product *domain.Product) (*domain.Product, error) {

	if result := repo.db.WithContext(ctx).Create(product); result.Error != nil {
		return nil, result.Error
	}

	return product, nil
}

func (repo *productRepository) FindAllProducts(ctx context.Context, params pagination.Params) ([]domain.Product, int64, error) {
	var products []domain.Product
	var total int64

	query := repo.db.WithContext(ctx).Model(&domain.Product{})

	if err := query.Count(&total).Error; err != nil {
		return nil, 0, err
	}

	if err := query.
		Limit(params.Limit).
		Offset(params.Offset).
		Order("created_at DESC").
		Find(&products).Error; err != nil {
		return nil, 0, err
	}

	return products, total, nil
}

func (repo *productRepository) FindProductById(ctx context.Context, id int) (*domain.Product, error) {
	var product *domain.Product

	if result := repo.db.WithContext(ctx).Where("id = ?", id).First(&product); result.Error != nil {
		return nil, result.Error
	}

	return product, nil
}

func (repo *productRepository) UpdateByProductId(ctx context.Context, id int, product *domain.Product) (*domain.Product, error) {
	productId, err := repo.FindProductById(ctx, id)

	if err != nil {
		return nil, err
	}

	result := repo.db.WithContext(ctx).Where("id = ?", productId.ID).Updates(&product)

	if result.Error != nil {
		return nil, result.Error
	}

	return product, nil
}

func (repo *productRepository) DeleteByProductId(ctx context.Context, id int) error {
	var product *domain.Product

	productId, err := repo.FindProductById(ctx, id)

	if err != nil {
		return nil
	}

	result := repo.db.WithContext(ctx).Where("id = ?", productId.ID).Delete(&product)

	if result.Error != nil {
		return nil
	}

	return nil
}
