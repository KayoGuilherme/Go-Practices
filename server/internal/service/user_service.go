package service

import (
	"context"
	"errors"
	"fmt"
	"server/internal/domain"
	"server/internal/domain/exceptions/users"
	"server/internal/util/pagination"
	"server/util"
	"time"

	"github.com/golang-jwt/jwt/v4"
)

type userService struct {
	repo    domain.Repository
	timeout time.Duration
}

type JwtClaims struct {
	ID   string `json:"id"`
	Name string `json:"name"`
	jwt.RegisteredClaims
}

func NewService(repo domain.Repository) domain.UserService {
	return &userService{
		repo:    repo,
		timeout: 2 * time.Second,
	}
}

func (s *userService) CreateUser(ctx context.Context, name, email, password, phone, cpf string) (*domain.User, error) {
	ctx, cancel := context.WithTimeout(ctx, s.timeout)
	defer cancel()

	user, err := s.repo.GetAndValidateUser(ctx, email, cpf, phone)

	if err != nil {
		return nil, fmt.Errorf("validate user uniqueness: %w", err)
	}

	if user != nil {
		if user.Email == email {
			return nil, users.NewUserEmailAlreadyExistsError()
		}
		if user.Cpf == cpf {
			return nil, users.NewUserCpfAlreadyExistsError()
		}
		if user.Phone == phone {
			return nil, users.NewUserPhoneAlreadyExistsError()
		}
	}

	hashedPassword, err := util.HashPassword(password)
	if err != nil {
		return nil, fmt.Errorf("hash password: %w", err)
	}

	user = &domain.User{
		Name:     name,
		Email:    email,
		Phone:    phone,
		Cpf:      cpf,
		Password: hashedPassword,
	}

	return s.repo.CreateUser(ctx, user)
}

func (s *userService) GetAllUsers(ctx context.Context, params pagination.Params) (*pagination.Response[domain.User], error) {
	ctx, cancel := context.WithTimeout(ctx, s.timeout)
	defer cancel()

	users, total, err := s.repo.GetAllUsers(ctx, params)

	if err != nil {
		return nil, fmt.Errorf("get all users: %w", err)
	}

	return &pagination.Response[domain.User]{
		Data:    users,
		Total:   total,
		Limit:   params.Limit,
		Offset:  params.Offset,
		HasMore: int64(params.Offset+params.Limit) < total,
	}, nil
}

func (s *userService) GetUserById(ctx context.Context, id int) (*domain.User, error) {
	ctx, cancel := context.WithTimeout(ctx, s.timeout)
	defer cancel()

	userExist, err := s.repo.GetUserById(ctx, id)

	if err != nil {
		if errors.Is(err, users.ErrUserNotFound) {
			return nil, users.NewUserNotFoundError()
		}
		return nil, fmt.Errorf("get user by id: %w", err)
	}

	return userExist, nil
}

func (s *userService) GetUserByCpf(ctx context.Context, cpf string) (*domain.User, error) {
	ctx, cancel := context.WithTimeout(ctx, s.timeout)
	defer cancel()

	userExist, err := s.repo.GetUserByCpf(ctx, cpf)

	if err != nil {
		if errors.Is(err, users.ErrUserNotFound) {
			return nil, users.NewUserNotFoundError()
		}
		return nil, fmt.Errorf("get user by cpf: %w", err)
	}

	return userExist, nil
}

func (s *userService) UpdateUserById(ctx context.Context, id int, name, email, phone, cpf string) (*domain.User, error) {
	ctx, cancel := context.WithTimeout(ctx, s.timeout)
	defer cancel()

	userExist, err := s.repo.GetUserById(ctx, id)

	if err != nil {
		if errors.Is(err, users.ErrUserNotFound) {
			return nil, users.NewUserNotFoundError()
		}
		return nil, fmt.Errorf("get user by id: %w", err)
	}

	updatedUser := domain.User{
		Name:  name,
		Email: email,
		Phone: phone,
		Cpf:   cpf,
	}

	updated, err := s.repo.UpdateUserById(ctx, int(userExist.ID), &updatedUser)
	if err != nil {
		if errors.Is(err, users.ErrUserNotFound) {
			return nil, users.NewUserNotFoundError()
		}
		return nil, fmt.Errorf("update user by id: %w", err)
	}

	return updated, nil
}

func (s *userService) DeleteUserById(ctx context.Context, id int) error {
	ctx, cancel := context.WithTimeout(ctx, s.timeout)
	defer cancel()

	if _, err := s.repo.GetUserById(ctx, id); err != nil {
		if errors.Is(err, users.ErrUserNotFound) {
			return users.NewUserNotFoundError()
		}
		return fmt.Errorf("get user by id: %w", err)
	}

	if err := s.repo.DeleteUserById(ctx, id); err != nil {
		if errors.Is(err, users.ErrUserNotFound) {
			return users.NewUserNotFoundError()
		}
		return fmt.Errorf("delete user by id: %w", err)
	}

	return nil
}
