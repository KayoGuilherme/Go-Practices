package domain

import (
	"context"
	"server/internal/util/pagination"
	"time"
)

type User struct {
    ID        int64     `gorm:"primaryKey"`
    Name      string    `gorm:"column:name"`
    Email     string    `gorm:"column:email"`
    Phone     string    `gorm:"column:phone"`
    Cpf       string    `gorm:"column:cpf"`
    Password  string    `gorm:"column:password"`
    CreatedAt time.Time
    UpdatedAt time.Time
}

type Repository interface {
	CreateUser(ctx context.Context, user *User) (*User, error)
	GetAllUsers(ctx context.Context, params pagination.Params) ([]User, int64, error)
	GetUserByEmail(ctx context.Context, email string) (*User, error)
	GetUserById(ctx context.Context, id int) (*User, error)
	GetUserByCpf(ctx context.Context, cpf string) (*User, error)
	GetUserByPhone(ctx context.Context, phone string) (*User, error)
	UpdateUserById(ctx context.Context, id int, user *User) (*User, error)
	DeleteUserById(ctx context.Context, id int) error 
	GetAndValidateUser(ctx context.Context, email, cpf, phone string) (*User, error)
}

type UserService interface {
	CreateUser(ctx context.Context, name, email, password, phone, cpf string) (*User, error)
	GetAllUsers(ctx context.Context, params pagination.Params) (*pagination.Response[User], error)
	GetUserById(ctx context.Context, id int) (*User, error)
	GetUserByCpf(ctx context.Context, cpf string) (*User, error)
	UpdateUserById(ctx context.Context, id int, name, email,  phone, cpf string) (*User, error)
	DeleteUserById(ctx context.Context, id int) error 
}

type AuthService interface {
	Login(ctx context.Context, email, password string) (accessToken string, user *User, err error)
}
