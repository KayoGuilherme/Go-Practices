package service_test

import (
	"context"
	"errors"
	"server/internal/domain"
	userexc "server/internal/domain/exceptions/users"
	"server/internal/util/pagination"
	"server/internal/service"
	"testing"
)

type mockRepository struct {
	CreateUserFn         func(ctx context.Context, user *domain.User) (*domain.User, error)
	GetAllUsersFn        func(ctx context.Context, params pagination.Params) ([]domain.User, int64, error)
	GetUserByEmailFn     func(ctx context.Context, email string) (*domain.User, error)
	GetUserByIdFn        func(ctx context.Context, id int) (*domain.User, error)
	GetUserByCpfFn       func(ctx context.Context, cpf string) (*domain.User, error)
	GetUserByPhoneFn     func(ctx context.Context, phone string) (*domain.User, error)
	UpdateUserByIdFn     func(ctx context.Context, id int, user *domain.User) (*domain.User, error)
	DeleteUserByIdFn     func(ctx context.Context, id int) error
	GetAndValidateUserFn func(ctx context.Context, email, cpf, phone string) (*domain.User, error)
}

func (m *mockRepository) CreateUser(ctx context.Context, user *domain.User) (*domain.User, error) {
	return m.CreateUserFn(ctx, user)
}

func (m *mockRepository) GetAndValidateUser(ctx context.Context, email, cpf, phone string) (*domain.User, error) {
	return m.GetAndValidateUserFn(ctx, email, cpf, phone)
}

func (m *mockRepository) GetAllUsers(ctx context.Context, params pagination.Params) ([]domain.User, int64,error) {
	return m.GetAllUsersFn(ctx, params)
}

func (m *mockRepository) GetUserByEmail(ctx context.Context, email string) (*domain.User, error) {
	return m.GetUserByEmailFn(ctx, email)
}

func (m *mockRepository) GetUserById(ctx context.Context, id int) (*domain.User, error) {
	return m.GetUserByIdFn(ctx, id)
}

func (m *mockRepository) GetUserByCpf(ctx context.Context, cpf string) (*domain.User, error) {
	return m.GetUserByCpfFn(ctx, cpf)
}

func (m *mockRepository) GetUserByPhone(ctx context.Context, phone string) (*domain.User, error) {
	return m.GetUserByPhoneFn(ctx, phone)
}

func (m *mockRepository) UpdateUserById(ctx context.Context, id int, user *domain.User) (*domain.User, error) {
	return m.UpdateUserByIdFn(ctx, id, user)
}

func (m *mockRepository) DeleteUserById(ctx context.Context, id int) error {
	return m.DeleteUserByIdFn(ctx, id)
}

func TestCreateUser_Success(t *testing.T) {

	mock := &mockRepository{
		GetAndValidateUserFn: func(ctx context.Context, email, cpf, phone string) (*domain.User, error) {
			return nil, nil // nenhum usuário duplicado
		},

		CreateUserFn: func(ctx context.Context, user *domain.User) (*domain.User, error) {
			return user, nil
		},
	}

	svc := service.NewService(mock)

	// --- ACT: executa o que estamos testando ---
	user, err := svc.CreateUser(
		context.Background(),
		"John Doe",
		"john@email.com",
		"senha123",
		"11999999999",
		"12345678901",
	)

	// --- ASSERT: verifica o resultado ---
	if err != nil {
		t.Fatalf("esperava err nil, got: %v", err)
	}
	if user == nil {
		t.Fatal("esperava um usuário, got nil")
	}
	if user.Name != "John Doe" {
		t.Errorf("esperava name 'John Doe', got: %s", user.Name)
	}
	if user.Email != "john@email.com" {
		t.Errorf("esperava email 'john@email.com', got: %s", user.Email)
	}
	// A senha deve ter sido hasheada — não pode ser igual à original
	if user.Password == "senha123" {
		t.Error("senha não foi hasheada")
	}

}

