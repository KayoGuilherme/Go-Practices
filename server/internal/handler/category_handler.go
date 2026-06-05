package handler

import (
	"server/internal/domain"
	"server/internal/dto/request"
	"server/internal/mapper"
	"server/internal/util/http/responses"
	"server/internal/validators"

	"github.com/gin-gonic/gin"
)

type CategoryHandler struct {
	categorysvc domain.CategoryService
}

func NewCategoryHandler(categorysvc domain.CategoryService) *CategoryHandler {
	return &CategoryHandler{
		categorysvc: categorysvc,
	}
}

func (h *CategoryHandler) CreateCategory(ctx *gin.Context) {
	var dto request.CategoryRequestDTO
	if validators.ShouldBindJSON(ctx, &dto) != nil {
		return
	}
	if validators.ValidateStruct(ctx, &dto) != nil {
		return
	}

	category := mapper.ToCategoryDomain(dto)

	_, err := h.categorysvc.CreateCategory(ctx, category)
	if err != nil {
		responses.HandleError(ctx, err)
		return
	}

	responses.Created(ctx, mapper.ToCategoryResponse(category))
}

func (h *CategoryHandler) GetAllCategorys(ctx *gin.Context) {
	products, err := h.categorysvc.GetAllCategory(ctx)
	if err != nil {
		responses.HandleError(ctx, err)
		return
	}

	responses.OK(ctx, mapper.ToCategoryResponses(products))
}

func (h *CategoryHandler) GetCategoryById(ctx *gin.Context) {
	id, err := validators.ParsePathInt(ctx, "id")
	if err != nil {
		return
	}

	category, err := h.categorysvc.GetCategory(ctx, id)
	if err != nil {
		responses.HandleError(ctx, err)
		return
	}

	responses.OK(ctx, mapper.ToCategoryResponse(category))
}
