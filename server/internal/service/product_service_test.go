package service_test

import (
	"context"
	"errors"
	"server/internal/domain"
	"server/internal/util/pagination"
	"server/internal/service"
	"testing"
)

type mockProductRepository struct {
	CreateProductFn     func(ctx context.Context, product *domain.Product) (*domain.Product, error)
	FindAllProductsFn   func(ctx context.Context, params pagination.Params) ([]domain.Product, int64, error)
	FindProductByIdFn   func(ctx context.Context, id int) (*domain.Product, error)
	UpdateByProductIdFn func(ctx context.Context, id int, product *domain.Product) (*domain.Product, error)
	DeleteByProductIdFn func(ctx context.Context, id int) error
}

func (m *mockProductRepository) CreateProduct(ctx context.Context, product *domain.Product) (*domain.Product, error) {
	return m.CreateProductFn(ctx, product)
}

func (m *mockProductRepository) FindAllProducts(ctx context.Context, params pagination.Params) ([]domain.Product, int64, error) {
	return m.FindAllProductsFn(ctx, params)
}

func (m *mockProductRepository) FindProductById(ctx context.Context, id int) (*domain.Product, error) {
	return m.FindProductByIdFn(ctx, id)
}

func (m *mockProductRepository) UpdateByProductId(ctx context.Context, id int, product *domain.Product) (*domain.Product, error) {
	return m.UpdateByProductIdFn(ctx, id, product)
}

func (m *mockProductRepository) DeleteByProductId(ctx context.Context, id int) error {
	return m.DeleteByProductIdFn(ctx, id)
}

func TestCreateProduct_Success(t *testing.T) {
	mock := &mockProductRepository{
		CreateProductFn: func(ctx context.Context, product *domain.Product) (*domain.Product, error) {
			return product, nil
		},
	}

	svc := service.NewProductService(mock)

	fakeProduct := &domain.Product{
		Name:        "Furadeira Profissional",
		Price:       49.90,
		Description: "Furadeira elétrica 650W para uso profissional",
		Stock:       "10",
		Is_on_sale:  false,
		Weight:      2.5,
		Width:       30.0,
		Diameter:    0,
		Length:      25,
	}

	product, err := svc.CreateProduct(context.Background(), fakeProduct)

	if err != nil {
		t.Fatalf("esperava err nil, got: %v", err)
	}
	if product == nil {
		t.Fatal("esperava produto, got nil")
	}
	if product.Name != fakeProduct.Name {
		t.Errorf("esperava name %q, got %q", fakeProduct.Name, product.Name)
	}
}

func TestCreateProduct_PriceIsLessThanFive(t *testing.T) {
	svc := service.NewProductService(&mockProductRepository{})

	fakeProduct := &domain.Product{
		Name:        "Furadeira Profissional",
		Price:       2.0,
		Description: "Furadeira elétrica 650W para uso profissional",
		Stock:       "10",
		Is_on_sale:  false,
		Weight:      2.5,
		Width:       30.0,
		Diameter:    0,
		Length:      25,
	}

	product, err := svc.CreateProduct(context.Background(), fakeProduct)

	if product != nil {
		t.Errorf("esperava product nil, got: %+v", product)
	}
	if !errors.Is(err, domain.ErrPriceNotBeLessThanFive) {
		t.Fatalf("esperava ErrPriceNotBeLessThanFive, got: %v", err)
	}
}

func TestGetAllProducts_Success(t *testing.T) {
	mock := &mockProductRepository{
		FindAllProductsFn: func(ctx context.Context, params pagination.Params) ([]domain.Product, int64, error) {
			return []domain.Product{
				{ID: 1, Name: "teste bem testado"},
			}, 1, nil
		},
	}
	svc := service.NewProductService(mock)

	products, err := svc.GetAllProducts(context.Background(), pagination.Params{
		Limit:  10,
		Offset: 100,
	})

	if err != nil {
		t.Fatalf("esperava err nil, got: %v", err)
	}
	if products == nil {
		t.Fatal("esperava lista de produtos, got nil")
	}
	if len(products.Data) != 1 {
		t.Fatalf("esperava 1 produto, got %d", len(products.Data))
	}
	if products.Data[0].Name != "teste bem testado" {
		t.Errorf("esperava name 'teste bem testado', got: %s", products.Data[0].Name)
	}
}

func TestGetAllProducts_Failed(t *testing.T) {
	mock := &mockProductRepository{
		FindAllProductsFn: func(ctx context.Context, params pagination.Params) ([]domain.Product, int64, error) {
			return nil, 0 ,errors.New("database error")
		},
	}
	svc := service.NewProductService(mock)

	products, err := svc.GetAllProducts(context.Background(), pagination.Params{
		Limit: 10,
		Offset: 100,
	},)

	if products != nil {
		t.Errorf("esperava products nil, got: %+v", products)
	}
	if err == nil {
		t.Fatal("esperava erro, got nil")
	}
}

func TestGetProductById_Success(t *testing.T) {
	mock := &mockProductRepository{
		FindProductByIdFn: func(ctx context.Context, id int) (*domain.Product, error) {
			return &domain.Product{
				ID: 1, Name: "teste bem testado",
			}, nil
		},
	}
	svc := service.NewProductService(mock)

	product, err := svc.GetProductById(context.Background(), 1)

	if err != nil {
		t.Fatalf("esperava err nil, got: %v", err)
	}
	if product == nil {
		t.Fatal("esperava produto, got nil")
	}
	if product.Name != "teste bem testado" {
		t.Errorf("esperava name 'teste bem testado', got: %s", product.Name)
	}
}

