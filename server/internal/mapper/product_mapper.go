package mapper

import (
	"server/internal/domain"
	"server/internal/dto/request"
	"server/internal/dto/response"
	"server/internal/util/pagination"
)

func ToProductDomain(dto request.CreateProductDTO) *domain.Product {
	return &domain.Product{
		Name:        dto.Name,
		Price:       dto.Price,
		Description: dto.Description,
		Stock:       dto.Stock,
		Is_on_sale:  dto.Is_on_sale,
		Weight:      dto.Weight,
		Width:       dto.Width,
		Diameter:    dto.Diameter,
		Length:      dto.Length,
	}
}

func ToProductDomainUpdate(dto request.UpdateProductDTO) *domain.Product {
	return &domain.Product{
		Name:        dto.Name,
		Price:       dto.Price,
		Description: dto.Description,
		Stock:       dto.Stock,
		Is_on_sale:  dto.Is_on_sale,
		Weight:      dto.Weight,
		Width:       dto.Width,
		Diameter:    dto.Diameter,
		Length:      dto.Length,
	}
}

func ToProductResponse(product *domain.Product) response.ProductResponseDTO {
	if product == nil {
		return response.ProductResponseDTO{}
	}
	return response.ProductResponseDTO{
		ID:          product.ID,
		Name:        product.Name,
		Price:       product.Price,
		Description: product.Description,
		Stock:       product.Stock,
		Is_on_sale:  product.Is_on_sale,
		Weight:      product.Weight,
		Width:       product.Width,
		Diameter:    product.Diameter,
		Length:      product.Length,
		CreatedAt:   product.CreatedAt,
		UpdatedAt:   product.UpdatedAt,
	}
}

func ToProductPaginateResponse(page *pagination.Response[domain.Product]) pagination.Response[response.ProductResponseDTO] {
	return pagination.Response[response.ProductResponseDTO]{
        Data:    ToProductResponses(page.Data),
        Total:   page.Total,
        Limit:   page.Limit,
        Offset:  page.Offset,
        HasMore: page.HasMore,
    }
}

func ToProductResponses(products []domain.Product) []response.ProductResponseDTO {
	out := make([]response.ProductResponseDTO, 0, len(products))
	for i := range products {
		out = append(out, ToProductResponse(&products[i]))
	}
	return out
}
