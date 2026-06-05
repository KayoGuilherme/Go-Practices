package domain

import (
	"context"
	"errors"
)

var (
	ErrLinkAlreadyExists = errors.New("product already linked to category")
	ErrLinkNotFound      = errors.New("link not found")
)

type CategoryProduct struct {
	ProductID  int64 `gorm:"column:product_id;primaryKey"`
	CategoryID int64 `gorm:"column:category_id;primaryKey"`
}

func (CategoryProduct) TableName() string {
	return "category_product"
}

type CategoryProductRepository interface {
	Link(ctx context.Context, productID, categoryID int64) error
	Unlink(ctx context.Context, productID, categoryID int64) error
	Exists(ctx context.Context, productID, categoryID int64) (bool, error)
	ListCategoriesByProduct(ctx context.Context, productID int64) ([]Category, error)
	ListProductsByCategory(ctx context.Context, categoryID int64) ([]Product, error)
}

type CategoryProductService interface {
	AttachCategory(ctx context.Context, productID, categoryID int) error
	DetachCategory(ctx context.Context, productID, categoryID int) error
	GetCategoriesByProduct(ctx context.Context, productID int) ([]Category, error)
	GetProductsByCategory(ctx context.Context, categoryID int) ([]Product, error)
}
