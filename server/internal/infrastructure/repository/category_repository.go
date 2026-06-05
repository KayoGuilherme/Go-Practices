package repository

import (
	"context"
	"server/internal/domain"

	"gorm.io/gorm"
)

type categoryRepository struct {
	db *gorm.DB
}

func NewCategoryRepository(db *gorm.DB) *categoryRepository {
	return &categoryRepository{
		db: db,
	}
}

func (repo *categoryRepository) CreateCategory(ctx context.Context, category *domain.Category) (*domain.Category, error) {
	if result := repo.db.WithContext(ctx).Create(&category); result.Error != nil {
		return nil, result.Error
	}

	return category, nil
}

func (repo *categoryRepository) FindAllCategory(ctx context.Context) ([]domain.Category, error) {
	var category []domain.Category

	if result := repo.db.WithContext(ctx).Find(&category); result.Error != nil {
		return nil, result.Error
	}

	return category, nil
}

func (repo *categoryRepository) FindByCategoryId(ctx context.Context, id int) (*domain.Category, error) {
	var category *domain.Category

	if result := repo.db.WithContext(ctx).Where("id = ?", id).Find(&category); result.Error != nil {
		return nil, result.Error
	}

	return category, nil
}
