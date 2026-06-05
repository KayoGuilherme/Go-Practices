package category

import (
	"errors"
	"server/internal/domain/exceptions"
)

var (
	ErrCategoryNotFound = errors.New("Category not found")
)

func NewCategoryNotFoundError() *exceptions.AppError {
	return &exceptions.AppError{
		StatusCode: 404,
		Message:    "Category not found",
		Err:        ErrCategoryNotFound,
	}
}