func TestCreateUser_UserEmailAlreadyExists(t *testing.T) {

	mock := &mockRepository{
		GetAndValidateUserFn: func(ctx context.Context, email, cpf, phone string) (*domain.User, error) {
			return &domain.User{
				Email: "john@email.com", // mesmo email que vamos tentar cadastrar
				Cpf:   "99999999999",
				Phone: "11888888888",
			}, nil // nenhum usuário duplicado
		},
	}

	svc := service.NewService(mock)

	// --- ACT: executa o que estamos testando ---
	user, err := svc.CreateUser(
		context.Background(),
		"John Doe",
		"john@email.com",
		"senha123",
		"11999999999",
		"12345678901",
	)

	if user != nil {
		t.Error("esperava user nil, got um usuário")
	}

	if !errors.Is(err, userexc.ErrEmailAlreadyExists) {
		t.Errorf("esperava ErrEmailAlreadyExists, got: %v", err)
	}

}

func TestGetAllUsers_Success(t *testing.T) {

	mock := &mockRepository{
		GetAllUsersFn:func(ctx context.Context, params pagination.Params) ([]domain.User, int64, error) {
			return []domain.User{
				{Name: "John", Email: "john@email.com"},
				{Name: "Jane", Email: "jane@email.com"},
			}, 0, nil
		},
	}

	svc := service.NewService(mock)

	user, err := svc.GetAllUsers(
		context.Background(),
		pagination.Params{
			Limit: 10,
			Offset: 100,
		},
	)

	if user == nil {
		t.Error("esperava users, got nil")
	}

	if err != nil {
		t.Errorf("esperava users nil, got users")
	}
}

func TestGetAllUsers_Failed(t *testing.T) {

	mock := &mockRepository{
		GetAllUsersFn: func(ctx context.Context, params pagination.Params) ([]domain.User, int64, error)  {
			return nil, 0,errors.New("database errors")
		},
	}

	svc := service.NewService(mock)

	user, err := svc.GetAllUsers(
		context.Background(),
		pagination.Params{
			Limit: 0,
			Offset: 0,
		},
	)

	if user != nil {
		t.Error("esperava nil, got users")
	}

	if err == nil {
		t.Error("esperava erro, got nil")
	}
}

func TestGetUserById_Success(t *testing.T) {

	mock := &mockRepository{
		GetUserByIdFn: func(ctx context.Context, id int) (*domain.User, error) {
			return &domain.User{
				Name: "John", Email: "john@email.com",
			}, nil
		},
	}

	svc := service.NewService(mock)

	user, err := svc.GetUserById(
		context.Background(),
		int(1),
	)

	if user == nil {
		t.Errorf("esperava user, got nil")
	}

	if err != nil {
		t.Errorf("esperava nil, got: %v", err)
	}
}

func TestGetUserById_Failed(t *testing.T) {

	mock := &mockRepository{
		GetUserByIdFn: func(ctx context.Context, id int) (*domain.User, error) {
			return nil, errors.New("database error")
		},
	}

	svc := service.NewService(mock)

	user, err := svc.GetUserById(
		context.Background(),
		int(1),
	)

	if user != nil {
		t.Errorf("esperava nil, got user")
	}

	if errors.Is(err, userexc.ErrUserNotFound) {
		t.Errorf("não esperava ErrUserNotFound para erro de banco, got: %v", err)
	}

	if err == nil {
		t.Error("esperava erro, got nil")
	}
}

func TestGetUserByCpf_Success(t *testing.T) {

	mock := &mockRepository{
		GetUserByCpfFn: func(ctx context.Context, cpf string) (*domain.User, error) {
			return &domain.User{
				Name: "John", Email: "john@email.com",
			}, nil
		},
	}

	svc := service.NewService(mock)

	user, err := svc.GetUserByCpf(
		context.Background(),
		string("1217654323"),
	)

	if user == nil {
		t.Errorf("esperava user, got nil")
	}

	if err != nil {
		t.Errorf("esperava nil, got: %v", err)
	}
}

func TestGetUserByCpf_Failed(t *testing.T) {

	mock := &mockRepository{
		GetUserByCpfFn: func(ctx context.Context, cpf string) (*domain.User, error) {
			return nil, errors.New("database error")
		},
	}

	svc := service.NewService(mock)

	user, err := svc.GetUserByCpf(
		context.Background(),
		string("1217654323"),
	)

	if user != nil {
		t.Errorf("esperava nil, got user")
	}

	if errors.Is(err, userexc.ErrUserNotFound) {
		t.Errorf("não esperava ErrUserNotFound para erro de banco, got: %v", err)
	}

	if err == nil {
		t.Error("esperava erro, got nil")
	}
}

