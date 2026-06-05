package users

import (
	"errors"
	"server/internal/domain/exceptions"
)

var (
	ErrInvalidCredentials = errors.New("invalid email or password")
	ErrEmailAlreadyExists = errors.New("email already exists")
	ErrCpfAlreadyExists = errors.New("cpf already exists")
	ErrPhoneAlreadyExists = errors.New("phone already exists")
	ErrUserNotFound = errors.New("User not found")
)


func NewUserInvalidCredentialsError() *exceptions.AppError {
	return &exceptions.AppError{
		StatusCode: 401,
		Message:    "invalid email or password",
		Err:        ErrInvalidCredentials,
	}
}

func NewUserEmailAlreadyExistsError() *exceptions.AppError {
	return &exceptions.AppError{
		StatusCode: 409,
		Message:    "email already exists",
		Err:        ErrEmailAlreadyExists,
	}
}

func NewUserCpfAlreadyExistsError() *exceptions.AppError {
	return &exceptions.AppError{
		StatusCode: 409,
		Message:    "cpf already exists",
		Err:        ErrCpfAlreadyExists,
	}
}

func NewUserPhoneAlreadyExistsError() *exceptions.AppError {
	return &exceptions.AppError{
		StatusCode: 409,
		Message:    "phone already exists",
		Err:        ErrPhoneAlreadyExists,
	}
}

func NewUserNotFoundError() *exceptions.AppError {
	return &exceptions.AppError{
		StatusCode: 404,
		Message:    "user not found",
		Err:        ErrUserNotFound,
	}
}
