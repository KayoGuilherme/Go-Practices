package repository

import (
	"context"

	"server/internal/domain"

	"gorm.io/gorm"
)

type categoryProductRepository struct {
	db *gorm.DB
}

func NewCategoryProductRepository(db *gorm.DB) domain.CategoryProductRepository {
	return &categoryProductRepository{db: db}
}

func (repo *categoryProductRepository) Link(ctx context.Context, productID, categoryID int64) error {
	link := domain.CategoryProduct{
		ProductID:  productID,
		CategoryID: categoryID,
	}
	return repo.db.WithContext(ctx).Create(&link).Error
}

func (repo *categoryProductRepository) Unlink(ctx context.Context, productID, categoryID int64) error {
	result := repo.db.WithContext(ctx).
		Where("product_id = ? AND category_id = ?", productID, categoryID).
		Delete(&domain.CategoryProduct{})

	if result.Error != nil {
		return result.Error
	}

	if result.RowsAffected == 0 {
		return domain.ErrLinkNotFound
	}
	return nil
}

func (repo *categoryProductRepository) Exists(ctx context.Context, productID, categoryID int64) (bool, error) {
	var count int64
	err := repo.db.WithContext(ctx).
		Model(&domain.CategoryProduct{}).
		Where("product_id = ? AND category_id = ?", productID, categoryID).
		Count(&count).Error
	if err != nil {
		return false, err
	}
	return count > 0, nil
}

func (repo *categoryProductRepository) ListCategoriesByProduct(ctx context.Context, productID int64) ([]domain.Category, error) {
	var categories []domain.Category
	err := repo.db.WithContext(ctx).
		Model(&domain.Category{}).
		Joins("INNER JOIN category_product ON category_product.category_id = categories.id").
		Where("category_product.product_id = ?", productID).
		Find(&categories).Error
	if err != nil {
		return nil, err
	}
	return categories, nil
}

func (repo *categoryProductRepository) ListProductsByCategory(ctx context.Context, categoryID int64) ([]domain.Product, error) {
	var products []domain.Product
	err := repo.db.WithContext(ctx).
		Model(&domain.Product{}).
		Joins("INNER JOIN category_product ON category_product.product_id = products.id").
		Where("category_product.category_id = ?", categoryID).
		Find(&products).Error
	if err != nil {
		return nil, err
	}
	return products, nil
}
