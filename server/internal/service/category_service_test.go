package service_test

import (
	"context"
	"errors"
	"server/internal/domain"
	"server/internal/service"
	"testing"
)

type mockCategoryRepository struct {
	CreateCategoryFn   func(ctx context.Context, category *domain.Category) (*domain.Category, error)
	FindAllCategoryFn  func(ctx context.Context) ([]domain.Category, error)
	FindByCategoryIdFn func(ctx context.Context, id int) (*domain.Category, error)
}

func (m *mockCategoryRepository) CreateCategory(ctx context.Context, category *domain.Category) (*domain.Category, error) {
	return m.CreateCategoryFn(ctx, category)
}

func (m *mockCategoryRepository) FindAllCategory(ctx context.Context) ([]domain.Category, error) {
	return m.FindAllCategoryFn(ctx)
}

func (m *mockCategoryRepository) FindByCategoryId(ctx context.Context, id int) (*domain.Category, error) {
	return m.FindByCategoryIdFn(ctx, id)
}

func TestCreateCategory_Success(t *testing.T) {
	mock := &mockCategoryRepository{
		CreateCategoryFn: func(ctx context.Context, category *domain.Category) (*domain.Category, error) {
			return category, nil
		},
	}

	svc := service.NewCategoryService(mock)

	fakeCategory := &domain.Category{Name: "Ferramentas"}

	category, err := svc.CreateCategory(context.Background(), fakeCategory)

	if err != nil {
		t.Fatalf("esperava err nil, got: %v", err)
	}
	if category == nil {
		t.Fatal("esperava categoria, got nil")
	}
	if category.Name != fakeCategory.Name {
		t.Errorf("esperava name %q, got %q", fakeCategory.Name, category.Name)
	}
}

func TestCreateCategory_Failed(t *testing.T) {
	mock := &mockCategoryRepository{
		CreateCategoryFn: func(ctx context.Context, category *domain.Category) (*domain.Category, error) {
			return nil, errors.New("database error")
		},
	}

	svc := service.NewCategoryService(mock)

	category, err := svc.CreateCategory(context.Background(), &domain.Category{Name: "Ferramentas"})

	if category != nil {
		t.Errorf("esperava category nil, got: %+v", category)
	}
	if err == nil {
		t.Fatal("esperava erro, got nil")
	}
}

func TestGetAllCategory_Success(t *testing.T) {
	mock := &mockCategoryRepository{
		FindAllCategoryFn: func(ctx context.Context) ([]domain.Category, error) {
			return []domain.Category{
				{ID: 1, Name: "Ferramentas"},
				{ID: 2, Name: "Construção"},
			}, nil
		},
	}

	svc := service.NewCategoryService(mock)

	categories, err := svc.GetAllCategory(context.Background())

	if err != nil {
		t.Fatalf("esperava err nil, got: %v", err)
	}
	if len(categories) != 2 {
		t.Fatalf("esperava 2 categorias, got %d", len(categories))
	}
	if categories[0].Name != "Ferramentas" {
		t.Errorf("esperava name 'Ferramentas', got: %s", categories[0].Name)
	}
}

func TestGetAllCategory_Failed(t *testing.T) {
	mock := &mockCategoryRepository{
		FindAllCategoryFn: func(ctx context.Context) ([]domain.Category, error) {
			return nil, errors.New("database error")
		},
	}

	svc := service.NewCategoryService(mock)

	categories, err := svc.GetAllCategory(context.Background())

	if categories != nil {
		t.Errorf("esperava categories nil, got: %+v", categories)
	}
	if err == nil {
		t.Fatal("esperava erro, got nil")
	}
}

func TestGetCategory_Success(t *testing.T) {
	mock := &mockCategoryRepository{
		FindByCategoryIdFn: func(ctx context.Context, id int) (*domain.Category, error) {
			return &domain.Category{ID: 1, Name: "Ferramentas"}, nil
		},
	}

	svc := service.NewCategoryService(mock)

	category, err := svc.GetCategory(context.Background(), 1)

	if err != nil {
		t.Fatalf("esperava err nil, got: %v", err)
	}
	if category == nil {
		t.Fatal("esperava categoria, got nil")
	}
	if category.Name != "Ferramentas" {
		t.Errorf("esperava name 'Ferramentas', got: %s", category.Name)
	}
}

func TestGetCategory_CategoryNotFound(t *testing.T) {
	mock := &mockCategoryRepository{
		FindByCategoryIdFn: func(ctx context.Context, id int) (*domain.Category, error) {
			return nil, domain.ErrCategoryNotFound
		},
	}

	svc := service.NewCategoryService(mock)

	category, err := svc.GetCategory(context.Background(), 999)

	if category != nil {
		t.Errorf("esperava category nil, got: %+v", category)
	}
	if !errors.Is(err, domain.ErrCategoryNotFound) {
		t.Fatalf("esperava ErrCategoryNotFound, got: %v", err)
	}
}

func TestGetCategory_Failed(t *testing.T) {
	mock := &mockCategoryRepository{
		FindByCategoryIdFn: func(ctx context.Context, id int) (*domain.Category, error) {
			return nil, errors.New("database error")
		},
	}

	svc := service.NewCategoryService(mock)

	category, err := svc.GetCategory(context.Background(), 1)

	if category != nil {
		t.Errorf("esperava category nil, got: %+v", category)
	}
	if err == nil {
		t.Fatal("esperava erro, got nil")
	}
	if errors.Is(err, domain.ErrCategoryNotFound) {
		t.Fatalf("esperava erro genérico, got ErrCategoryNotFound")
	}
}