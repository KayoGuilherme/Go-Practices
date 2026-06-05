package repository

import (
	"context"
	"errors"
	"fmt"
	"server/internal/domain"
	userexc "server/internal/domain/exceptions/users"
	"server/internal/util/pagination"

	"gorm.io/gorm"
)

type repository struct {
    db *gorm.DB
}

func NewRepository(db *gorm.DB) (domain.Repository) {
    return &repository{db: db}
}

func (repo *repository) CreateUser(ctx context.Context, user *domain.User) (*domain.User, error) {
    result := repo.db.WithContext(ctx).Create(user)
    if result.Error != nil {
        return nil, fmt.Errorf("create user: %w", result.Error)
    }
    return user, nil
}

func (repo *repository) GetAllUsers(ctx context.Context, params pagination.Params) ([]domain.User, int64, error) {    
    var users []domain.User
    var total int64

    
    query := repo.db.WithContext(ctx).Model(&domain.User{})

    if err := query.Count(&total).Error; err != nil {
        return nil, 0, err
    }

    if err := query.
        Limit(params.Limit).
        Offset(params.Offset).
        Order("created_at DESC"). 
        Find(&users).Error; err != nil {
        return nil, 0, err
    }

    return users, total, nil
}

func (repo *repository) GetAndValidateUser(ctx context.Context, email, cpf, phone string) (*domain.User, error) {
    var user domain.User
    result := repo.db.WithContext(ctx).Where("email = ?", email).Or("cpf = ?", cpf).Or("phone = ?", phone).First(&user)
    if result.Error != nil {
        if errors.Is(result.Error, gorm.ErrRecordNotFound) {
            return nil, nil
        }
        return nil, fmt.Errorf("get and validate user: %w", result.Error)
    }
    return &user, nil
}

func (repo *repository) GetUserByEmail(ctx context.Context, email string) (*domain.User, error) {
    var user domain.User
    result := repo.db.WithContext(ctx).Where("email = ?", email).First(&user)
    if result.Error != nil {
        if errors.Is(result.Error, gorm.ErrRecordNotFound) {
            return nil, fmt.Errorf("get user by email: %w", userexc.ErrUserNotFound)
        }
        return nil, fmt.Errorf("get user by email: %w", result.Error)
        }
    return &user, nil
}

func (repo *repository) GetUserById(ctx context.Context, id int) (*domain.User, error) {
    var user domain.User
    result := repo.db.WithContext(ctx).Where("id = ?", id).First(&user)
    if result.Error != nil {
        if errors.Is(result.Error, gorm.ErrRecordNotFound) {
            return nil, fmt.Errorf("get user by id %d: %w", id, userexc.ErrUserNotFound)
        }
        return nil, fmt.Errorf("get user by id %d: %w", id, result.Error)
    }
    return &user, nil
}

func (repo *repository) GetUserByCpf(ctx context.Context, cpf string) (*domain.User, error) {
    var user domain.User
    result := repo.db.WithContext(ctx).Where("cpf = ?", cpf).First(&user)
    if result.Error != nil {
        if errors.Is(result.Error, gorm.ErrRecordNotFound) {
            return nil, fmt.Errorf("get user by cpf: %w", userexc.ErrUserNotFound)
        }
        return nil, fmt.Errorf("get user by cpf: %w", result.Error)
    }
    return &user, nil
}

func (repo *repository) GetUserByPhone(ctx context.Context, phone string) (*domain.User, error) {
    var user domain.User
    result := repo.db.WithContext(ctx).Where("phone = ?", phone).First(&user)
    if result.Error != nil {
        if errors.Is(result.Error, gorm.ErrRecordNotFound) {
            return nil, fmt.Errorf("get user by phone: %w", userexc.ErrUserNotFound)
        }
        return nil, fmt.Errorf("get user by phone: %w", result.Error)
    }
    return &user, nil
}

func (repo *repository) UpdateUserById(ctx context.Context, id int, user *domain.User) (*domain.User, error) {
    result := repo.db.WithContext(ctx).Where("id = ?", id).Updates(user)
    if result.Error != nil {
        return nil, fmt.Errorf("update user by id %d: %w", id, result.Error)
    }

    if result.RowsAffected == 0 {
        return nil, fmt.Errorf("update user by id %d: %w", id, userexc.ErrUserNotFound)
    }

    var updatedUser domain.User
    if err := repo.db.WithContext(ctx).Where("id = ?", id).First(&updatedUser).Error; err != nil {
        return nil, fmt.Errorf("get updated user by id %d: %w", id, err)
    }
    return &updatedUser, nil
}

func (repo *repository) DeleteUserById(ctx context.Context, id int) error {
    result := repo.db.WithContext(ctx).Where("id = ?", id).Delete(&domain.User{})
    if result.Error != nil {
        return fmt.Errorf("delete user by id %d: %w", id, result.Error)
    }
    if result.RowsAffected == 0 {
        return fmt.Errorf("delete user by id %d: %w", id, userexc.ErrUserNotFound)
    }
    return nil
}