package handler

import (
	"errors"

	"server/internal/domain"
	"server/internal/dto/request"
	"server/internal/mapper"
	"server/internal/util/http/responses"
	"server/internal/validators"

	"github.com/gin-gonic/gin"
)

type CategoryProductHandler struct {
	linkService domain.CategoryProductService
}

func NewCategoryProductHandler(linkService domain.CategoryProductService) *CategoryProductHandler {
	return &CategoryProductHandler{linkService: linkService}
}

func (h *CategoryProductHandler) Attach(ctx *gin.Context) {
	productID, err := validators.ParsePathInt(ctx, "id")
	if err != nil {
		return
	}

	var dto request.LinkCategoryDTO
	if validators.ShouldBindJSON(ctx, &dto) != nil {
		return
	}
	if validators.ValidateStruct(ctx, &dto) != nil {
		return
	}

	if err := h.linkService.AttachCategory(ctx.Request.Context(), productID, dto.CategoryID); err != nil {
		handleLinkError(ctx, err)
		return
	}

	responses.Created(ctx, gin.H{"product_id": productID, "category_id": dto.CategoryID})
}

func (h *CategoryProductHandler) Detach(ctx *gin.Context) {
	productID, err := validators.ParsePathInt(ctx, "id")
	if err != nil {
		return
	}

	categoryID, err := validators.ParsePathInt(ctx, "categoryId")
	if err != nil {
		return
	}

	if err := h.linkService.DetachCategory(ctx.Request.Context(), productID, categoryID); err != nil {
		handleLinkError(ctx, err)
		return
	}

	responses.NoContent(ctx)
}

func (h *CategoryProductHandler) ListCategoriesByProduct(ctx *gin.Context) {
	productID, err := validators.ParsePathInt(ctx, "id")
	if err != nil {
		return
	}

	categories, err := h.linkService.GetCategoriesByProduct(ctx.Request.Context(), productID)
	if err != nil {
		handleLinkError(ctx, err)
		return
	}

	responses.OK(ctx, mapper.ToCategoryResponses(categories))
}

func (h *CategoryProductHandler) ListProductsByCategory(ctx *gin.Context) {
	categoryID, err := validators.ParsePathInt(ctx, "id")
	if err != nil {
		return
	}

	products, err := h.linkService.GetProductsByCategory(ctx.Request.Context(), categoryID)
	if err != nil {
		handleLinkError(ctx, err)
		return
	}

	responses.OK(ctx, mapper.ToProductResponses(products))
}

func handleLinkError(ctx *gin.Context, err error) {
	switch {
	case errors.Is(err, domain.ErrProductNotFound),
		errors.Is(err, domain.ErrCategoryNotFound),
		errors.Is(err, domain.ErrLinkNotFound):
		responses.NotFound(ctx, err)
	case errors.Is(err, domain.ErrLinkAlreadyExists):
		responses.Conflict(ctx, err)
	default:
		responses.InternalServerError(ctx, err)
	}
}
