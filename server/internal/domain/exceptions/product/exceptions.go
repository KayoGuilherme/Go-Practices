package product

import (
	"errors"
	"server/internal/domain/exceptions"
)

var (
	ErrPriceNotBeLessThanFive = errors.New("Price not be less than R$5,00")
	ErrProductNotFound = errors.New("Product not found")
)

func NewProductNotFoundError() *exceptions.AppError {
	return &exceptions.AppError{
		StatusCode: 404,
		Message:    "product not found",
        Err:        ErrProductNotFound,
	}
}

func NewProductNotBeLessThanFiveError() *exceptions.AppError {
	return &exceptions.AppError{
		StatusCode: 400,
		Message: "Price not be less than R$5,00",
		Err: ErrPriceNotBeLessThanFive,
	}
}