func TestGetProductById_ProductNotFound(t *testing.T) {
	mock := &mockProductRepository{
		FindProductByIdFn: func(ctx context.Context, id int) (*domain.Product, error) {
			return nil, domain.ErrProductNotFound
		},
	}

	svc := service.NewProductService(mock)

	product, err := svc.GetProductById(context.Background(), 1)

	if product != nil {
		t.Errorf("esperava product nil, got: %+v", product)
	}
	if !errors.Is(err, domain.ErrProductNotFound) {
		t.Fatalf("esperava ErrProductNotFound, got: %v", err)
	}
}

func TestUpdateProduct_Success(t *testing.T) {
	const productID = 1

	oldProduct := &domain.Product{
		Name:        "Furadeira Profissional",
		Price:       49.90,
		Description: "Furadeira elétrica 650W para uso profissional",
		Stock:       "10",
		Is_on_sale:  false,
		Weight:      2.5,
		Width:       30.0,
		Diameter:    0,
		Length:      25,
	}

	updatedProduct := &domain.Product{
		Name:        "Furadeira Profissional 2",
		Price:       49.91,
		Description: "Furadeira elétrica 650W para uso profissional 2",
		Stock:       "10",
		Is_on_sale:  true,
		Weight:      2.5,
		Width:       30.0,
		Diameter:    0,
		Length:      25,
	}

	mock := &mockProductRepository{
		UpdateByProductIdFn: func(ctx context.Context, id int, product *domain.Product) (*domain.Product, error) {
			return product, nil
		},
	}

	svc := service.NewProductService(mock)

	product, err := svc.UpdateByProduct(context.Background(), productID, updatedProduct)

	if err != nil {
		t.Fatalf("esperava err nil, got: %v", err)
	}
	if product == nil {
		t.Fatal("esperava produto, got nil")
	}

	if product.Name != updatedProduct.Name {
		t.Errorf("Name: esperava %q, got %q", updatedProduct.Name, product.Name)
	}
	if product.Price != updatedProduct.Price {
		t.Errorf("Price: esperava %v, got %v", updatedProduct.Price, product.Price)
	}
	if product.Description != updatedProduct.Description {
		t.Errorf("Description: esperava %q, got %q", updatedProduct.Description, product.Description)
	}
	if product.Is_on_sale != updatedProduct.Is_on_sale {
		t.Errorf("Is_on_sale: esperava %v, got %v", updatedProduct.Is_on_sale, product.Is_on_sale)
	}

	if product.Name == oldProduct.Name {
		t.Errorf("Name não foi substituído: ainda é %q", oldProduct.Name)
	}
	if product.Price == oldProduct.Price {
		t.Errorf("Price não foi substituído: ainda é %v", oldProduct.Price)
	}
	if product.Description == oldProduct.Description {
		t.Errorf("Description não foi substituída: ainda é %q", oldProduct.Description)
	}
	if product.Is_on_sale == oldProduct.Is_on_sale {
		t.Error("Is_on_sale não foi substituído: ainda é false")
	}
}

func TestUpdateProduct_FailedProductNotFound(t *testing.T) {
	mock := &mockProductRepository{
		UpdateByProductIdFn: func(ctx context.Context, id int, product *domain.Product) (*domain.Product, error) {
			return nil, domain.ErrProductNotFound
		},
	}

	svc := service.NewProductService(mock)

	fakeProduct := &domain.Product{
		Name:        "Furadeira Profissional 2",
		Price:       49.90,
		Description: "Furadeira elétrica 650W para uso profissional 2",
		Stock:       "10",
		Is_on_sale:  false,
		Weight:      2.5,
		Width:       30.0,
		Diameter:    0,
		Length:      25,
	}

	product, err := svc.UpdateByProduct(context.Background(), 1, fakeProduct)

	if product != nil {
		t.Errorf("esperava product nil, got: %+v", product)
	}
	if !errors.Is(err, domain.ErrProductNotFound) {
		t.Fatalf("esperava ErrProductNotFound, got: %v", err)
	}
}

func TestDeleteProduct_Success(t *testing.T) {
	mock := &mockProductRepository{
		DeleteByProductIdFn: func(ctx context.Context, id int) error {
			return nil
		},
	}

	svc := service.NewProductService(mock)
	err := svc.DeleteByProduct(
		context.Background(),
		int(1),
	)

	if errors.Is(err, domain.ErrProductNotFound) {
		t.Fatalf("esperava ErrProductNotFound, got: %v", err)
	}

	if err != nil {
		t.Fatalf("esperava err, got: %v", err)
	}
}

func TestDeleteProduct_FailedProductNotFound(t *testing.T) {
	mock := &mockProductRepository{
		DeleteByProductIdFn: func(ctx context.Context, id int) error {
			return domain.ErrProductNotFound
		},
	}

	svc := service.NewProductService(mock)
	err := svc.DeleteByProduct(
		context.Background(),
		int(1),
	)

	if !errors.Is(err, domain.ErrProductNotFound) {
		t.Fatalf("esperava ErrProductNotFound, got: %v", err)
	}
}

func TestDeleteProduct_Failed(t *testing.T) {
	mock := &mockProductRepository{
		DeleteByProductIdFn: func(ctx context.Context, id int) error {
			return errors.New("database error")
		},
	}

	svc := service.NewProductService(mock)
	err := svc.DeleteByProduct(
		context.Background(),
		int(1),
	)

	if err == nil {
		t.Fatalf("esperava err, got: %v", err)
	}
}
