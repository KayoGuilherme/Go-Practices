package mapper

import (
	"server/internal/domain"
	"server/internal/dto/request"
	"server/internal/dto/response"
)

func ToCategoryDomain(dto request.CategoryRequestDTO) *domain.Category {
	return &domain.Category{
		Name: dto.Name,
	}
}

func ToCategoryResponse(domain *domain.Category) response.CategoryResponseDTO {

	if domain == nil {
		return response.CategoryResponseDTO{}
	}

	return response.CategoryResponseDTO{
		ID:        domain.ID,
		Name:      domain.Name,
		CreatedAt: domain.CreatedAt,
	}
}

func ToCategoryResponses(categorys []domain.Category) []response.CategoryResponseDTO {
	out := make([]response.CategoryResponseDTO, 0, len(categorys))
	for i := range categorys {
		out = append(out, ToCategoryResponse(&categorys[i]))
	}
	return out
}