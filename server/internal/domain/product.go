package domain

import (
	"context"
	"errors"
	"server/internal/util/pagination"
	"time"
)


var (
	ErrPriceNotBeLessThanFive = errors.New("Price not be less than R$5,00")
	ErrProductNotFound = errors.New("Product not found")
)


type Product struct {
	ID          int64   `gorm:"primaryKey"`
	Name        string  `gorm:"column:name"`
	Price       float64 `gorm:"column:price;type:numeric(10,2)"`
	Description string  `gorm:"column:description"`
	Stock       string  `gorm:"column:stock"`
	Is_on_sale  bool    `gorm:"column:is_on_sale"`
	Weight      float64 `gorm:"column:weight;type:numeric(10,2)"`
	Width       float64 `gorm:"column:width;type:numeric(10,2)"`
	Diameter    int32   `gorm:"column:diameter"`
	Length      int32   `gorm:"column:length"`
	CreatedAt   time.Time
	UpdatedAt   time.Time
}

type ProductRepository interface {
	CreateProduct(ctx context.Context, product *Product) (*Product, error)
	FindAllProducts(ctx context.Context, params pagination.Params) ([]Product, int64, error)
	FindProductById(ctx context.Context, id int) (*Product, error)
	UpdateByProductId(ctx context.Context, id int, product *Product) (*Product, error)
	DeleteByProductId(ctx context.Context, id int) error
}

type ProductService interface {
	CreateProduct(ctx context.Context, product *Product) (*Product, error)
	GetAllProducts(ctx context.Context, params pagination.Params) (*pagination.Response[Product], error)
	GetProductById(ctx context.Context, id int) (*Product, error)
	UpdateByProduct(ctx context.Context, id int, product *Product) (*Product, error)
	DeleteByProduct(ctx context.Context, id int) error
}
