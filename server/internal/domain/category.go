package domain

import (
	"context"
	"errors"
	"time"
)


type Category struct {
	ID        int64     `gorm:"primaryKey"`
	Name 	  string 	`gorm:"column:name"`
	CreatedAt time.Time `gorm:"column:created_at"` 
}

var (
	ErrCategoryNotFound = errors.New("Category not found")
)


type CategoryRepository interface {
	CreateCategory(ctx context.Context, category *Category) (*Category, error)
	FindAllCategory(ctx context.Context) ([]Category, error)
	FindByCategoryId(ctx context.Context, id int) (*Category, error)
}


type CategoryService interface {
	CreateCategory(ctx context.Context, category *Category) (*Category, error)
	GetAllCategory(ctx context.Context) ([]Category, error)
	GetCategory(ctx context.Context, id int) (*Category, error)
}