func TestUpdateByUserId_Success(t *testing.T) {
	mock := &mockRepository{
		UpdateUserByIdFn: func(ctx context.Context, id int, user *domain.User) (*domain.User, error) {
			return user, nil
		},
		GetUserByIdFn: func(ctx context.Context, id int) (*domain.User, error) {
			return &domain.User{ID: 1}, nil
		},
	}

	svc := service.NewService(mock)

	user, err := svc.UpdateUserById(
		context.Background(),
		1,
		"John2",
		"john2@gmail.com",
		"859947261212",
		"12167898732",
	)

	if err != nil {
		t.Fatalf("esperava err nil, got: %v", err)
	}

	if user == nil {
		t.Fatal("esperava um usuário, got nil")
	}

	if user.Name != "John2" {
		t.Errorf("esperava name 'John2', got: %s", user.Name)
	}

	if user.Email != "john2@gmail.com" {
		t.Errorf("esperava email 'john2@gmail.com', got: %s", user.Email)
	}

	if user.Phone != "859947261212" {
		t.Errorf("esperava phone '859947261212', got: %s", user.Phone)
	}

	if user.Cpf != "12167898732" {
		t.Errorf("esperava cpf '12167898732', got: %s", user.Cpf)
	}
}

func TestUpdateByUserId_UserNotFound(t *testing.T) {
	mock := &mockRepository{
		GetUserByIdFn: func(ctx context.Context, id int) (*domain.User, error) {
			return nil, userexc.ErrUserNotFound
		},
	}

	svc := service.NewService(mock)

	user, err := svc.UpdateUserById(
		context.Background(),
		0,
		"",
		"",
		"",
		"",
	)

	if !errors.Is(err, userexc.ErrUserNotFound) {
		t.Fatalf("esperava ErrUserNotFound, got: %v", err)
	}

	if user != nil {
		t.Fatal("esperava nil, got usuário")
	}
}

func TestUpdateByUserId_UserNotUpdated(t *testing.T) {
	mock := &mockRepository{
		UpdateUserByIdFn: func(ctx context.Context, id int, user *domain.User) (*domain.User, error) {
			return nil, userexc.ErrUserNotFound
		},
		GetUserByIdFn: func(ctx context.Context, id int) (*domain.User, error) {
			return &domain.User{ID: 1}, nil
		},
	}

	svc := service.NewService(mock)

	user, err := svc.UpdateUserById(
		context.Background(),
		1,
		"John2",
		"john2@gmail.com",
		"859947261212",
		"12167898732",
	)

	if !errors.Is(err, userexc.ErrUserNotFound) {
		t.Fatalf("esperava ErrUserNotFound, got: %v", err)
	}

	if user != nil {
		t.Fatal("esperava nil, got usuário")
	}
}

func TestDeleteUserById_Success(t *testing.T) {
	mock := &mockRepository{
		DeleteUserByIdFn: func(ctx context.Context, id int) error {
			return nil
		},
		GetUserByIdFn: func(ctx context.Context, id int) (*domain.User, error) {
			return &domain.User{
				Name:  "John doe",
				Email: "johndoe@gmail.com",
			}, nil
		},
	}

	svc := service.NewService(mock)

	err := svc.DeleteUserById(
		context.Background(),
		int(1),
	)

	if err != nil {
		t.Fatalf("esperava err nil, got: %v", err)
	}
}

func TestDeleteUserById_UserNotFound(t *testing.T) {
	mock := &mockRepository{
		GetUserByIdFn: func(ctx context.Context, id int) (*domain.User, error) {
			return nil, userexc.ErrUserNotFound
		},
	}

	svc := service.NewService(mock)

	err := svc.DeleteUserById(
		context.Background(),
		1,
	)

	if !errors.Is(err, userexc.ErrUserNotFound) {
		t.Fatalf("esperava ErrUserNotFound, got: %v", err)
	}
}
